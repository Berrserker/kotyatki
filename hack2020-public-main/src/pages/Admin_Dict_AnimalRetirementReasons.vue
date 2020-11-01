<template lang="pug">
.adminDictAnimalRetirementReasonsPage
  dictionary-table(dictionary-key="animal_retirement_reason")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
