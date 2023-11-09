import './style.css'
import { createApp } from 'vue'
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueCookies from 'vue3-cookies'
import { globalCookiesConfig } from "vue3-cookies";
import VueGoogleMaps from '@fawmi/vue-google-maps'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueVirtualScroller from 'vue-virtual-scroller'
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css'

globalCookiesConfig({
    path: "/api",
    domain: "localhost",
    secure: false,
    sameSite: "None",
})

const app = createApp(App)
app.use(router)
router.app = app
app.use(VueCookies)
app.use(store)
app.use(VueVirtualScroller)
app.use(VueGoogleMaps, {
    load: {
        key: import.meta.env.VITE_GOOGLE_MAP_API_KEY
    }
})
app.mount('#app')
