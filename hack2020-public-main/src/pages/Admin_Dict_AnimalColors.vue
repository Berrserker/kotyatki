<template lang="pug">
.adminDictAnimalColorsPage
  dictionary-table(dictionary-key="animal_color")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
