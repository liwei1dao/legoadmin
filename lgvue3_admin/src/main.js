import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import config from './config.js';
import websocket from './plugins/websocket'
import vuetify from './plugins/vuetify/vuetify'
import { loadFonts } from './plugins/vuetify/webfontloader'
import VueVideoPlayer from './plugins/videoplayer'

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .use(VueVideoPlayer)
  .provide('$config', config)
  .provide('$websocket', websocket)
  .mount('#app')
