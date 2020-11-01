<template lang="pug">
.adminNavigationComponent
  router-link.item(
    v-for="(item, index) in items",
    :to="item.route",
    :key="index",
    active-class="_active"
  ) {{ item.label }}
</template>

<script>
const ADMIN_NAVIGATION_ITEMS = [
  {
    label: 'Реестр',
    route: {
      name: 'Admin_Animals'
    }
  },
  {
    label: 'Справочники',
    route: {
      name: 'Admin_Dictionaries'
    }
  },
  {
    label: 'Отчеты',
    route: {
      name: 'Admin_Reports'
    }
  }
  // {
  //   label: 'Пользователи',
  //   route: {
  //     name: 'Admin_Users'
  //   }
  // }
]

export default {
  computed: {
    items () {
      return ADMIN_NAVIGATION_ITEMS
    }
  }
}
</script>

<style lang="sass" scoped>
.adminNavigationComponent
  display: flex
  align-items: center
  border-bottom: 1px solid $colorGray2

  > .item
    font-family: Lato
    font-style: normal
    font-weight: 600
    font-size: 16px
    line-height: 19px
    color: $colorBlack
    padding: 10px 0
    border-bottom: 2px solid transparent
    margin-bottom: -1px
    text-decoration: none
    display: block
    outline: none

    &:not(:last-child)
      margin-right: 44px

    &._active
      border-bottom-color: $colorRed
</style>
