import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import vuetify from './plugins/vuetify/vuetify'
import { loadFonts } from './/plugins/vuetify/webfontloader'
import installPiniaPlugin from '@/plugins/pinia/installPiniaPlugin'; // Pinia 状态监控
// import installFFmpeg from '@/plugins/ffmpeg/installFFmpeg'; // ffmpeg 集成
import './index.css'

loadFonts()
const pinia = createPinia()
pinia.use(installPiniaPlugin);

createApp(App)
  .use(router)
  .use(pinia)
  .use(vuetify)
  // .use(installFFmpeg)
  .mount('#app')
