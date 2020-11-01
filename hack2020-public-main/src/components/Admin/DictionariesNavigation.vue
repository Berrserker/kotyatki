<template lang="pug">
.dictionariesNavigationComponent
  .group(v-for="group in groups")
    router-link.item(
      v-for="(item, index) in group",
      :to="item.route",
      :key="index",
      active-class="_active"
    ) {{ item.label }}
</template>

<script>
import { DICTIONARIES } from '@/config'
export default {
  computed: {
    items () {
      return DICTIONARIES
    },
    groups () {
      return [
        this.items.slice(0, 2),
        this.items.slice(2, 4),
        this.items.slice(4, 8),
        this.items.slice(8, 12),
        this.items.slice(12, 15)
      ]
    }
  }
}
</script>

<style lang="sass" scoped>
.dictionariesNavigationComponent
  display: flex
  background: no-repeat bottom right url('~@/assets/images/dictionaries-bg.svg')
  padding-bottom: 320px

  > .group
    &:not(:last-child)
      margin-right: 62px

    > .item
      font-family: 'PT Sans'
      font-style: normal
      font-weight: normal
      font-size: 14px
      line-height: 18px
      color: $colorGray
      text-decoration: none
      display: block
      margin: 20px 0
</style>
