export const API_BASE_URL = process.env.NODE_ENV === 'production' ? 'http://46.173.216.27/' : 'http://46.173.216.27/'

export const AUTH_TOKEN_NAME = 'auth_token'

export const DICTIONARIES = [
  {
    key: 'shelter',
    apiKey: 'Приют',
    label: 'Приюты',
    addNewText: 'Добавить приют',
    route: {
      name: 'Admin_Dict_Shelters'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Оф. Краткое наименование приюта',
          type: 'text'
        },
        {
          code: 'organization',
          title: 'Подчинение (например Префектура ЮВАО или "ГБУ "Автомобильные дороги»)',
          type: 'select',
          model: 'organization'
        },
        {
          code: 'address',
          title: 'Адрес приюта',
          type: 'text'
        },
        {
          code: 'phone',
          title: 'Телефон приюта',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'organization',
    apiKey: 'Эксплуатационные организации',
    label: 'Эксплуатирующие организации',
    addNewText: 'Добавить организацию',
    route: {
      name: 'Admin_Dict_Organizations'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Наименование организации',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_kind',
    apiKey: 'Вид',
    label: 'Виды животных',
    addNewText: 'Добавить вид',
    disableAdd: true,
    route: {
      name: 'Admin_Dict_AnimalKinds'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Вид животного',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_gender',
    apiKey: 'Пол',
    label: 'Пол животных',
    addNewText: 'Добавить пол',
    disableAdd: true,
    route: {
      name: 'Admin_Dict_AnimalGenders'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Пол животного',
          type: 'text'
        }
      ]
    }
  },
  // {
  //   key: 'animal_breed',
  //   apiKey: 'poroda',
  //   label: 'Породы',
  //   addNewText: 'Добавить породу',
  //   route: {
  //     name: 'Admin_Dict_AnimalBreeds'
  //   },
  //   schema: {
  //     fields: [
  //       {
  //         code: 'title',
  //         title: 'Порода',
  //         type: 'text'
  //       }
  //     ]
  //   }
  // },
  {
    key: 'animal_breed_cat',
    apiKey: 'Породы Кошка',
    label: 'Породы кошек',
    addNewText: 'Добавить породу',
    route: {
      name: 'Admin_Dict_AnimalBreeds_Cat'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Порода кошки',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_breed_dog',
    apiKey: 'Породы Собака',
    label: 'Породы собак',
    addNewText: 'Добавить породу',
    route: {
      name: 'Admin_Dict_AnimalBreeds_Dog'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Порода собаки',
          type: 'text'
        }
      ]
    }
  },
  // {
  //   key: 'animal_color',
  //   apiKey: 'color',
  //   label: 'Окрасы',
  //   addNewText: 'Добавить окрас',
  //   route: {
  //     name: 'Admin_Dict_AnimalColors'
  //   },
  //   schema: {
  //     fields: [
  //       {
  //         code: 'title',
  //         title: 'Окрас',
  //         type: 'text'
  //       }
  //     ]
  //   }
  // },
  {
    key: 'animal_color_cat',
    apiKey: 'Окрас Кошка',
    label: 'Окрасы кошек',
    addNewText: 'Добавить окрас',
    route: {
      name: 'Admin_Dict_AnimalColors_Cat'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Окрас кошки',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_color_dog',
    apiKey: 'Окрас Собака',
    label: 'Окрасы собак',
    addNewText: 'Добавить окрас',
    route: {
      name: 'Admin_Dict_AnimalColors_Dog'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Окрас собаки',
          type: 'text'
        }
      ]
    }
  },
  // {
  //   key: 'animal_hair_type',
  //   apiKey: 'furr',
  //   label: 'Тип шерсти',
  //   addNewText: 'Добавить тип',
  //   route: {
  //     name: 'Admin_Dict_AnimalHairTypes'
  //   },
  //   schema: {
  //     fields: [
  //       {
  //         code: 'title',
  //         title: 'Тип шерсти',
  //         type: 'text'
  //       }
  //     ]
  //   }
  // },
  {
    key: 'animal_hair_type_cat',
    apiKey: 'Шерсть Кошка',
    label: 'Тип шерсти кошек',
    addNewText: 'Добавить тип',
    route: {
      name: 'Admin_Dict_AnimalHairTypes_Cat'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Тип шерсти кошек',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_hair_type_dog',
    apiKey: 'Шерсть Собака',
    label: 'Тип шерсти собак',
    addNewText: 'Добавить тип',
    route: {
      name: 'Admin_Dict_AnimalHairTypes_Dog'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Тип шерсти собаки',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_ear_type',
    apiKey: 'Тип ушей',
    label: 'Тип ушей',
    addNewText: 'Добавить тип',
    route: {
      name: 'Admin_Dict_AnimalEarTypes'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Тип ушей',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_tail_type',
    apiKey: 'Хвост',
    label: 'Тип хвостов',
    addNewText: 'Добавить тип',
    route: {
      name: 'Admin_Dict_AnimalTailTypes'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Тип хвоста',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_retirement_reason',
    apiKey: 'Причина выбытия',
    label: 'Причины выбытия из приюта',
    route: {
      name: 'Admin_Dict_AnimalRetirementReasons'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Причина выбытия из приюта',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_death_cause',
    apiKey: 'Причина смерти',
    label: 'Причины смерти',
    route: {
      name: 'Admin_Dict_AnimalDeathCauses'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Причина смерти',
          type: 'text'
        }
      ]
    }
  },
  {
    key: 'animal_euthanasia_cause',
    apiKey: 'Причина эвтаназии',
    label: 'Причины эвтаназии',
    route: {
      name: 'Admin_Dict_AnimalEuthanasiaCauses'
    },
    schema: {
      fields: [
        {
          code: 'title',
          title: 'Причина эвтаназии',
          type: 'text'
        }
      ]
    }
  }
]
