<template lang="pug">
.svgIconComponent(:style="style", v-html="rawSvg", :class="{ _inheritColors: !originalColors}")
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      default: 'default'
    },
    size: {
      type: Number,
      default: 20
    },
    originalColors: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    rawSvg () {
      let rawSvg = ''
      try {
        rawSvg = require(`@/assets/icons/${this.name}.svg?raw`)
      } catch (e) {}
      return rawSvg
    },
    style () {
      return {
        width: `${this.size}px`,
        height: `${this.size}px`
      }
    }
  }
}
</script>

<style lang="sass" scoped>
.svgIconComponent
  line-height: 0
  &._inheritColors
    /deep/ svg
      fill: currentColor
      width: 100%
      height: 100%
      > path
        fill: inherit
</style>
