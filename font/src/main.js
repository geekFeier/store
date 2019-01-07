import Vue from 'vue'
import iView from 'iview'
import 'iview/dist/styles/iview.css'
import VueResource from 'vue-resource'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(VueResource);

//import HelloWorld from './components/HelloWorld.vue'
import App from './App.vue'


/*
import axios from 'axios'

axios.defaults.withCredentials = true; //让ajax携带cookie
Vue.prototype.$axios = axios;
*/

Vue.config.productionTip = false

Vue.use(iView)



const router = new VueRouter({
    mode: 'history',
    base: __dirname,
    routes: [{
        path: '*',
        components: () =>
            import ('./components/HelloWorld.vue')
    }]
})
new Vue({
    el: '#app',
    router,
    render: h => h(App),
})