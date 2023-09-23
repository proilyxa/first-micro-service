import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'

import './css/style.css'
import 'primevue/resources/themes/lara-light-blue/theme.css'
import 'primeicons/primeicons.css'

import App from './App.vue'
import router from './router'

import './lib/http'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue, { ripple: true })

app.component('InputText', InputText)
app.component('InputNumber', InputNumber)

app.mount('#app')
