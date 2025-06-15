<script setup>
import { ref, watch, computed } from 'vue';
import { useToast } from 'primevue/usetoast';
import { valibotResolver } from '@primevue/forms/resolvers/valibot';
import * as v from 'valibot';
import { useUserTemplateManagerStore } from '@/stores/userTemplateManager';

const templateManagerStore = useUserTemplateManagerStore()
const props = defineProps({
  initialData: {
    type: Object,
    default: () => ({
      name: '',
      expire_at: null,
      slug: '',
      url_type: 'default',
      custom_url: '',
      file: null,
      cover_image: null
    })
  }
});

const visible = defineModel('visible', { default: false });
const emit = defineEmits(['submit']);

const toast = useToast();

// Options
const urlTypes = ref([
  { name: 'Default Access', value: 'default' },
  { name: 'Custom URL', value: 'custom' }
]);

// Form schema
const schema = v.object({
  name: v.pipe(
    v.string('Name must be a string'),
    v.minLength(3, 'Name must be at least 3 characters'),
    v.maxLength(100, 'Name cannot exceed 100 characters')
  ),
  expire_at: v.pipe(
    v.date('Expiration date must be a valid date'),
  ),
  slug: v.pipe(
    v.string('Slug must be a string'),
    v.minLength(3, 'Slug must be at least 3 characters'),
    v.maxLength(50, 'Slug cannot exceed 50 characters'),
    v.regex(/^[a-z0-9-]+$/, 'Slug can only contain lowercase letters, numbers, and hyphens')
  ),
  url_type: v.pipe(
    v.string('URL type must be a string'),
    v.includes(['default', 'custom'], 'Invalid URL type')
  ),
  custom_url: v.optional(
    v.pipe(
      v.string('URL must be a string'),
      v.url('Must be a valid URL')
    )
  ),
  file: v.pipe(
    v.instance(File, 'File must be a file'),
    v.mimeType(['application/zip'], 'Only ZIP files allowed'),
    v.maxSize(50 * 1024 * 1024, 'Maximum file size is 50MB')
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
const zipFile = ref(null);

const onImageSelect = (event) => {
  const file = event.files[0];
  if (file) {
    fileUpload.value = file;
    const reader = new FileReader();
    reader.onload = (e) => {
      previewImage.value = e.target.result;
      formData.value.cover_image = file;
    };
    reader.readAsDataURL(file);
  }
};

const onZipSelect = (event) => {
  const file = event.files[0];
  if (file) {
    zipFile.value = file;
    formData.value.file = file;
    console.log("Selected ZIP file:", file);
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
    return;
  }

  try {
    const payload = {
      name: formValues.states.name.value,
      expire_at: formValues.states.expire_at.value,
      slug: formValues.states.slug.value,
      url: formValues.states.url_type.value === 'default'
        ? `http://localhost:8085/u/${formValues.states.slug.value}`
        : (formValues.states?.custom_url?.value || ""),
      file: zipFile.value,
      cover_image: fileUpload.value,
      whatsapp_template: {
        text: formValues.states.description.value
      },
    };

    await templateManagerStore.createTemplate(payload);

    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Template created successfully',
      life: 3000
    });
    visible.value = false;
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: error.message || 'Failed to create template',
      life: 3000
    });
  }
};

function submitForm() {
  const formEl = formData.value.$el;
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
        <InputText id="name" name="name" type="text" placeholder="Template name"
          :class="{ 'p-invalid': $field.invalid }" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Expiration Date -->
      <FormField v-slot="$field" name="expire_at" class="flex flex-col gap-1">
        <label for="expire_at">Expiration Date</label>
        <Calendar id="expire_at" v-model="$field.value" showIcon iconDisplay="input" dateFormat="yy-mm-dd"
          :class="{ 'p-invalid': $field.invalid }" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Slug -->
      <FormField v-slot="$field" name="slug" class="flex flex-col gap-1">
        <label for="slug">Slug</label>
        <InputText id="slug" v-model="$field.value" placeholder="unique-slug-here"
          :class="{ 'p-invalid': $field.invalid }" />
        <small class="text-gray-500">Only lowercase letters, numbers, and hyphens allowed</small>
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <FormField v-slot="$field" name="description" class="flex flex-col gap-1">
        <label for="description" class="block text-sm font-medium text-gray-700">Whatsapp Template</label>
        <Textarea id="description" rows="3" v-model="$field.value" :class="{ 'p-invalid': $field.invalid }"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- URL Type -->
      <FormField v-slot="$field" name="url_type" class="flex flex-col gap-1">
        <label for="url_type">URL Type</label>
        <SelectButton id="url_type" v-model="$field.value" :options="urlTypes" optionLabel="name" optionValue="value"
          :class="{ 'p-invalid': $field.invalid }" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Custom URL (conditionally shown) -->
      <FormField v-slot="$field" name="custom_url" class="flex flex-col gap-1"
        v-if="$form?.url_type?.value === 'custom'">
        <label for="custom_url">Custom URL</label>
        <InputText id="custom_url" name="custom_url" v-model="$field.value" placeholder="https://example.com/path"
          :class="{ 'p-invalid': $field.invalid }" />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <FormField v-slot="$field" name="default_url" class="flex flex-col gap-1"
        v-if="$form?.url_type?.value !== 'custom'">
        <label for=" custom_url">Default URL</label>
        <InputText :value="`http://localhost:8085/u/${($form?.slug?.value || '{{slug}}')}`" name="default_url"
          :class="{ 'p-invalid': $field.invalid }" disable />
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Default URL preview -->
      <div v-if="formData.url_type === 'default' && formData.slug" class="p-3 bg-gray-100 rounded-lg">
        <p class="font-medium">Default URL:</p>
        <p>http://localhost:8085/u/{{ formData.slug }}</p>
      </div>

      <!-- ZIP File Upload -->
      <FormField v-slot="$field" name="file" class="flex flex-col gap-1">
        <label for="file">ZIP File</label>
        <FileUpload mode="basic" name="file" :auto="true" accept=".zip" :maxFileSize="50000000"
          chooseLabel="Select ZIP File" @select="onZipSelect" :class="{ 'p-invalid': $field.invalid }" />
        <small class="text-gray-500">Max file size: 50MB</small>
        <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
          {{ $field.error?.message }}
        </Message>
      </FormField>

      <!-- Cover Image Upload -->
      <FormField v-slot="$field" name="cover_image" class="flex flex-col gap-1">
        <label for="cover_image">Cover Image</label>
        <FileUpload mode="basic" name="cover_image" :auto="true" accept="image/*" :maxFileSize="10000000"
          chooseLabel="Select Image" @select="onImageSelect" :class="{ 'p-invalid': $field.invalid }" />
        <small class="text-gray-500">Max file size: 10MB</small>
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
