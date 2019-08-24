import Vue from 'vue'
import Router from 'vue-router'
import Bitcoin from './views/Bitcoin'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'btc',
      component: Bitcoin
    }
  ]
})
