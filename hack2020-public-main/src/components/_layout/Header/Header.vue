<template lang="pug">
.headerComponent
  router-link.logo(:to="{ name: 'Main' }")
  .text База животных без владельцев
  .auth
    template(v-if="!user.isAuthorized")
      .loginButton(@click="openLoginModal") Войти
    template(v-else)
      //- .login Вошли как admin.
      router-link.adminLink(:to="{ name: 'Admin' }", active-class="_active").adminLink Перейти в панель администрирования
      .logoutButton(@click="logOut") Выйти
</template>

<script>
import { mapActions } from 'vuex'
import UiButton from '@/components/_ui/Button'
export default {
  inject: ['user'],
  components: {
    UiButton
  },
  methods: {
    ...mapActions('ui', ['openLoginModal']),
    ...mapActions('user', ['logOut'])
  }
}
</script>

<style lang="sass" scoped>
.headerComponent
  display: flex
  align-items: center
  padding: 20px 0
  border-bottom: 1px solid $colorGray2

  > .logo
    width: 118px
    height: 24px
    background: center no-repeat url('~@/assets/images/logo-line.svg')
    background-size: contain
    margin-right: 12px
    display: block
    outline: none

  > .text
    font-family: Lato
    font-style: normal
    font-weight: 500
    font-size: 14px
    line-height: 17px
    color: #C92723

  > .auth
    margin-left: auto
    display: flex
    align-items: center

    > .loginButton
      cursor: pointer
      font-family: Lato
      font-style: normal
      font-weight: 600
      font-size: 14px
      line-height: 19px

    > .adminLink
      color: $colorBlue
      display: block
      margin-right: 6px
      font-family: Lato
      font-style: normal
      font-weight: 600
      font-size: 16px
      line-height: 19px
      color: $colorRed
      text-decoration: none

      &._active
        display: none

    > .logoutButton
      cursor: pointer
      font-family: Lato
      font-style: normal
      font-weight: 600
      font-size: 14px
      line-height: 19px
</style>
