import Vue from 'vue'

import { logIn, logOut } from '@/services/auth'
import { fetchUser } from '@/services/user'

const defaultState = {
  isAuthorized: false,
  isLoading: false,
  id: null,
  email: '',
  role: ''
}

export default {
  namespaced: true,
  state: { ...defaultState },
  mutations: {
    setIsAuthorized (state, payload) {
      state.isAuthorized = payload
    },
    setIsLoading (state, payload) {
      state.isLoading = payload
    },
    setUserData (state, payload) {
      for (const key of Object.keys(payload)) {
        if (Object.prototype.hasOwnProperty.call(state, key)) {
          Vue.set(state, key, payload[key])
        }
      }
    },
    resetState (state) {
      for (const key of Object.keys(state)) {
        Vue.set(state, key, defaultState[key])
      }
    }
  },
  actions: {
    async logIn ({ commit }, data) {
      commit('setIsLoading', true)
      const response = await logIn(data)
      if (response.data) {
        commit('setIsAuthorized', true)
      }
      commit('setIsLoading', false)
      return response
    },
    logOut ({ commit }) {
      commit('resetState')
      logOut()
    },
    async fetchUser ({ commit }) {
      commit('setIsLoading', true)
      const response = await fetchUser()
      if (response.data) {
        commit('setUserData', response.data.account)
        commit('setIsAuthorized', true)
      }
      commit('setIsLoading', false)
    }
  }
}
