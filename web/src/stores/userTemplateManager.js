import { defineStore } from 'pinia'
import time from '@/pkg/time';
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
    async createTemplate(templateData) {
      const authStore = useAuthUser();
      if (!authStore.token) {
        throw new Error('No authentication token found');
      }

      try {
        console.log("insert")
        // Create FormData for multipart/form-data
        const formData = new FormData();

        const str = JSON.stringify([{
          text: templateData.whatsapp_template.text,
          provier: "whatsapp",
        }])
        // Append all fields
        formData.append('base_template_id', "xx");
        formData.append('name', templateData.name);
        formData.append('expire_at', time(templateData.expire_at).utc().format());
        formData.append('slug', templateData.slug);
        formData.append('url', templateData.url);
        formData.append('message_template', str)

        // Append cover image if it exists
        if (templateData.cover_image instanceof File) {
          formData.append('cover_image', templateData.cover_image);
        }

        if (templateData.file instanceof File) {
          formData.append('zip_file', templateData.file);
        }
        console.log(formData)
        const response = await fetch('http://localhost:8085/private/user-templates', {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
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
    async fetchTemplates(userID = "") {
      const authStore = useAuthUser()

      try {
        const response = await fetch('http://localhost:8085/private/user-templates?page=1&limit=10000&user_id=' + userID, {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${authStore.token}`,
          }
        })

        if (!response.ok) throw new Error('Failed to fetch templates')

        const { data } = await response.json()

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
  }
})
