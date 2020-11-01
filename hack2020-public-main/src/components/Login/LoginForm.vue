<template lang="pug">
.loginFormComponent
  .logo
  .title.ui-subheader-lato База животных без владельцев
  .subtitle.ui-subheader-pt-sans Вход для сотрудников
  ui-input.login(placeholder-text="Логин (e-mail)", v-model="formData.email")
  ui-input.password(placeholder-text="Пароль", type="password", v-model="formData.password")
  template(v-if="!isBusy")
    ui-button.button(@click="handleButtonClick") Вход
  template(v-else)
    .loader
      ui-loader
  //- .registerText Вы сотрудник приюта, но у вас нет аккаунта?
  //-   .link Отправить заявку
</template>

<script>
import { mapState, mapActions } from 'vuex'
import UiInput from '@/components/_ui/Input'
import UiButton from '@/components/_ui/Button'
import UiLoader from '@/components/_ui/Loader'

export default {
  components: {
    UiInput,
    UiButton,
    UiLoader
  },
  data () {
    return {
      formData: {
        email: '',
        password: ''
      }
    }
  },
  computed: {
    ...mapState({
      isBusy: state => state.user.isLoading
    })
  },
  methods: {
    ...mapActions('user', ['logIn']),
    async handleButtonClick () {
      this.logIn(this.formData)
    }
  }
}
</script>

<style lang="sass" scoped>
.loginFormComponent
  display: flex
  flex-direction: column
  align-items: center

  > .logo
    background: center no-repeat url('~@/assets/images/logo-full.svg')
    background-size: contain
    width: 80px
    height: 99px
    margin-bottom: 21px

  > .title
    margin-bottom: 21px
    text-align: center

  > .subtitle
    margin-bottom: 32px
    text-align: center

  > .login
    margin-bottom: 24px

  > .password
    margin-bottom: 45px

  > .button
    width: 256px
    // margin-bottom: 63px

  > .loader
      height: 56px
      // margin-bottom: 63px

  > .registerText
    font-family: PT Sans
    font-style: normal
    font-weight: normal
    font-size: 16px
    line-height: 21px
    color: #333333
    text-align: center

    > .link
      display: inline
      color: $colorBlue
      cursor: pointer

      &:active
        color: $colorLightBlue2

      &::before
        content: ' '
</style>
