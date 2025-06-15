<script setup>
import { useRoute } from 'vue-router'
import { ref, onMounted, reactive } from "vue"
import { useAuthUser } from '@/stores/authStore';

const route = useRoute()
const activeUser = reactive({})
//const isActive = (path) => route.path === path

onMounted(() => {
  const authUser = useAuthUser()
  Object.assign(activeUser, authUser.activeUser,)
  if (activeUser.role == 2) {
    items.value = [
      {
        separator: true
      },
      {
        label: 'Template',
        items: [
          {
            label: 'Browse',
            icon: 'pi pi-search',
            url: "google.com",
            shortcut: '⌘+N',
            route: "/user/browse",
          },
          {
            label: 'My Template',
            icon: 'pi pi-file',
            shortcut: '⌘+S',
            route: "/user/my-template",
          }
        ]
      },
      {
        separator: true
      },
      {
        items: [
          {
            label: 'Logout',
            icon: 'pi pi-sign-out',
            shortcut: '⌘+Q'
          }
        ]
      },
      {
        separator: true
      }
    ]

  } else if (activeUser.role == 1) {
    items.value = [
      {
        separator: true
      },
      {
        label: 'Template',
        items: [
          {
            label: 'Template Manager',
            icon: 'pi pi-file',
            url: "google.com",
            shortcut: '⌘+N',
            route: "/admin/template-manager",
          },
          {
            label: 'User Manager',
            icon: 'pi pi-user',
            shortcut: '⌘+S',
            route: "/admin/user-manager",
          },
        ]
      },
      {
        separator: true
      },
      {
        items: [
          {
            label: 'Logout',
            icon: 'pi pi-sign-out',
            shortcut: '⌘+Q'
          }
        ]
      },
      {
        separator: true
      }
    ]
  }

});

const items = ref([]);

</script>
<template>
  <div class="card justify-center w-55 h-screen bg-white shadow-md fixed top-0 left-0 flex">
    <Menu :model="items" class="w-full md:w-55">
      <template #start>
        <span class="flex items-center justify-start gap-2 p-2 pl-10 mb-3 h-10 mt-6">
          <i class="pi pi-objects-column" />
          <span class="text-xl font-semibold"><span class="text-primary">{{ activeUser.role == 2 ? "User" : "Admin"
          }} Panel</span></span>
        </span>
        <Divider />
        <button v-ripple
          class="relative overflow-hidden w-full border-0 bg-transparent flex items-start gap-2 p-2 pl-4 hover:bg-surface-100 dark:hover:bg-surface-800 rounded-none cursor-pointer transition-colors duration-200">
          <Avatar class="mr-2 self-center min-w-12" shape="circle" :image="activeUser.profile" size="large" />
          <span class="inline-flex flex-col items-start">
            <span class="font-bold truncate" v-if="activeUser && activeUser.name">{{ activeUser?.name.slice(0,
              1000) }}</span>
            <span class="text-sm" v-if="activeUser && activeUser.name">{{ activeUser.role == 2 ? "User" : "Admin"
              }}</span>
          </span>
        </button>
      </template>
      <template #submenulabel="{ item }">
        <span class="text-primary font-bold">{{ item.label }}</span>
      </template>
      <template #item="{ item, props }">
        <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
          <a v-ripple :href="href" v-bind="props.action" @click="navigate" :class="[
            route.path === item.route ? 'bg-primary-300 text-primary font-semibold' : '',
            'flex items-center p-2 transition-colors duration-200'
          ]">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
          </a>
        </router-link>
        <a v-else v-ripple :href="item.url" :target="item.target" v-bind="props.action" :class="[
          'flex items-center p-2 transition-colors duration-200'
        ]">
          <span :class="item.icon" />
          <span class="ml-2">{{ item.label }}</span>
        </a>
      </template>
      <template #end>
      </template>
    </Menu>
  </div>
</template>
