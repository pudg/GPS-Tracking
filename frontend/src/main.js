import './style.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueGoogleMaps from '@fawmi/vue-google-maps'

import Vue from 'vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VirtualList from 'vue-virtual-scroll-list'

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
// Vue.component('virtual-list', VirtualList)
// Vue.config.productionTip = false

const app = createApp(App)
app.use(router)
app.use(store)
app.use(VueGoogleMaps, {
    load: {
        key: import.meta.env.VITE_GOOGLE_MAP_API_KEY
    }
})
app.mount('#app')
