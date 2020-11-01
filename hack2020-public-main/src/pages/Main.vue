<template lang="pug">
.mainPage
  .ui-h1.title Поиск лучшего друга
  .cards
    ui-animal-card(v-for="card in cards", :id="card.id", :name="card.name", :year="card.year", :gender="card.gender", :key="card.id")
  .text.ui-h2-lato Здесь будет поиск с фильтрами.
  ui-svg-icon.icon(name="paw", :size="100")
</template>

<script>
import UiSvgIcon from '@/components/_ui/SvgIcon'
import UiAnimalCard from '@/components/_ui/AnimalCard/AnimalCard'
export default {
  components: {
    UiSvgIcon,
    UiAnimalCard
  },
  data () {
    return {
      cards: [
        {
          id: 1,
          name: 'Барсик',
          gender: 'male',
          year: 2020
        },
        {
          id: 2,
          name: 'Ласка',
          gender: 'female',
          year: 2020
        },
        {
          id: 3,
          name: 'Чак',
          gender: 'male',
          year: 2020
        },
        {
          id: 4,
          name: 'Джоли',
          gender: 'female',
          year: 2020
        }
      ]
    }
  }
}
</script>

<style lang="sass" scoped>
.mainPage
  > .title
    margin-bottom: 20px

  > .cards
    display: flex
    justify-content: space-between
  > .text
    margin-top: 120px
    text-align: center
    margin-bottom: 20px

  > .icon
    margin: 0 auto
    color: $colorGray2
</style>
