<template lang="pug">
.adminDictOrganizationsPage
  dictionary-table(dictionary-key="organization")
</template>

<script>
import DictionaryTable from '@/components/Admin/DictionaryTable/DictionaryTable'

export default {
  components: {
    DictionaryTable
  }
}
</script>
