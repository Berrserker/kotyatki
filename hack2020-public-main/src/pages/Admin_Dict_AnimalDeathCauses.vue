<template lang="pug">
.adminDictAnimalDeathCausesPage
  dictionary-table(dictionary-key="animal_death_cause")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
