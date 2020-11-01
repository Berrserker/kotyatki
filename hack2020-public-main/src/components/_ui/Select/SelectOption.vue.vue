<template lang="pug">
.selectOptionComponent {{ label }}
</template>

<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    },
    value: {
      type: [String, Number],
      default: ''
    }
  }
}
</script>

<style lang="sass" scoped>
.selectOptionComponent
</style>
