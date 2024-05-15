import { createStore, createLogger } from 'vuex'
import product from './modules/product'

const debug = process.env.NODE_ENV !== 'production'

export default createStore({
  modules: {
    product
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})