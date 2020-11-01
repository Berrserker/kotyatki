<template lang="pug">
.adminDictAnimalColorsCatPage
  dictionary-table(dictionary-key="animal_color_cat")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
