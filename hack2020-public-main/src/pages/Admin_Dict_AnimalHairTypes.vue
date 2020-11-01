<template lang="pug">
.adminDictAnimalHairTypesPage
  dictionary-table(dictionary-key="animal_hair_type")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
