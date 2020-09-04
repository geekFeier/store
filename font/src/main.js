import Vue from 'vue'
import iView from 'iview'
import 'iview/dist/styles/iview.css'
import VueResource from 'vue-resource'
import VueRouter from 'vue-router'
import VueCookies from 'vue-cookies'

Vue.use(VueRouter)
Vue.use(VueResource);
Vue.use(VueCookies);
VueCookies.config('7d')

import App from './App.vue'

Vue.config.productionTip = false

Vue.use(iView)



const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [{
    path: '/',
    name: 'index',
    component: () =>
      import( /* webpackChunkName: "index" */ '@/views/index/index.vue'),
    meta: {
      title: 'sealYun',
    }
  }, {
    path: '/comment',
    name: 'comment',
    component: () =>
      import( /* webpackChunkName: "comment" */ '@/views/comment/index.vue'),
    meta: {
      title: 'sealYun 评论区',
    }
  }, {
    path: '*',
    components: () =>
      import('./components/404.vue')
  }]
})
new Vue({
  el: '#app',
  router,
  render: h => h(App),
})