import { defineStore } from 'pinia'
import { reactive } from 'vue'

export const useAuthUser = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    activeUser: reactive({}),
  }),

  getters: {
    getActiveUser(state) {
      return () => {
        return state.activeUser
      }
    },
  },

  actions: {
    async login(email, password) {
      try {
        const response = await fetch('http://localhost:8085/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        })

        if (!response.ok) throw new Error('Login failed')

        const data = await response.json()
        if (!data.token) throw new Error('No token received')

        // Store token and fetch user profile
        this.token = data.token
        await this.fetchUserProfile()
        return true
      } catch (error) {
        console.error('Login error:', error)
        this.logout()
        throw error
      }
    },

    async fetchUserProfile() {
      if (!this.token) return

      try {
        const response = await fetch('http://localhost:8085/auth/me', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${this.token}`,
          },
        })

        if (!response.ok) throw new Error('Failed to fetch user profile')

        const user = await response.json()
        this.activeUser = user
        return user
      } catch (error) {
        console.error('Fetch profile error:', error)
        this.logout()
        throw error
      }
    },

    logout() {
      this.token = null
      this.activeUser = {}
      localStorage.removeItem('token')
    },

    // Initialize user if token exists
    async initialize() {
      if (this.token) {
        await this.fetchUserProfile()
      }
    }
  },

  persist: true,
})
