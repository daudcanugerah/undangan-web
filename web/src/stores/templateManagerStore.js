import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { useAuthUser } from './authStore' // Import the auth store

export const useTemplateManagerStore = defineStore('templateManager', {
  state: () => ({
    list: reactive([]),
    error: null,
    total: 0
  }),
  getters: {
    getList(state) {
      return () => {
        return state.list
      }
    }
  },
  actions: {
    async fetchTemplates() {
      const authStore = useAuthUser() // Get the auth store instance
      if (!authStore.token) {
        throw new Error('No authentication token found')
      }

      try {
        const response = await fetch('http://localhost:8085/private/public-templates', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
            'Content-Type': 'application/json'
          }
        })

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }

        const data = await response.json()
        this.list = reactive([])
        this.total = data.total
        data.data.forEach(element => {
          this.list.push(element)
        })
      } catch (error) {
        this.error = error.message
        console.error('Error fetching templates:', error)
        throw error // Re-throw to let components handle it
      }
    },

    // ... (keep your other actions: deleteTemplate, createTemplate, updateTemplate)
  }
})
