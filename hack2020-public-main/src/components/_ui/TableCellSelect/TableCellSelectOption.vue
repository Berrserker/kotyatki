<template lang="pug">
.tableCellSelectOptionComponent(@click="handleClick") {{ option }}
</template>

<script>
export default {
  props: {
    option: {
      type: String,
      default: ''
    }
  },
  methods: {
    handleClick () {
      this.$emit('click', this.option)
    }
  }
}
</script>

<style lang="sass" scoped>
.tableCellSelectOptionComponent
  width: 100%
  background: $colorLightBlue
  border: 1px solid $colorDarkBlue
  border-radius: 5px
  padding: 3px 20px
  cursor: pointer

  &:hover
    background: $colorLightBlue3
</style>
