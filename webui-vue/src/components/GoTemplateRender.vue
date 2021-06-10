<template>
  <div id="app">
    <form class="container">
      <div class="row">
        <div class="col-md-6">
          <h1>Template</h1>
          <PrismEditor
            class="my-editor"
            id="template"
            v-model="template"
            :highlight="highlighterGo"
            line-numbers
          ></PrismEditor>
        </div>
        <div class="col-md-6">
          <h1>Values (Yaml)</h1>
          <PrismEditor
            class="my-editor"
            id="values"
            v-model="values"
            :highlight="highlighterYml"
            line-numbers
          ></PrismEditor>
        </div>
      </div>
      <div class="row">
        <div class="col-md-12 buttons">
          <vs-tooltip
            bottom
            color="#7d33ff">
            <vs-switch 
              val="I'm debugging Helm" 
              v-model="helmEnabled"
              primary
              color="#7d33ff"
            >
              I'm debugging Helm
            </vs-switch>
            <template #tooltip>
              <div class="content-tooltip">
                <h4>
                  What does it mean?
                </h4>
                <p>
                  When debugging Helm charts you will have the '.Values' prefix with your values.
                  Toggling this will save you adding 'Values' to the top of the yaml here.
                </p>
              </div>
            </template>
          </vs-tooltip>
          <vs-button
            type="button"
            gradient
            size="xl"
            color="danger"
            @click="clear"
          >
            Clear
          </vs-button>
          <vs-button
            :loading="renderLoading"
            type="button"
            gradient
            size="xl"
            color="success"
            @click="render"
          >
            Render
          </vs-button>
        </div>
        <div class="col-md-12">
          <h1>Render</h1>
            <PrismEditor
              class="my-editor"
              id="render"
              v-model="renderData"
              :highlight="highlighterYml"
              line-numbers
            ></PrismEditor>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import { PrismEditor } from "vue-prism-editor";
import "vue-prism-editor/dist/prismeditor.min.css"; // import the styles somewhere

// import highlighting library (you can use any library you want just return html string)
import { highlight, languages } from "prismjs/components/prism-core";
// import "prismjs/components/prism-javascript";
import "prismjs/components/prism-yaml";
// import "prismjs/themes/prism-tomorrow.css"; // import syntax highlighting styles
const yaml = require("js-yaml");

export default {
  name: "GoTemplateRender",
  components: {
    PrismEditor,
  },
  data() {
    return {
      values: "myval: 5",
      template: "{{ .myval }}",
      renderLoading: false,
      renderData: "",
      helmEnabled: false,
    };
  },
  methods: {
    highlighterYml(code) {
      return highlight(code, languages.yml, "yaml"); //returns html
    },
    highlighterGo(code) {
      return highlight(code, languages.yaml, "yaml"); //returns html
    },
    yamlToJson() {
      try {
        const doc = yaml.load(this.values);
        return doc;
      } catch (e) {
        return e;
      }
    },

    render() {
      this.renderLoading = true;
      let myval = this.yamlToJson(this.values);
      if (this.helmEnabled) {
        myval = { Values: myval };
      }
      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          template: JSON.stringify(this.template),
          values: JSON.stringify(myval),
        }),
      };
      fetch("/render", requestOptions)
        .then((response) => {
          if (!response.ok) {
            this.renderLoading = false;
            return response.text();
          }
          return response.text();
        })
        .then((data) => {
          this.renderData = data
          this.renderLoading = false;
          });
    },
    clear() {
      this.renderData = "";
    }
  },
};
</script>
<style >
.prism-editor-wrapper {
  height: auto;
}

.my-editor {
  background: whitesmoke;
  color: #ccc;
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 5px;
  border-radius: 1em;
}

textarea:focus-visible {
  outline: none !important;
}

#template,
#render,
#values {
  width: 100%;
  min-height: 400px;
  white-space: pre;
  text-align: start;
  border-radius: 1em;
}

#render {
  background: #eee;
  padding: 1em;
}

span.whitespace {
  color: #ccc;
}

label {
  font-weight: normal;
}

.buttons {
  margin-top: 10px;
  color: white;
}

.buttons .vs-button {
  margin: 0px 5px;
  font-size: 2em;
  position: relative;
  float: right;
  outline: none;
}

.buttons .vs-switch {
  margin: 0px 5px;
  width: 18vw;
  position: relative;
  float: left;
  height: 2em;
  font-size: 1em;
}
.vs-switch .vs-switch__circle {
  background-color:rgb(178, 195, 197);
  width: 1.2em;
  height: 1.2em;
}

.vs-switch .vs-switch__text {
  font-size: 1em;
}

.vs-tooltip-content {
  position: absolute;
}
</style>
