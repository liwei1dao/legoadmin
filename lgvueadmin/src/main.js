import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import websocket from './plugins/websocket'
import vuetify from './plugins/vuetify/vuetify'
import { loadFonts } from './plugins/vuetify/webfontloader'

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .provide('$websocket', websocket)
  .mount('#app')
