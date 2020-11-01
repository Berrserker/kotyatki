<template lang="pug">
.dictionaryTableComponent
  .controls
    .backButton(@click="handleBackButtonClick")
      ui-svg-icon.icon(name="arrow-left")
    .title {{ title }}
    .addNewButton(v-if="!disableAdd", @click="handleAddButtonClick") {{ addNewText }}
      ui-svg-icon.icon(name="plus", :size="16")
  .table._th(:style="tableStyles")
    table-cell.cell._th(v-for="item in tableSchema.fields", :value="item.title", :key="item.code")
  .table(:style="tableStyles")
    template(v-for="(row, rowIndex) in localData")
      table-cell.cell(
        v-for="(item, index) in row",
        v-model="localData[rowIndex][index]",
        :edit-type="tableSchema.fields[index].type",
        :key="`${rowIndex}_${index}`",
        :select-options-dictionary-key="tableSchema.fields[index].model"
      )
</template>

<script>
import { DICTIONARIES } from '@/config'
import UiSvgIcon from '@/components/_ui/SvgIcon'
import TableCell from './DictionaryTableCell'

export default {
  components: {
    UiSvgIcon,
    TableCell
  },
  props: {
    dictionaryKey: {
      type: String,
      default: null
    }
  },
  data () {
    return {
      title: '',
      addNewText: '',
      localData: [],
      disableAdd: false,
      tableSchema: {
        fields: []
      }
    }
  },
  computed: {
    tableStyles () {
      if (!this.tableSchema) return {}
      return {
        gridTemplateColumns: `repeat(${this.tableSchema.fields.length}, 1fr)`
      }
    }
  },
  methods: {
    handleBackButtonClick () {
      this.goToDictionaries()
    },
    goToDictionaries () {
      this.$router.push({
        name: 'Admin_Dictionaries'
      })
    },
    handleAddButtonClick () {
      this.handleAdd()
    },
    handleAdd () {
      const newRow = []
      for (let i = 0; i < this.tableSchema.fields.length; i++) {
        newRow.push('')
      }
      this.localData.push(newRow)
      setTimeout(() => { window.scrollTo(0, document.body.scrollHeight) }, 0)
    }
  },
  created () {
    this.title = DICTIONARIES.find(o => o.key === this.dictionaryKey).label
    this.addNewText = DICTIONARIES.find(o => o.key === this.dictionaryKey).addNewText
    this.localData = this.$store.state.dictionary.dictionaries[this.dictionaryKey]
    this.tableSchema = DICTIONARIES.find(o => o.key === this.dictionaryKey).schema
    this.disableAdd = DICTIONARIES.find(o => o.key === this.dictionaryKey).disableAdd || false
  }
}
</script>

<style lang="sass" scoped>
.dictionaryTableComponent
  width: 100%
  padding-bottom: 42px

  > .controls
    display: flex
    align-items: center
    margin-bottom: 25px

    > .backButton
      cursor: pointer

    > .title
      font-family: Lato
      font-style: normal
      font-weight: 600
      font-size: 16px
      line-height: 19px
      text-align: center
      color: $colorBlack
      margin: 0 auto

    > .addNewButton
      display: flex
      align-items: center
      font-family: Lato
      font-style: normal
      font-weight: 600
      font-size: 16px
      line-height: 19px
      color: $colorBlue
      cursor: pointer

      > .icon
        margin-left: 8px

  > .table
    display: grid
    grid-gap: 2px

    &._th
      margin-bottom: 2px
      position: sticky
      top: 0
      z-index: 1
</style>
