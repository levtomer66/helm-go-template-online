<template>
  <div id="app">
    <vs-tooltip shadow>
      <a href="https://lusha.com" target="_blank">
        <div class="company-logo">
          <img src="./assets/lusha.svg" />
        </div>
      </a>
      <template #tooltip>
        Made by Lusha
      </template>
    </vs-tooltip>
    <div class="product-logo">
      <img src="./assets/lusha.svg" />
    </div>
    <GoTemplateRender />
    
  </div>
</template>

<script>
import 'boxicons'
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
.social-buttons {
  float: right;
  margin-top: 20vh;
}

.social-buttons .vs-button {
  width: 50px;
  height: 50px;
}

.product-logo img {
  max-width: 25vw;
}

.company-logo {
  padding-left: 2em;
  padding-bottom: 2em;
  position: fixed;
  top: 0;
  right: 0;
  background-color:black;
  border-bottom-left-radius: 10em;
  box-shadow:
    0 0 30px #fff,
    -10px 0 60px #f0f,
    10px 0 60px #0ff;
  overflow: hidden;
  transition: transform .5s ease;
}

.company-logo:hover {
  transform: scale(1.05);
}

.company-logo img {
  max-width: 10em;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  margin-bottom: 60px;
}
</style>
