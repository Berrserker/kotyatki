<template lang="pug">
.adminDictShaltersPage
  dictionary-table(dictionary-key="shelter")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
