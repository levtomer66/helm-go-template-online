<template>
  <div id="app">
    <form class="container">
      <div class="row">
        <div class="col-md-5">
          <h1>Template</h1>
          <PrismEditor
            class="my-editor"
            id="template"
            v-model="template"
            :highlight="highlighter"
            line-numbers
          ></PrismEditor>
        </div>
        <div class="col-md-5">
          <h1>Render</h1>
          <div id="render">{{ renderData }}</div>
        </div>
        <div class="col-md-2">
          <h1>Settings</h1>
          <div id="settings">
            <label
              ><input
                type="checkbox"
                name="showwhitespaces"
                checked="checked"
              />
              Show whitespaces</label
            ><br />
            <label
              ><input type="checkbox" name="dummyvalues" /> Use dummy
              values</label
            >
            <label
              ><input
                type="radio"
                name="input_type"
                value="json"
                checked="checked"
              />
              JSON</label
            >
            <label
              ><input type="radio" name="input_type" value="yaml" /> YAML</label
            >
            <input
              type="button"
              class="btn btn-success"
              @click="render"
              id="convert"
              value="Convert"
            />
            <input
              type="button"
              class="btn btn-danger"
              id="clear"
              value="Clear"
            />
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col-md-5">
          <h1>Values</h1>
          <PrismEditor
            class="my-editor"
            id="values"
            v-model="values"
            :highlight="highlighter"
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
import "prismjs/components/prism-clike";
import "prismjs/components/prism-javascript";
import "prismjs/themes/prism-tomorrow.css"; // import syntax highlighting styles
const yaml = require("js-yaml");

export default {
  name: "GoTemplateRender",
  components: {
    PrismEditor,
  },
  data() {
    return {
      values: "",
      template: "",
      renderData: "",
    };
  },
  methods: {
    highlighter(code) {
      return highlight(code, languages.js); //returns html
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
      console.log(this.values);
      console.log(this.template);
      const myval = this.yamlToJson(this.values);
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
            return response.text()
            // throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => (this.renderData = data))
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style >
.prism-editor-wrapper {
  height: auto;
}
/* required class */
.my-editor {
  /* we dont use `language-` classes anymore so thats why we need to add background and text color manually */
  background: #4e4e4e;
  color: #ccc;

  /* you must provide font-family font-size line-height. Example: */
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 5px;
}

/* optional class for removing the outline */
textarea:focus-visible {
  outline: none !important;
}

#template,
#render,
#values {
  width: 100%;
  min-height: 385px;
}

#render {
  background: #eee;
}

span.whitespace {
  color: #ccc;
}

label {
  font-weight: normal;
}
</style>
