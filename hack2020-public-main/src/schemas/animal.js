const schema = {
  fields: [
    {
      code: 'card_no',
      title: 'карточка учета животного №',
      type: 'text'
    },
    {
      code: 'name',
      title: 'Кличка',
      type: 'text'
    },
    {
      code: 'weight',
      title: 'Вес, кг',
      type: 'text'
    },
    {
      code: 'year',
      title: 'Возраст, год',
      type: 'text'
    },
    {
      code: 'shelter',
      title: 'Приют',
      type: 'select',
      model: 'shelter'
    },
    {
      code: 'kind',
      title: 'Вид',
      type: 'select',
      model: 'animal_kind'
    },
    {
      code: 'gender',
      title: 'Пол',
      type: 'select',
      model: 'animal_gender'
    },
    {
      code: 'breed',
      title: 'Порода',
      type: 'select',
      model: 'animal_breed'
    },
    {
      code: 'color',
      title: 'Окрас',
      type: 'select',
      model: 'animal_color'
    },
    {
      code: 'hair_type',
      title: 'Тип шерсти',
      type: 'select',
      model: 'animal_hair_type'
    },
    {
      code: 'ear_type',
      title: 'Тип ушей',
      type: 'select',
      model: 'animal_ear_type'
    },
    {
      code: 'tail_type',
      title: 'Тип хвоста',
      type: 'select',
      model: 'animal_tail_type'
    },
    {
      code: 'retirement_reason',
      title: 'Причина выбытия из приюта',
      type: 'select',
      model: 'animal_retirement_reason'
    },
    {
      code: 'death_cause',
      title: 'Причина смерти',
      type: 'select',
      model: 'animal_death_cause'
    },
    {
      code: 'euthanasia_cause',
      title: 'Причина эвтаназии',
      type: 'select',
      model: 'animal_euthanasia_cause'
    }
  ]
}

export default schema
