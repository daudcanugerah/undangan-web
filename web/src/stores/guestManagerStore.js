import { defineStore } from 'pinia'
import { reactive } from 'vue'
import { useAuthUser } from './authStore'

export const useGuestManagerStore = defineStore('guestManager', {
  state: () => ({
    list: reactive({ data: [], total: 0 }),
  }),

  getters: {
    getGuestByTemplateID: (state) => (templateId) => {
      return state.list.filter(guest => guest.user_template_id === templateId)
    },
    getList: (state) => () => {
      return state.list
    },
  },

  actions: {
    async fetchGuests(userTemplateId, page = 1, limit = 100) {
      const authStore = useAuthUser()
      try {
        const response = await fetch(
          `http://localhost:8085/private/guests?page=${page}&limit=${limit}&user_template_id=${userTemplateId}`,
          {
            method: 'GET',
            headers: {
              'Authorization': `Bearer ${authStore.token}`,
              'Content-Type': 'application/json'
            }
          }
        )

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }

        const data = await response.json()
        this.list.data.splice(0, this.list.length)
        this.list.total = 0
        this.list.data = data?.data || []
        this.list.total = data?.total | 0
      } catch (error) {
        this.error = error.message
        console.error('Error fetching guests:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async deleteGuest(id) {
      try {
        const response = await fetch(
          `http://localhost:8085/private/guests/${id}`,
          {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${authStore.token}`,
            }
          }
        )

        if (!response.ok) {
          throw new Error('Failed to delete guest')
        }

        this.list = this.list.filter(guest => guest.id !== id)
      } catch (error) {
        console.error('Error deleting guest:', error)
        throw error
      }
    },

    async createGuest(guestData) {
      const authStore = useAuthUser()
      debugger
      try {
        const response = await fetch(
          'http://localhost:8085/private/guests',
          {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${authStore.token}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              ...guestData,
            })
          }
        )

        if (!response.ok) {
          throw new Error('Failed to create guest')
        }

        const newGuest = await response.json()
        this.list.data.push({
          ...newGuest,
        })
        await this.fetchGuests(guestData.user_template_id)
        return
      } catch (error) {
        console.error('Error creating guest:', error)
        throw error
      }
    },

    async updateGuest(id, updates) {
      try {
        const response = await fetch(
          `http://localhost:8085/private/guests/${id}`,
          {
            method: 'PUT',
            headers: {
              'Authorization': `Bearer ${localStorage.getItem('token')}`,
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              ...updates,
              tags: JSON.stringify(updates.tags || [])
            })
          }
        )

        if (!response.ok) {
          throw new Error('Failed to update guest')
        }

        const updatedGuest = await response.json()
        const index = this.list.findIndex(g => g.id === id)
        if (index !== -1) {
          this.list[index] = {
            ...this.list[index],
            ...updatedGuest,
            tags: updates.tags || this.list[index].tags
          }
        }
        return updatedGuest
      } catch (error) {
        console.error('Error updating guest:', error)
        throw error
      }
    }
  }
})
