import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: () => import(/* webpackChunkName: "main" */ '../pages/Main.vue')
  },
  {
    path: '/admin',
    name: 'Admin',
    redirect: { name: 'Admin_Animals' },
    component: () => import(/* webpackChunkName: "admin" */ '../pages/Admin.vue'),
    children: [
      {
        path: 'animals',
        name: 'Admin_Animals',
        component: () => import(/* webpackChunkName: "admin_animals" */ '../pages/Admin_Animals.vue')
      },
      {
        path: 'dictionaries',
        name: 'Admin_Dictionaries',
        component: () => import(/* webpackChunkName: "admin_dictionaries" */ '../pages/Admin_Dictionaries.vue'),
        children: [
          {
            path: 'shelters',
            name: 'Admin_Dict_Shelters',
            component: () => import(/* webpackChunkName: "admin_dict_shelters" */ '../pages/Admin_Dict_Shelters.vue')
          },
          {
            path: 'organizations',
            name: 'Admin_Dict_Organizations',
            component: () => import(/* webpackChunkName: "admin_dict_organizations" */ '../pages/Admin_Dict_Organizations.vue')
          },
          {
            path: 'animal_kinds',
            name: 'Admin_Dict_AnimalKinds',
            component: () => import(/* webpackChunkName: "admin_dict_animal_kinds" */ '../pages/Admin_Dict_AnimalKinds.vue')
          },
          {
            path: 'animal_genders',
            name: 'Admin_Dict_AnimalGenders',
            component: () => import(/* webpackChunkName: "admin_dict_animal_genders" */ '../pages/Admin_Dict_AnimalGenders.vue')
          },
          // {
          //   path: 'animal_breeds',
          //   name: 'Admin_Dict_AnimalBreeds',
          //   component: () => import(/* webpackChunkName: "admin_dict_animal_breeds" */ '../pages/Admin_Dict_AnimalBreeds.vue')
          // },
          {
            path: 'animal_breeds_cat',
            name: 'Admin_Dict_AnimalBreeds_Cat',
            component: () => import(/* webpackChunkName: "admin_dict_animal_breeds_cat" */ '../pages/Admin_Dict_AnimalBreeds_Cat.vue')
          },
          {
            path: 'animal_breeds_dog',
            name: 'Admin_Dict_AnimalBreeds_Dog',
            component: () => import(/* webpackChunkName: "admin_dict_animal_breeds_dog" */ '../pages/Admin_Dict_AnimalBreeds_Dog.vue')
          },
          // {
          //   path: 'animal_colors',
          //   name: 'Admin_Dict_AnimalColors',
          //   component: () => import(/* webpackChunkName: "admin_dict_animal_colors" */ '../pages/Admin_Dict_AnimalColors.vue')
          // },
          {
            path: 'animal_colors_cat',
            name: 'Admin_Dict_AnimalColors_Cat',
            component: () => import(/* webpackChunkName: "admin_dict_animal_colors_cat" */ '../pages/Admin_Dict_AnimalColors_Cat.vue')
          },
          {
            path: 'animal_colors_dog',
            name: 'Admin_Dict_AnimalColors_Dog',
            component: () => import(/* webpackChunkName: "admin_dict_animal_colors_dog" */ '../pages/Admin_Dict_AnimalColors_Dog.vue')
          },
          // {
          //   path: 'animal_hair_types',
          //   name: 'Admin_Dict_AnimalHairTypes',
          //   component: () => import(/* webpackChunkName: "admin_dict_animal_hair_types" */ '../pages/Admin_Dict_AnimalHairTypes.vue')
          // },
          {
            path: 'animal_hair_types_cat',
            name: 'Admin_Dict_AnimalHairTypes_Cat',
            component: () => import(/* webpackChunkName: "admin_dict_animal_hair_types_cat" */ '../pages/Admin_Dict_AnimalHairTypes_Cat.vue')
          },
          {
            path: 'animal_hair_types_dog',
            name: 'Admin_Dict_AnimalHairTypes_Dog',
            component: () => import(/* webpackChunkName: "admin_dict_animal_hair_types_dog" */ '../pages/Admin_Dict_AnimalHairTypes_Dog.vue')
          },
          {
            path: 'animal_ear_types',
            name: 'Admin_Dict_AnimalEarTypes',
            component: () => import(/* webpackChunkName: "admin_dict_animal_ear_types" */ '../pages/Admin_Dict_AnimalEarTypes.vue')
          },
          {
            path: 'animal_tail_types',
            name: 'Admin_Dict_AnimalTailTypes',
            component: () => import(/* webpackChunkName: "admin_dict_animal_tail_types" */ '../pages/Admin_Dict_AnimalTailTypes.vue')
          },
          {
            path: 'animal_retirement_reasons',
            name: 'Admin_Dict_AnimalRetirementReasons',
            component: () => import(/* webpackChunkName: "admin_dict_animal_retirement_reasons" */ '../pages/Admin_Dict_AnimalRetirementReasons.vue')
          },
          {
            path: 'animal_death_causes',
            name: 'Admin_Dict_AnimalDeathCauses',
            component: () => import(/* webpackChunkName: "admin_dict_animal_death_causes" */ '../pages/Admin_Dict_AnimalDeathCauses.vue')
          },
          {
            path: 'animal_euthanasia_causes',
            name: 'Admin_Dict_AnimalEuthanasiaCauses',
            component: () => import(/* webpackChunkName: "admin_dict_animal_euthanasia_causes" */ '../pages/Admin_Dict_AnimalEuthanasiaCauses.vue')
          }
        ]
      },
      {
        path: 'reports',
        name: 'Admin_Reports',
        component: () => import(/* webpackChunkName: "admin_reports" */ '../pages/Admin_Reports.vue')
      },
      {
        path: 'users',
        name: 'Admin_Users',
        component: () => import(/* webpackChunkName: "admin_users" */ '../pages/Admin_Users.vue')
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
