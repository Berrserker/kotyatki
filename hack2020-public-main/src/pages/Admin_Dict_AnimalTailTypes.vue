<template lang="pug">
.adminDictAnimalTailTypesPage
  dictionary-table(dictionary-key="animal_tail_type")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
