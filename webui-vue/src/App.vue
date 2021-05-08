<template>
  <div id="app">
    <GoTemplateRender />
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
