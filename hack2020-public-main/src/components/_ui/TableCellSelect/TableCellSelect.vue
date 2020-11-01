<template lang="pug">
.tableCellSelectComponent
  .selectedValue(@click="handleSelectedValueClick") {{ selectedValue || 'Выберите значение...' }}
  template(v-if="isOpened")
    .options
      table-cell-select-option.option(v-for="option in options", :key="option", :option="option", @click="handleOptionClick")
</template>

<script>
import { DICTIONARIES } from '@/config'

import TableCellSelectOption from './TableCellSelectOption'
export default {
  components: {
    TableCellSelectOption
  },
  props: {
    value: {
      type: String,
      default: ''
    },
    dictionaryKey: {
      type: String,
      default: null
    }
  },
  data () {
    return {
      isOpened: true,
      selectedValue: ''
    }
  },
  computed: {
    options () {
      if (!this.dictionaryKey) return []
      const filterFieldIndex = DICTIONARIES.find(o => o.key === this.dictionaryKey).filterFieldIndex || 0
      return this.$store.state.dictionary.dictionaries[this.dictionaryKey].map(item => item[filterFieldIndex])
    }
  },
  watch: {
    value () {
      this.isOpened = false
    },
    selectedValue (val) {
      this.$emit('input', val)
    }
  },
  methods: {
    handleSelectedValueClick () {
      this.toggleIsOpened()
    },
    handleOptionClick (option) {
      this.selectedValue = this.selectedValue === option ? '' : option
    },
    toggleIsOpened () {
      this.isOpened = !this.isOpened
    }
  },
  created () {
    this.selectedValue = this.value
    console.log(this.dictionaryKey)
  }
}
</script>

<style lang="sass" scoped>
.tableCellSelectComponent
  position: relative
  width: 100%

  > .value
    font-family: PT Sans
    font-style: normal
    font-weight: normal
    font-size: 14px
    line-height: 18px
    text-align: center

  > .options
    position: absolute
    top: calc(100% + 22px)
    left: -20px
    width: calc(100% + 40px)
    z-index: 1

    > .option
      &:not(:last-child)
        margin-bottom: 2px
</style>
