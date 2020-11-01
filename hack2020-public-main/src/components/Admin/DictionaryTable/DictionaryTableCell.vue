<template lang="pug">
.dictionaryTableCellComponent(:class="{ _th: isHeader, _editing: editMode }", @click="handleClick", v-click-outside="handleClickOutside")
  template(v-if="editMode")
    ui-svg-icon(:name="editType === 'text' ? 'edit' : 'select'", :size="16").editIcon
  template(v-if="!editMode")
    .value {{ localValue }}
  template(v-else-if="editType === 'text'")
    ui-table-cell-input.input(v-model="localValue")
  template(v-else-if="editType === 'select'")
    ui-table-cell-select.select(v-model="localValue", :dictionary-key="selectOptionsDictionaryKey")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'
import UiTableCellInput from '@/components/_ui/TableCellInput'
import UiTableCellSelect from '@/components/_ui/TableCellSelect/TableCellSelect'
export default {
  components: {
    UiSvgIcon,
    UiTableCellInput,
    UiTableCellSelect
  },
  props: {
    isHeader: {
      type: Boolean,
      default: false
    },
    // none | text | select
    editType: {
      type: String,
      default: 'none'
    },
    value: {
      type: String,
      default: ''
    },
    selectOptionsDictionaryKey: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      editMode: false,
      localValue: ''
    }
  },
  watch: {
    editMode (val) {
      if (val) {
        window.addEventListener('keydown', this.handleKeyPress)
      } else {
        window.removeEventListener('keydown', this.handleKeyPress)
      }
    },
    localValue (val) {
      this.$emit('input', val)
    }
  },
  methods: {
    handleClick () {
      if (this.editType !== 'none' && !this.editMode) setTimeout(this.enableEditMode, 0)
    },
    handleClickOutside () {
      if (this.editMode) this.disableEditMode()
    },
    handleKeyPress (e) {
      if ([13, 27].includes(e.keyCode)) this.disableEditMode()
    },
    enableEditMode () {
      this.editMode = true
    },
    disableEditMode () {
      this.editMode = false
    }
  },
  created () {
    this.localValue = this.value
  }
}
</script>

<style lang="sass" scoped>
.dictionaryTableCellComponent
  border: 0.5px solid $colorLightBlue
  border-radius: 5px
  padding: 20px
  font-family: PT Sans
  font-style: normal
  font-weight: normal
  text-align: center
  font-size: 14px
  line-height: 18px
  display: flex
  align-items: center
  justify-content: center
  color: $colorBlack
  position: relative
  word-break: break-all

  &._th
    color: $colorDarkBlue
    background-color: $colorLightBlue

  &._editing
    border-color: $colorDarkBlue

  > .editIcon
    position: absolute
    top: 4px
    right: 4px
    color: $colorDarkBlue
</style>
