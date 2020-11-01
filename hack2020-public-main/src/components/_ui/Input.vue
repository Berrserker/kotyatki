<template lang="pug">
.inputComponent
  input.input(
    :type="trueType",
    :placeholder="placeholderText",
    :class="`_type_${type}`",
    @input="handleInput"
  )
  template(v-if="type === 'password'")
    .passwordSwitch(@click="handlePasswordVisibilitySwitchClick")
      ui-svg-icon.icon(:name="isPasswordHidden ? 'eye-hidden' : 'eye'")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'
export default {
  components: {
    UiSvgIcon
  },
  props: {
    type: {
      type: String,
      default: 'text'
    },
    placeholderText: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      isPasswordHidden: true
    }
  },
  computed: {
    trueType () {
      if (this.type === 'password' && !this.isPasswordHidden) return 'text'
      return this.type
    }
  },
  methods: {
    handleInput (e) {
      this.$emit('input', e.target.value)
    },
    handlePasswordVisibilitySwitchClick () {
      this.switchPasswordVisibility()
    },
    switchPasswordVisibility () {
      this.isPasswordHidden = !this.isPasswordHidden
    }
  }
}
</script>

<style lang="sass" scoped>
.inputComponent
  width: 100%
  position: relative

  > .input
    display: block
    width: 100%
    background-color: $colorWhite
    border: 0.5px solid $colorGray3
    border-radius: 4px
    padding: 14px 16px
    font-family: PT Sans
    font-style: normal
    font-weight: normal
    font-size: 22px
    line-height: 28px
    color: $colorBlack

    &:focus
      background-color: $colorLightBlue

    &._type_password
      padding-right: 52px

  > .passwordSwitch
    width: 20px
    height: 20px
    position: absolute
    top: 19px
    right: 16px
    cursor: pointer
</style>
