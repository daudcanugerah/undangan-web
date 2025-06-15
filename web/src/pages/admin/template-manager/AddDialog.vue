<script setup>
import { ref, watch, computed } from 'vue';
import { useToast } from 'primevue/usetoast';
import { valibotResolver } from '@primevue/forms/resolvers/valibot';
import * as v from 'valibot';
import { useTemplateManagerStore } from '@/stores/TemplateManagerStore';

const templateManagerStore = useTemplateManagerStore()
const props = defineProps({
  initialData: {
    type: Object,
    default: () => ({
      name: '',
      description: '',
      price_interval: '',
      price: 0,
      state: 0,
      type: '',
      tags: [],
      cover_image: ''
    })
  }
});

const visible = defineModel('visible', { default: false });
const emit = defineEmits(['submit']);

const toast = useToast();


const processedTags = (tagString) => {
  console.log("Processing tags:", tagString);
  if (!tagString || tagString.length == 0) return [];
  return tagString.split(',')
    .map(tag => tag.trim())
    .filter(tag => tag.length > 0);
};

const removeTag = (tagsField, index) => {
  const currentTags = processedTags(tagsField.value);
  currentTags.splice(index, 1);
  tagsField.value = currentTags.join(', ');
};

// Options
const priceIntervals = ref([
  { name: 'Day', code: 'day' },
  { name: 'Week', code: 'week' },
  { name: 'Month', code: 'month' }
]);

const productTypes = ref(['wedding', 'birthday', 'corporate']);
const availableTags = ref(['mantap', 'sekali', 'bagus', 'keren']);

// Form schema
const schema = v.object({
  name: v.pipe(
    v.string('Name must be a string'),
    v.minLength(3, 'Name must be at least 3 characters'),
    v.maxLength(100, 'Name cannot exceed 100 characters')
  ),
  description: v.pipe(
    v.string('Description must be a string'),
    v.minLength(5, 'Description must be at least 5 characters'),
    v.maxLength(500, 'Description cannot exceed 500 characters')
  ),
  price_interval: v.pipe(
    v.string('Price interval must be a string'),
    v.minLength(1, 'Description must be at least 5 characters'),
  ),
  price: v.pipe(
    v.number('Price must be a number'),
    v.minValue(1000, 'Minimum price is 1000'),
    v.maxValue(100000000, 'Maximum price is 100,000,000')
  ),
  state: v.pipe(
    v.number('State must be a number'),
    v.includes([0, 1], 'Must be 0 (inactive) or 1 (active)')
  ),
  type: v.pipe(
    v.string('Type must be a string'),
    v.minLength(3, 'Type must be at least 3 characters'),
    v.maxLength(50, 'Type cannot exceed 50 characters')
  ),
  cover_image: v.pipe(
    v.instance(File, 'Cover image must be a file'),
    v.mimeType(['image/jpeg', 'image/png'], 'Only JPEG and PNG images allowed'),
    v.maxSize(10 * 1024 * 1024, 'Maximum file size is 10MB')
  )
});

const resolver = valibotResolver(schema);

// Form data
const formData = ref({});

watch(visible, (newVal) => {
  if (newVal) {
    formData.value = { ...props.initialData };
    previewImage.value = null;
  }
});

const headerTitle = computed(() =>
  props.initialData?.id ? 'Edit Template' : 'New Template'
);

const previewImage = ref(null);
const fileUpload = ref(null);
const onImageSelect = (event) => {
  const file = event.files[0];
  if (file) {
    fileUpload.value = file
    const reader = new FileReader();
    reader.onload = (e) => {
      previewImage.value = e.target.result;
      formData.value.cover_image = file;
    };
    reader.readAsDataURL(file);
  }
};

const onFormSubmit = async (formValues) => {
  if (!formValues.valid) {
    toast.add({
      severity: 'error',
      summary: 'Validation Error',
      detail: 'Please fix the errors in the form',
      life: 3000
    });

    return
  }

  console.log("Submitting form with values:", formValues.states);
  try {
    debugger
    await templateManagerStore.createTemplate({
      "name": formValues.states.name.value,
      "description": formValues.states.description.value,
      "price": formValues.states.price.value,
      "price_interval": formValues.states.price_interval.value,
      "state": formValues.states.state.value,
      "type": formValues.states.type.value,
      "tags": (formValues.states.tags.value || "").split(","),
      "cover_image": fileUpload.value,
    });
    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Template created successfully',
      life: 3000
    });
    // Close dialog or reset form
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: error.message || 'Failed to create template',
      life: 3000
    });
  } finally {
    //isSubmitting.value = false;
  }
};

function submitForm() {
  const formEl = formData.value.$el
  var event = new Event('submit', {
    'bubbles': true,
    'cancelable': true
  });

  formEl.dispatchEvent(event);
}

</script>

<template>
  <Dialog v-model:visible="visible" modal :header="headerTitle" :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
    class="w-full max-w-2xl" :draggable="false">
    <Form v-slot="$form" ref="formData" @submit="onFormSubmit" class="flex flex-col gap-6 p-4">
      <!-- Name -->
      <FormField v-slot="$field" name="name" class="flex flex-col gap-1">
        <label for="name">Name</label>
        <InputText id="name" name="name" type="text" placeholder="Product name"
          :class="{ 'p-invalid': $field.invalid }" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>
      <!---->
      <FormField v-slot="$field" name="description" class="flex flex-col gap-1">
        <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
        <Textarea id="description" rows="3" v-model="$field.value" :class="{ 'p-invalid': $field.invalid }"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Price Fields -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <FormField v-slot="$field" name="price" class="flex flex-col gap-1">
          <label for="price" class="block text-sm font-medium text-gray-700">Price</label>
          <InputNumber id="price" v-model="$field.value" mode="currency" currency="IDR" locale="id-ID"
            :class="{ 'p-invalid': $field.invalid }"
            class="w-full [&>input]:p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <FormField v-slot="$field" name="price_interval" class="flex flex-col gap-1">
          <label for="price_interval" class="block text-sm font-medium text-gray-700">Interval</label>
          <Dropdown id="price_interval" v-model="$field.value" :options="priceIntervals" optionLabel="code"
            optionValue="code" :class="{ 'p-invalid': $field.invalid }"
            class="w-full [&>div]:p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>
      </div>

      <!-- Status Fields -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <FormField v-slot="$field" name="state" class="flex flex-col gap-1">
          <label for="state" class="block text-sm font-medium text-gray-700">Status</label>
          <Dropdown id="state" v-model="$field.value" :options="[0, 1]" :class="{ 'p-invalid': $field.invalid }"
            class="w-full [&>div]:p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500">
            <template #option="slotProps">
              {{ slotProps.option === 1 ? 'Active' : 'Disabled' }}
            </template>
          </Dropdown>
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <FormField v-slot="$field" name="type" class="flex flex-col gap-1">
          <label for="type" class="block text-sm font-medium text-gray-700">Type</label>
          <Dropdown id="type" v-model="$field.value" :options="productTypes" :class="{ 'p-invalid': $field.invalid }"
            class="w-full [&>div]:p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>
      </div>

      <FormField v-slot="$field" name="tags" class="flex flex-col gap-1">
        <label for="tags" class="block text-sm font-medium text-gray-700">Tags</label>
        <InputText id="tags" v-model="$field.value" placeholder="Type tags, separated by commas"
          :class="{ 'p-invalid': $field.invalid }" />
        <div class="flex flex-wrap gap-2 mt-2">
          <template v-if="$field.value">
            <Chip v-for="(tag, index) in processedTags($field.value)" :key="index" :label="tag"
              @remove="removeTag($field, index)" />
          </template>
        </div>
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Image Upload -->
      <FormField v-slot="$field" name="cover_image" class="flex flex-col gap-1">
        <label for="cover_image" class="block text-sm font-medium text-gray-700">Cover Image</label>
        <FileUpload mode="basic" name="cover_image" :auto="true" accept="image/*" :maxFileSize="10000000"
          chooseLabel="Select Image" @select="onImageSelect" :class="{ 'p-invalid': $field.invalid }"
          class="w-full border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors" />
        <small class="text-gray-500 text-xs">Max file size: 10MB</small>
        <div v-if="previewImage" class="mt-2 flex justify-center">
          <img :src="previewImage" alt="Preview" class="max-h-40 rounded-lg border border-gray-200" />
        </div>
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>
    </Form>

    <template #footer>
      <div class="flex justify-end gap-3 p-4 border-t border-gray-200">
        <Button label="Cancel" icon="pi pi-times" @click="visible = false"
          class="px-4 py-2 text-gray-700 hover:bg-gray-100 rounded-lg border border-gray-300" />
        <Button label="Save" icon="pi pi-check" @click="submitForm" autofocus
          class="px-4 py-2 bg-primary-500 text-white hover:bg-primary-600 rounded-lg" />
      </div>
    </template>
  </Dialog>
</template>
