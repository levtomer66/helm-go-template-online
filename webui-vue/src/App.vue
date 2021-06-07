<template>
  <div id="app">
    <div class="product-logo">
      <img src="./assets/GoTemplateDebugger.png" />
    </div>
    <GoTemplateRender />
    <div class="company-logo">
      <img src="./assets/MeteorOps.png" />
    </div>
  </div>
</template>

<script>

import GoTemplateRender from './components/GoTemplateRender.vue'

export default {
  name: 'App',
  components: {
    GoTemplateRender
  }
}
</script>

<style>
body {
  background-color: #0d1c24;
}

.product-logo img {
  max-width: 25vw;
}

.company-logo img {
  max-width: 12vw;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
