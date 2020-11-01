<template lang="pug">
.adminDictAnimalGendersPage
  dictionary-table(dictionary-key="animal_gender")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
