<template lang="pug">
.adminReportsPage
  .ui-h1 Отчеты
  .text.ui-h2-lato Здесь будут отчеты.
  ui-svg-icon.icon(name="paw", :size="100")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'
export default {
  components: {
    UiSvgIcon
  }
}
</script>

<style lang="sass" scoped>
.text
  margin-top: 120px
  text-align: center
  margin-bottom: 20px

.icon
  margin: 0 auto
  color: $colorGray2
</style>
