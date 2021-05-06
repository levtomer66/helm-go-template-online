<template>
  <div id="app">
    <form class="container">
      <div class="row">
        <div class="col-md-5">
          <h1>Template</h1>
          <textarea v-model="templates" id="template"></textarea>
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
          <textarea v-model="values" id="values"></textarea>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  name: "GoTemplateRender",
  data() {
    return {
      values: "",
      templates: "",
      renderData: ""
    };
  },
  methods: {
    render() {
      console.log(this.values);
      console.log(this.template);
      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify ({
                template: JSON.stringify(this.templates),
                values: this.values,
            })
      };
      fetch("/render", requestOptions)
        .then((response) => response.json())
        .then((data) => (this.renderData = data));
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#template,
#render,
#values {
  width: 100%;
  min-height: 300px;
  resize: vertical;
  font-family: "Courier New", Courier, monospace;
  border: 1px solid #ccc;
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
