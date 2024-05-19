
// initial state
const state = () => ({
  products:{}
})

// getters
const getters = {
  products (state) {
    return state.products
  }
}

// actions
const actions = {

}

// mutations
const mutations = {
  add (state, product) {
      state.products[product._id] = product
  },
  adds (state, products) {
    for (const product of products) {
      state.products[product._id] = product
    }
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}