import Vue from 'vue'
import App from './App.vue'
import Vuesax from "vuesax";
import "vuesax/dist/vuesax.css";
import VueAnalytics from 'vue-analytics';

Vue.use(Vuesax);
Vue.use(VueAnalytics, {
  id: 'G-QQ0YMTLM6B',
})
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
