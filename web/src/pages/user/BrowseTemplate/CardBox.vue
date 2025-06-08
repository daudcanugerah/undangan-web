<script setup>
import SkeletonImage from '@/pages/user/BrowseTemplate/ImageCard.vue'

defineProps({
  item: {
    type: Object,
    default: '',
  },
});

const template = `
Hi! I'm interested in this template:

ID: #{{id}}  
Name: *{{name}}*  
Description: {{description}}  
Type: {{type}}  

Please let me know how to proceed.
`;

const callAdmin = (data) => {
  const filled = template
    .replace('{{id}}', data.id)
    .replace('{{name}}', data.name)
    .replace('{{description}}', data.description)
    .replace('{{type}}', data.type)

  const phone = '+6282180868730';
  const url = `https://wa.me/${phone}?text=${encodeURIComponent(filled)}`;
  window.open(url, '_blank');
}

</script>

<template>
  <Card style="width: 20rem; overflow: hidden" class="p-2">
    <template #header>
      <div class="flex justify-center items-center mb-0 h-90 overflow-hidden">
        <SkeletonImage :src="item.cover_image" alt="Example Image" width="200rem" height="100rem" />
      </div>
    </template>

    <template #title class="m-0 p-0">
      <p class="text-sm">
        {{ item.type }} #{{ item.id }}
      </p>
      <p class="text-base font-semibold">
        {{ item.name }}
      </p>

    </template>
    <template #content>
      <p class="text-sm text-gray-700">
        {{ item.description || 'No description available.' }}
      </p>

      <div class="flex gap-3 mt-2">
        <Tag v-for="(label, index) in item.tags" :severity="tagSeverity" :value="label"
          :style="{ fontSize: '0.725rem', padding: '0.1rem 0.4rem' }" />
      </div>

    </template>
    <template #footer>
      <div class="flex gap-3 mt-1 justify-between">
        <p class="mt-2 text-sm font-bold text-green-600" v-if="item.price">
          {{ item.price ? item.price.toLocaleString('id-ID', {
            style: 'currency', currency: 'IDR',
            minimumFractionDigits: 0,
            maximumFractionDigits: 0
          }) : 'Free' }}/{{
            item.price_interval }}
        </p>
        <Button label="Call Admin" @click="callAdmin(item)" icon="pi pi-phone" size="small" />
      </div>
    </template>
  </Card>
</template>
