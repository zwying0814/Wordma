import { createApp } from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import piniaPluginPersist from 'pinia-plugin-persistedstate'
import router from "./router";

import 'virtual:uno.css'
import '@unocss/reset/tailwind-compat.css'
import './style.css'


const pinia = createPinia()
pinia.use(piniaPluginPersist)

const app = createApp(App)

app.use(router)
app.use(pinia)

app.mount('#app')
