<template lang="pug">
.adminDictAnimalEuthanasiaCausesPage
  dictionary-table(dictionary-key="animal_euthanasia_cause")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
