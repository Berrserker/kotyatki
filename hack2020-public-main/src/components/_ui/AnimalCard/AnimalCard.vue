<template lang="pug">
.animalCardComponent
  .image(:style="imageStyles")
  .info
    .top
      ui-svg-icon.genderIcon(:name="`gender-${gender}`", :original-colors="true", :size="14")
      .name {{ name }}
    .text Дата поступления в приют: {{ year }}
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'

export default {
  components: {
    UiSvgIcon
  },
  props: {
    id: {
      type: Number,
      default: 0
    },
    name: {
      type: String,
      default: ''
    },
    // male | female
    gender: {
      type: String,
      default: 'male'
    },
    year: {
      type: Number,
      default: null
    }
  },
  computed: {
    imageStyles () {
      return {
        backgroundImage: `url('sample_${this.id}.jpg')`
      }
    }
  }
}
</script>

<style lang="sass" scoped>
.animalCardComponent
  background: $colorLightGray
  width: 256px
  border-radius: 5px

  > .image
    width: 256px
    height: 176px
    background: center no-repeat
    background-size: cover
    border-radius: 5px 5px 0 0

  > .info
    padding: 10px
    > .top
      display: flex
      align-items: center
      margin-bottom: 4px

      > .genderIcon
          margin-right: 9px

      > .name
        font-family: Lato
        font-style: normal
        font-weight: 600
        font-size: 16px
        line-height: 19px
        color: $colorBlack

    > .text
      padding-left: 23px
      font-family: Lato
      font-style: normal
      font-weight: 500
      font-size: 12px
      line-height: 14px
      color: $colorGray
</style>
