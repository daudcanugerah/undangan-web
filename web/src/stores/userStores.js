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
    profile: "https://yt3.ggpht.com/ytc/AIdro_kPHVv90nVAdc5jKMJO6YMmT0qQy2Sr0u-8gxOxmbQ=s88-c-k-c0x00ffffff-no-rj",
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
    profile: "https://yt3.ggpht.com/ytc/AIdro_kPHVv90nVAdc5jKMJO6YMmT0qQy2Sr0u-8gxOxmbQ=s88-c-k-c0x00ffffff-no-rj",
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
