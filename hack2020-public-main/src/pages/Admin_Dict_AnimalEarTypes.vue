<template lang="pug">
.adminDictAnimalEarTypesPage
  dictionary-table(dictionary-key="animal_ear_type")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
