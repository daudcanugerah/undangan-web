import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { useAuthUser } from './authStore' // Import the auth store

export const useTemplateManagerStore = defineStore('templateManager', {
  state: () => ({
    list: reactive({ data: [], total: 0 }),
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

        this.list.data = []
        this.list.total = 0

        const data = await response.json()
        this.list.total = data.total
        data.data?.forEach(element => {
          this.list.data.push({ ...element, cover_image: "http://localhost:8085/" + element.cover_image })
        })
      } catch (error) {
        this.error = error.message
        console.error('Error fetching templates:', error)
        throw error // Re-throw to let components handle it
      }
    },
    async createTemplate(templateData) {
      const authStore = useAuthUser();
      if (!authStore.token) {
        throw new Error('No authentication token found');
      }

      try {
        // Create FormData for multipart/form-data
        const formData = new FormData();

        // Append all fields
        formData.append('name', templateData.name);
        formData.append('description', templateData.description);
        formData.append('price', templateData.price.toString());
        formData.append('price_interval', templateData.price_interval);
        formData.append('state', templateData.state.toString());
        formData.append('type', templateData.type);

        // Append each tag individually
        templateData.tags.forEach(tag => {
          formData.append('tags', tag);
        });

        // Append cover image if it exists
        if (templateData.cover_image instanceof File) {
          formData.append('cover_image', templateData.cover_image);
        }

        console.log(templateData)
        debugger
        const response = await fetch('http://localhost:8085/private/public-templates', {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
            // Don't set Content-Type - let browser set it with boundary
          },
          body: formData
        });

        if (!response.ok) {
          const errorData = await response.json().catch(() => ({}));
          throw new Error(errorData.message || `HTTP error! status: ${response.status}`);
        }

        const data = await response.json();

        // Optionally add the new template to your list
        this.list.data.unshift(data);
        this.list.total += 1;

        return;
      } catch (error) {
        console.error('Error creating template:', error);
        throw error; // Re-throw to let components handle it
      }
    },
  },
})
