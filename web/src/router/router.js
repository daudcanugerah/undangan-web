import { createWebHistory, createRouter } from 'vue-router'

import UserManagerView from '@/pages/admin/user-manager/UserManger.vue'
import MyTemplateView from '@/pages/user/MyTemplate/MyTemplate.vue'
import TemplateMangerView from '@/pages/admin/template-manager/TemplateManger.vue'
import GuestManagerView from '@/pages/user/GuestManager/GuestManager.vue'
import BrowseTemplateView from '@/pages/user/BrowseTemplate/BrowseTemplate.vue'
import Login from '@/pages/Login.vue'
import Layout from "@/components/Layout.vue"


const adminRoutes = [
  { path: '/admin/login', component: Login, name: "admin.login" },
  { path: '/user/login', component: Login, name: "user.login" },
  { path: '/user/login', component: Login, name: "user.login" },
  {
    component: Layout, path: '/user', children: [
      { path: '/user/browse', component: BrowseTemplateView, name: "user.browse" },
      { path: '/user/my-template', component: MyTemplateView, name: "user.myTemplate" },
      { path: '/user/guest-manager/:id', component: GuestManagerView, name: "user.guestManager" },
    ]
  },
  {
    component: Layout, path: '/admin', children: [
      { path: '/admin/user-manager', component: UserManagerView, name: "admin.userManager" },
      { path: '/admin/template-manager', component: TemplateMangerView, name: "admin.templateManager" },
      { path: '/admin/user-template-manager/:user_id', component: UserManagerView, name: "admin.userTemplateManager" },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: adminRoutes,
})

export default router
