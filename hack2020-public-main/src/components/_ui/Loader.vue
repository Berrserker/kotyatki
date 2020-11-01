<template lang="pug">
.loaderComponent
  ui-svg-icon.icon(name="paw", :size="size")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'

export default {
  components: {
    UiSvgIcon
  },
  props: {
    size: {
      type: Number,
      default: 56
    }
  }
}
</script>

<style lang="sass" scoped>
.icon
  animation: 3s breathe linear infinite
  color: $colorBlue

@keyframes breathe
  0%
    transform: scale(1)
  50%
    transform: scale(.8)
  100%
    transform: scale(1)
</style>
