import { defineStore } from 'pinia'
import { reactive } from 'vue'


// https://example.com/undangan/:template_id/invit/:guest_id
const templateManagerList = [
  {
    id: '1038',
    name: 'Golden Wedding Classic',
    description: 'Classic style template perfect for golden wedding invites',
    price_interval: "day",
    price: 15000,
    type: "wedding",
    tags: ['luxury', 'classic', 'timeless'],
    cover_image: "https://wevitation.com/img/slider/slide-1.jpg",
    state: 1,
    created_at: new Date('2025-05-30T19:25:00Z'),
    updated_at: new Date('2025-05-30T19:25:00Z'),
  },
  {
    id: '1039',
    name: 'Golden Wedding Deluxe',
    description: 'An elegant deluxe template for golden wedding invitations',
    price_interval: "day",
    price: 17000,
    state: 1,
    cover_image: "https://i.etsystatic.com/13562776/r/il/40a658/3012779508/il_fullxfull.3012779508_jnom.jpg",
    type: "wedding",
    tags: ['elegant', 'refined', 'sophisticated'],
    created_at: new Date('2025-05-31T10:15:00Z'),
    updated_at: new Date('2025-05-31T10:15:00Z'),
  },
  {
    id: '1040',
    name: 'Modern Minimalist',
    description: 'Simple and sleek template for modern weddings',
    price_interval: "day",
    price: 12000,
    state: 1,
    type: "wedding",
    tags: ['minimalist', 'modern', 'clean'],
    cover_image: "https://wevitation.com/img/slider/slide-2.jpg",
    created_at: new Date('2025-06-01T08:00:00Z'),
    updated_at: new Date('2025-06-01T08:00:00Z'),
  },
  {
    id: '1041',
    name: 'Rustic Charm',
    description: 'Rustic themed template with warm earthy tones',
    price_interval: "day",
    price: 14000,
    state: 1,
    type: "wedding",
    tags: ['rustic', 'warm', 'nature'],
    cover_image: "https://wevitation.com/img/slider/slide-3.jpg",
    created_at: new Date('2025-06-01T12:30:00Z'),
    updated_at: new Date('2025-06-01T12:30:00Z'),
  },
  {
    id: '1042',
    name: 'Beach Bliss Wedding Gold',
    description: 'Light and airy template for beach weddings',
    price_interval: "day",
    price: 16000,
    state: 1,
    type: "wedding",
    tags: ['beach', 'light', 'airy'],
    cover_image: "https://wevitation.com/img/slider/slide-4.jpg",
    created_at: new Date('2025-06-02T09:45:00Z'),
    updated_at: new Date('2025-06-02T09:45:00Z'),
  },
  // {
  //   id: '1042',
  //   name: 'Beach Bliss Wedding Premium',
  //   description: 'Light and airy template for beach weddings',
  //   price_interval: "day",
  //   price: 16000,
  //   state: 1,
  //   type: "wedding",
  //   tags: ['beach', 'light', 'airy'],
  //   cover_image: "https://wevitation.com/img/slider/slide-5.jpg",
  //   created_at: new Date('2025-06-02T09:45:00Z'),
  //   updated_at: new Date('2025-06-02T09:45:00Z'),
  // },
]

export const useTemplateManagerStore = defineStore('templateManager', {
  state: () => ({
    list: [], // make sure this is defined!
  }),
  getters: {
    getList(state) {
      return state.list
    }
  },
  actions: {
    fetch() {
      this.list = templateManagerList.map(e => ({
        ...e,
        gallery: ["https://wevitation.com/img/slider/slide-1.jpg", "https://wevitation.com/img/slider/slide-2.jpg", "https://wevitation.com/img/slider/slide-3.jpg", "https://wevitation.com/img/slider/slide-4.jpg"],
      }))
    },
    delete(id = "") {
      this.list = this.list.filter(user => user.id !== id);
    },
    create(data = {}) {
      this.list.push(data);
    },
    update(id, data = {}) {
      const index = this.list.findIndex(user => user.id === id);
      if (index !== -1) {
        this.list[index] = { ...this.list[index], ...data };
      }
    },
  },
});
