import { defineStore } from 'pinia'
import { reactive } from 'vue'
import { useAuthUser } from './authStore'

export const useUserTemplateManagerStore = defineStore('userTemplateManager', {
  state: () => ({
    list: reactive({ data: [], total: 0 }), // This will store our templates
  }),
  getters: {
    getList(state) {
      return () => {
        return state.list
      }
    },
    getTemplate(state) {
      return (id) => {
        return state.list.data.filter(e => e.id == id)[0]
      }
    },
  },
  actions: {
    async fetchTemplates() {
      const authStore = useAuthUser()

      try {
        const response = await fetch('http://localhost:8085/private/user-templates?page=1&limit=10000', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
          }
        })

        if (!response.ok) throw new Error('Failed to fetch templates')

        const { data } = await response.json()
        console.log(data)

        this.list.data = []
        this.list.total = 0

        data.forEach(e => {
          this.list.data.push(e)
        }) // Store the fetched templates in the list
      } catch (error) {
        console.error('Error fetching templates:', error)
        throw error
      }
    },

    // Basic CRUD operations
    addTemplate(template) {
      this.list.push(template)
    },
    removeTemplate(id) {
      this.list = this.list.filter(t => t.id !== id)
    },

    updateTemplate(updatedTemplate) {
      const index = this.list.findIndex(t => t.id === updatedTemplate.id)
      if (index !== -1) {
        this.list[index] = updatedTemplate
      }
    }
  }
})
