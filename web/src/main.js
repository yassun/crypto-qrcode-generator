import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import axios from 'axios'
import titleMixin from './utils/title';

Vue.config.productionTip = false

Vue.prototype.$axios = axios

Vue.mixin(titleMixin)

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
