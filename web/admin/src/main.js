import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
axios.defaults.baseURL = 'http://localhost:3000/api/v1'
import setupAtnd from '@/plugin/antui'

const app = createApp(App)
app.config.globalProperties.$http = axios
setupAtnd(app)
app.use(router).mount('#app')
