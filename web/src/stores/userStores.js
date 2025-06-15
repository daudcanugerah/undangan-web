import { defineStore } from 'pinia'
import { useAuthUser } from './authStore' // Import the auth store
import { reactive } from 'vue'


export const useUserStore = defineStore('user', {
  state: () => ({
    list: reactive({ data: [], total: 0 }), // make sure this is defined!
  }),
  getters: {
    getList(state) {
      return () => {
        return state.list
      }
    }
  },
  actions: {
    async fetchUser(page = 1, limit = 100) {
      const authStore = useAuthUser() // Get the auth store instance
      if (!authStore.token) {
        throw new Error('No authentication token found')
      }
      const url = `http://localhost:8085/private/users?page=${page}&limit=${limit}`;
      try {
        const response = await fetch(url, {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
            'Content-Type': 'application/json'
          }
        });

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        this.list.data = []
        this.list.total = 0

        const data = await response.json()
        this.list.total = data.total
        data.data?.forEach(element => {
          console.log(element)
          this.list.data.push(element)
        })

        return data;
      } catch (error) {
        console.error('Error fetching users:', error);
        throw error;
      }
    }
  },
});
