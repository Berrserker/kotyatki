<template lang="pug">
.adminDictAnimalBreedsPage
  dictionary-table(dictionary-key="animal_breed")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
