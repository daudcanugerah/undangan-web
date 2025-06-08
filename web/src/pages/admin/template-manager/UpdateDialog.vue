<script setup>
import { defineProps, computed, ref } from "vue"
import { useTemplateManagerStore } from '@/stores/templateManagerStore';
import { useToast } from 'primevue/usetoast';
import { reactive } from 'vue';
import { valibotResolver } from '@primevue/forms/resolvers/valibot';
import { yupResolver } from '@primevue/forms/resolvers/yup';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import AutoComplete from 'primevue/autocomplete';
import Dropdown from 'primevue/dropdown';
import InputNumber from 'primevue/inputnumber';
import MultiSelect from 'primevue/multiselect';
import FileUpload from 'primevue/fileupload';
import Button from 'primevue/button';
import Message from 'primevue/message';
import * as v from 'valibot';

const templateManagerStore = useTemplateManagerStore()
const priceIntervals = ref([
  { name: 'Day', code: 'day' },
  { name: 'Month', code: 'month' },
  { name: 'Week', code: 'week' },
])
const props = defineProps({
  templateData: {}
})

const formData = ref({
  name: '',
  description: '',
  price_interval: '',
  price: 0,
  state: 0,
  type: '',
  tags: [],
  cover_image: '',
});


const visible = defineModel("visible", { default: false })
const toast = useToast();

const headerTitle = computed(() =>
  props.templateData?.id !== '' ? `New Template` : 'Update Template'
);


const initialValues = reactive({
  name: '',
  description: '',
  price_interval: '',
  price: 0,
  state: 0,
  type: '',
  tags: [],
  cover_image: '',
});

const schema = v.object({
  name: v.string([v.minLength(3, 'Name must be at least 3 characters')]),
  file_name: v.string([v.minLength(3, 'File name must be at least 3 characters')]),
  description: v.string([v.minLength(5, 'Description must be at least 5 characters')]),
  price_interval: v.union([v.literal('day'), v.literal('week'), v.literal('month')]),
  price: v.number([v.minValue(1, 'Price must be positive')]),
  type: v.string([v.minLength(3, 'Type must be at least 3 characters')]),
  tags: v.array(v.string()),
  cover_image: v.string([v.minLength(5, 'Cover image URL must be valid')]),
});

const resolver = valibotResolver(schema)

const onImageSelect = (event) => {
  const file = event.files[0];
  if (file) {
    // Create preview
    const reader = new FileReader();
    reader.onload = (e) => {
      previewImage.value = e.target.result;
    };
    reader.readAsDataURL(file);

    // Update form value
    initialValues.value.cover_image = file;
  }
};

const onFormSubmit = ({ valid, originalEvent }) => {
  if (valid) {
    toast.add({ severity: 'success', summary: 'Form is submitted.', life: 3000 });
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
  <Dialog v-model:visible="visible" :modal="true" :header="headerTitle"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">
    <div class="flex flex-col gap-6">
      <Form v-slot="$form" ref="formData" :initialValues="initialValues" :resolver="resolver" @submit="onFormSubmit"
        class="flex flex-col gap-4 w-full sm:w-80">
        <!-- Name Field -->
        <FormField v-slot="$field" name="name" class="flex flex-col gap-1">
          <label for="name">Name</label>
          <InputText id="name" type="text" v-model="$field.value" placeholder="Product name"
            :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Description Field -->
        <FormField v-slot="$field" name="description" class="flex flex-col gap-1">
          <label for="description">Description</label>
          <Textarea id="description" v-model="$field.value" placeholder="Product description"
            :class="{ 'p-invalid': $field.invalid }" rows="3" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Price Interval Field -->
        <FormField v-slot="$field" name="price_interval" class="flex flex-col gap-1">
          <label for="price_interval">Price Interval</label>
          <Dropdown id="price_interval" v-model="$field.value" :options="priceIntervals" optionLabel="name"
            optionValue="code" placeholder="Select price interval" :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Price Field -->
        <FormField v-slot="$field" name="price" class="flex flex-col gap-1">
          <label for="price">Price</label>
          <InputNumber id="price" v-model="$field.value" mode="currency" currency="IDR"
            :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- State Field -->
        <FormField v-slot="$field" name="state" class="flex flex-col gap-1">
          <label for="state">State</label>
          <InputNumber id="state" v-model="$field.value" :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Type Field -->
        <FormField v-slot="$field" name="type" class="flex flex-col gap-1">
          <label for="type">Type</label>
          <Dropdown id="type" v-model="$field.value" :options="productTypes" placeholder="Select product type"
            :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Tags Field -->
        <FormField v-slot="$field" name="tags" class="flex flex-col gap-1">
          <label for="tags">Tags</label>
          <MultiSelect id="tags" v-model="$field.value" :options="availableTags" placeholder="Select tags"
            :class="{ 'p-invalid': $field.invalid }" />
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>

        <!-- Cover Image Upload -->
        <FormField v-slot="$field" name="cover_image" class="flex flex-col gap-1">
          <label for="cover_image">Cover Image</label>
          <FileUpload id="cover_image" mode="basic" name="cover_image" :auto="true" :customUpload="true"
            @select="onImageSelect" accept="image/*" :maxFileSize="10000000" chooseLabel="Select Image"
            :class="{ 'p-invalid': $field.invalid }" />
          <small class="text-gray-500">Accepted formats: JPG, PNG, GIF</small>
          <div v-if="previewImage" class="mt-2">
            <img :src="previewImage" alt="Preview" class="max-h-40 rounded-md border" />
          </div>
          <Message v-if="$field?.invalid" severity="error" size="small" variant="simple">
            {{ $field.error?.message }}
          </Message>
        </FormField>
      </Form>
    </div>
    <template #footer>
      <Button label="Cancel" icon="pi pi-times" text />
      <Button label="Save" icon="pi pi-check" @click="submitForm" />
    </template>
  </Dialog>

</template>
