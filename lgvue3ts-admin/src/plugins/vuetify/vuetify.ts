// Styles
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify } from 'vuetify'
import { mdi } from 'vuetify/iconsets/mdi'

export default createVuetify(
  // https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
  {
      icons: {
          defaultSet: 'mdi',
          sets: {
              mdi
          }
      }
  }
)
