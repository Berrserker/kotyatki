<template lang="pug">
#app
  template(v-if="isReady")
    .wrapper
      app-header.header
      router-view
    login-modal(v-if="isLoginModalOpened")
  template(v-else)
    ui-loader.loader(:size="120")
</template>

<script>
import { mapState, mapActions } from 'vuex'
import { DICTIONARIES } from '@/config'
import AppHeader from '@/components/_layout/Header/Header'
import LoginModal from '@/components/Login/LoginModal'
import UiLoader from '@/components/_ui/Loader'
export default {
  components: {
    AppHeader,
    LoginModal,
    UiLoader
  },
  provide () {
    return {
      user: this.user
    }
  },
  data () {
    return {
      isReady: false
    }
  },
  computed: {
    ...mapState({
      user: state => state.user,
      isLoginModalOpened: state => state.ui.isLoginModalOpened
    })
  },
  watch: {
    'user.isAuthorized' (val) {
      if (val) this.fetchDictionary(DICTIONARIES.map(o => o.apiKey))
    }
  },
  methods: {
    ...mapActions('user', ['fetchUser']),
    ...mapActions('dictionary', ['fetchDictionary'])
  },
  async created () {
    await this.fetchUser()
    await this.fetchDictionary(DICTIONARIES.map(o => o.apiKey))
    this.isReady = true
  }
}
</script>

<style lang="sass" scoped>
#app
  display: flex
  min-width: 100vw
  min-height: 100vh

  > .wrapper
    width: 100%
    padding: 0 20px
    max-width: 1280px
    margin: 0 auto

    > .header
      margin-bottom: 24px

  > .loader
    margin: auto
</style>
