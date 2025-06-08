<template>
  <div>
    <template v-if="isLoading || hasError">
      <Skeleton :width="width" :height="height" />
    </template>
    <Image v-show="!isLoading && !hasError" :src="src" :alt="alt" :width="width" :height="height" @load="onLoad"
      @error="onError" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Image from 'primevue/image';
import Skeleton from 'primevue/skeleton';

defineProps({
  src: { type: String, required: true },
  alt: { type: String, default: '' },
  width: { type: [String, Number], default: '200px' },
  height: { type: [String, Number], default: '200px' }
});

const isLoading = ref(true);
const hasError = ref(false);

function onLoad() {
  isLoading.value = false;
}

function onError() {
  isLoading.value = false;
  hasError.value = true;
}
</script>
