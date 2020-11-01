<template lang="pug">
.tableCellInputComponent
  input.input(type="text", :value="value", ref="input", @input="handleInput")
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  methods: {
    handleInput (e) {
      this.$emit('input', e.target.value)
    }
  },
  mounted () {
    this.$refs.input.focus()
  }
}
</script>

<style lang="sass" scoped>
.tableCellInputComponent
  width: 100%

  > .input
    font-family: PT Sans
    font-style: normal
    font-weight: normal
    font-size: 14px
    line-height: 18px
    border: none
    width: 100%
</style>
