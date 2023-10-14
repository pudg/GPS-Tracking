import './style.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueGoogleMaps from '@fawmi/vue-google-maps'
const app = createApp(App)

app.use(router)
app.use(store)
app.use(VueGoogleMaps, {
    load: {
        key: import.meta.env.VITE_GOOGLE_MAP_API_KEY
    }
})
app.mount('#app')
