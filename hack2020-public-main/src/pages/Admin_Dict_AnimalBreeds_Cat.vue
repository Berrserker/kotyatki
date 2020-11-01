<template lang="pug">
.adminDictAnimalBreedsCatPage
  dictionary-table(dictionary-key="animal_breed_cat")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
