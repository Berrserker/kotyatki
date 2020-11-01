<template lang="pug">
.adminDictionariesPage
  template(v-if="showNavigation")
    .ui-h1 Справочники
    dictionaries-navigation.navigation
  router-view
</template>

<script>
import DictionariesNavigation from '@/components/Admin/DictionariesNavigation'
export default {
  components: {
    DictionariesNavigation
  },
  computed: {
    showNavigation () {
      return this.$route.name === 'Admin_Dictionaries'
    }
  }
}
</script>
