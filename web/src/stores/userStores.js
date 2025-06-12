import { defineStore } from 'pinia'
import { reactive } from 'vue'

const userList = [
  {
    id: 2,
    username: "user",
    name: "Muhamad Daud Anugerah",
    password: "user",
    email: "daudcanugerah@gmail.com",
    role: 2,
    profile: "https://w7.pngwing.com/pngs/340/946/png-transparent-avatar-user-computer-icons-software-developer-avatar-child-face-heroes-thumbnail.png",
    created_at: new Date('2025-05-30T19:25:00Z'),
    updated_at: new Date('2025-05-30T19:25:00Z'),
  },
  {
    id: 1,
    username: "admin",
    name: "Verlyna Zahra",
    password: "admin",
    email: "admin@gmail.com",
    role: 1,
    profile: "https://e7.pngegg.com/pngimages/799/987/png-clipart-computer-icons-avatar-icon-design-avatar-heroes-computer-wallpaper-thumbnail.png",
    created_at: new Date('2025-05-30T19:25:00Z'),
    updated_at: new Date('2025-05-30T19:25:00Z'),
  },
]

export const useUserStore = defineStore('user', {
  state: () => ({
    activeUser: {},
    list: userList.map(e => ({
      ...e,
    })), // make sure this is defined!
  }),
  getters: {
    findByCredential(state) {
      return (username, password, role) => {
        let data = (state.list.find(user => user.username === username && user.password === password && user.role == role) || {})
        delete data.password
        return reactive(data)
      }
    },
    getActiveUser(state) {
      return () => {
        return reactive(state.activeUser)
      }
    }
  },
  actions: {
    setActiveUser(user) {
      this.activeUser = user;
    },
    fetch() {
      // fetch logic here
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
