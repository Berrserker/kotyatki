<template lang="pug">
.selectComponent
  select-option(v-for="(option, index) in options", :key="index")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'
import SelectOption from './SelectOption'
export default {
  components: {
    UiSvgIcon,
    SelectOption
  },
  props: {
    type: {
      type: String,
      default: 'text'
    },
    placeholderText: {
      type: String,
      default: ''
    },
    options: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style lang="sass" scoped>
.selectComponent
  width: 100%
  position: relative
  user-select: none

  > .select
</style>
