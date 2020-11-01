import Vue from 'vue'
import Vuex from 'vuex'
import user from './user'
import ui from './ui'
import dictionary from './dictionary'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user,
    ui,
    dictionary
  }
})
