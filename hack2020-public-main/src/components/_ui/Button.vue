<template lang="pug">
.buttonComponent
  button.button(@click="handleClick")
    slot
</template>

<script>
export default {
  methods: {
    handleClick (e) {
      this.$emit('click', e)
    }
  }
}
</script>

<style lang="sass" scoped>
.buttonComponent
  width: 100%
  position: relative

  > .button
    display: block
    width: 100%
    background-color: $colorBlue
    color: $colorWhite
    font-family: 'PT Sans'
    font-style: normal
    font-weight: normal
    font-size: 22px
    line-height: 28px
    color: $colorWhite
    border: none
    padding: 14px
    border-radius: 8px

    &:active
      background-color: $colorLightBlue2
</style>
