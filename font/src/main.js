import Vue from 'vue'
import App from './App.vue'
import iView from 'iview'
import 'iview/dist/styles/iview.css'
import VueResource from 'vue-resource'
/*
import axios from 'axios'

axios.defaults.withCredentials = true; //让ajax携带cookie
Vue.prototype.$axios = axios;
*/

Vue.config.productionTip = false

Vue.use(iView)
Vue.use(VueResource);
new Vue({
    render: h => h(App),
}).$mount('#app')