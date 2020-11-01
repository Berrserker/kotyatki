<template lang="pug">
.loginModalComponent
  .inner(:class="{ _closing: isClosing }")
    .close(@click="close")
      ui-svg-icon(name="close", :size="14")
    login-form
</template>

<script>
import { mapActions } from 'vuex'
import LoginForm from '@/components/Login/LoginForm'
import UiSvgIcon from '@/components/_ui/SvgIcon'
export default {
  inject: ['user'],
  components: {
    LoginForm,
    UiSvgIcon
  },
  data () {
    return {
      isClosing: false
    }
  },
  watch: {
    'user.isAuthorized' (val) {
      if (val) this.close()
    }
  },
  methods: {
    ...mapActions('ui', ['closeLoginModal']),
    handleEscKeyPress (e) {
      if (e.keyCode === 27) this.close()
    },
    close () {
      this.isClosing = true
      setTimeout(this.closeLoginModal, 300)
    }
  },
  created () {
    window.addEventListener('keydown', this.handleEscKeyPress)
  },
  beforeDestroy () {
    window.removeEventListener('keydown', this.handleEscKeyPress)
  }
}
</script>

<style lang="sass" scoped>
@keyframes slideDown
  0%
    top: -1000px
  100%
    top: 0
@keyframes slideUp
  0%
    top: 0
  100%
    top: -1000px

.loginModalComponent
  position: fixed
  top: 0
  left: 0
  width: 100vw
  height: 100vh
  background: rgba(#000, .24)
  display: flex
  flex-direction: column
  padding: 42px
  overflow: auto

  > .inner
    margin: auto
    background-color: $colorWhite
    border-radius: 4px
    padding: 32px 48px
    position: relative
    max-width: 480px
    animation: slideDown .3s ease-out

    > .close
      position: absolute
      top: 17px
      right: 17px
      cursor: pointer

    &._closing
      animation: slideUp .3s ease-out
</style>
