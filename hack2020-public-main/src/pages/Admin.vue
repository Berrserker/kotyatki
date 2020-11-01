<template lang="pug">
.adminPage
  admin-navigation.navigation
  router-view
</template>

<script>
import AdminNavigation from '@/components/Admin/Navigation'

export default {
  inject: ['user'],
  components: {
    AdminNavigation
  },
  watch: {
    'user.isAuthorized': {
      handler (val) {
        if (!val) {
          this.$router.push({
            name: 'Main'
          })
        }
      },
      immediate: true
    }
  }
}
</script>

<style lang="sass" scoped>
.adminPage
  > .navigation
    margin-bottom: 24px
</style>
