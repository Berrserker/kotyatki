import Vue from 'vue'

import { DICTIONARIES } from '@/config'

import { fetchDictionary, updateDictionary } from '@/services/dictionary'

const defaultState = {
  dictionaries: {},
  isLoading: false
}

export default {
  namespaced: true,
  state: { ...defaultState },
  mutations: {
    setDictionary (state, payload) {
      Vue.set(state.dictionaries, payload.key, payload.value)
    },
    setIsLoading (state, payload) {
      state.isLoading = payload
    },
    resetState (state) {
      for (const key of Object.keys(state)) {
        Vue.set(state, key, defaultState[key])
      }
    }
  },
  actions: {
    async fetchDictionary ({ commit }, keys) {
      commit('setIsLoading', true)
      const response = await fetchDictionary(keys)
      if (response.data) {
        for (const dict of response.data.vocs) {
          const key = DICTIONARIES.find(o => o.apiKey === dict.Name).key
          const value = dict.Voc.map(o => o.split('|'))
          commit('setDictionary', { key, value })
        }
      }
      commit('setIsLoading', false)
    },
    async updateDictionary ({ commit }, data) {
      const response = await updateDictionary(data.login, data.name, data.voc)
      console.log(response)
    }
  }
}
