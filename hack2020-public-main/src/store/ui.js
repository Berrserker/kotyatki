const defaultState = {
  isLoginModalOpened: false
}

export default {
  namespaced: true,
  state: { ...defaultState },
  mutations: {
    setLoginModalOpened (state, payload) {
      state.isLoginModalOpened = payload
    }
  },
  actions: {
    openLoginModal ({ commit }) {
      commit('setLoginModalOpened', true)
    },
    closeLoginModal ({ commit }) {
      commit('setLoginModalOpened', false)
    }
  }
}
