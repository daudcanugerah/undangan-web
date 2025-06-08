import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const useAuthUser = defineStore('auth_user', {
  state : () => ({
    activeUser: reactive({})
  }),
  persist: true,
});

// () => {
//   const activeUser = reactive({})
//   return { activeUser }
// }, {
//   }
