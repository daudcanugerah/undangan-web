<script setup>
import { defineProps } from "vue"
import { useTemplateManagerStore } from '@/stores/templateManagerStore';

import { useToast } from 'primevue/usetoast';

const userManagerStore = useTemplateManagerStore()
const visible = defineModel("visible", { default: false })
const toast = useToast();


const props = defineProps({
  userData: {}
})

async function deleteTemplateManager() {
  userManagerStore.delete(props.userData.id)
  toast.add({
    id: 'delete-user',
    severity: 'success',
    summary: 'Success',
    detail: `Template ${props.userData.name} deleted successfully.`,
    life: 3000
  })
  visible.value = false
}
</script>

<template>
  <Toast />
  <Dialog :visible="visible" :style="{ width: '450px' }" header="Confirm" :modal="true">
    <div class="flex items-center gap-4">
      <i class="pi pi-exclamation-triangle !text-3xl" />
      <span v-if="userData">Are you sure you want to delete <b>{{ userData.name }}</b>?</span>
    </div>
    <template #footer>
      <Button label="No" icon="pi pi-times" text @click="visible = false" />
      <Button label="Yes" icon="pi pi-check" @click="deleteTemplateManager" />
    </template>
  </Dialog>
</template>
