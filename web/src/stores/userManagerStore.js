import { defineStore } from 'pinia'

const userManagerList = [
  {
    id: '1029',
    name: 'John Doe',
    email: 'joh@doe.com',
    created_at: new Date('2025-05-30T08:00:00Z'),
    state: 1,
    total_active_template: 1,
    total_inactive_template: 0,
  },
  {
    id: '1030',
    name: 'Jane Smith',
    email: 'jane.smith@example.com',
    created_at: new Date('2025-05-30T09:30:00Z'),
    state: 1,
    total_active_template: 1,
    total_inactive_template: 1,
  },
  {
    id: '1031',
    name: 'Alice Johnson',
    email: 'alice.j@example.net',
    created_at: new Date('2025-05-30T10:15:00Z'),
    state: 0,
    total_active_template: 0,
    total_inactive_template: 1,
  },
  {
    id: '1032',
    name: 'Bob Williams',
    email: 'bob.w@example.org',
    created_at: new Date('2025-05-30T11:45:00Z'),
    state: 1,
    total_active_template: 2,
    total_inactive_template: 1,
  },
  {
    id: '1033',
    name: 'Charlie Brown',
    email: 'charlie.brown@mail.com',
    created_at: new Date('2025-05-30T12:00:00Z'),
    state: 0,
    total_active_template: 1,
    total_inactive_template: 1,
  },
  {
    id: '1034',
    name: 'Diana Prince',
    email: 'diana.prince@hero.org',
    created_at: new Date('2025-05-30T13:20:00Z'),
    state: 1,
    total_active_template: 3,
    total_inactive_template: 1,
  },
  {
    id: '1035',
    name: 'Ethan Hunt',
    email: 'ethan.h@mission.net',
    created_at: new Date('2025-05-30T14:55:00Z'),
    state: 1,
    total_active_template: 2,
    total_inactive_template: 0,
  },
  {
    id: '1036',
    name: 'Fiona Gallagher',
    email: 'fiona.g@mail.net',
    created_at: new Date('2025-05-30T15:40:00Z'),
    state: 0,
    total_active_template: 0,
    total_inactive_template: 1,
  },
  {
    id: '1037',
    name: 'George Clooney',
    email: 'george.c@example.org',
    created_at: new Date('2025-05-30T16:10:00Z'),
    state: 1,
    total_active_template: 2,
    total_inactive_template: 1,
  },
  {
    id: '1038',
    name: 'Hannah Baker',
    email: 'hannah.baker@example.com',
    created_at: new Date('2025-05-30T19:25:00Z'),
    state: 1,
    total_active_template: 1,
    total_inactive_template: 0,
  },
]

export const useUserManagerStore = defineStore('userManager', {
  state: () => ({
    list: userManagerList, // make sure this is defined!
  }),
  getters: {
    getList(state) {
      return state.list
    }
  },
  actions: {
    fetch() {
      // fetch logic here
    },
    delete(id = "") {
      console.log(id, this.list.filter(user => user.id !== id).length)
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
