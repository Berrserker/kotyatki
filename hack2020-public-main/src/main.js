import Vue from 'vue'
import '@/plugins/axios'
import App from '@/App.vue'
import router from '@/router'
import store from '@/store'

import '@/directives/click-outside'

import '@/assets/styles/reset.css'
import '@/assets/styles/fonts.sass'
import '@/assets/styles/general.sass'
import '@/assets/styles/ui.sass'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
