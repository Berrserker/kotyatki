<template lang="pug">
.adminDictAnimalBreedsDogPage
  dictionary-table(dictionary-key="animal_breed_dog")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
