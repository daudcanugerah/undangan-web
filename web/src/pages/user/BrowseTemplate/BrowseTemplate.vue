<script setup>
import CardBox from "./CardBox.vue"
import { useTemplateManagerStore } from "@/stores/templateManagerStore"
import { computed, onMounted, ref } from "vue"


let isLoading = ref(true)
const templateManagerStore = useTemplateManagerStore()
const data = computed(() => {
  if (isLoading.value) return []
  return templateManagerStore.getList()
})

onMounted(async () => {
  await templateManagerStore.fetchTemplates()
  isLoading.value = false
})

</script>

<template>
  <div class="p-5 mt-10 flex gap-5 flex-wrap justify-center">
    <CardBox v-for="item in data.data" :item="item" />
  </div>
</template>
