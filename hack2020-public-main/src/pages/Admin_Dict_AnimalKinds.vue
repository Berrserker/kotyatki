<template lang="pug">
.adminDictAnimalKindsPage
  dictionary-table(dictionary-key="animal_kind")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
