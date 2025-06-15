<template>
  <Dialog v-model:visible="visible" modal header="Add Person" :style="{ width: '30rem' }"
    :breakpoints="{ '960px': '75vw', '641px': '90vw' }">
    <!-- Form Content (Same as before) -->
    <Form v-slot="$form" ref="formData" :initialValues="initialValues" :resolver="resolver" @submit="onFormSubmit"
      class="flex flex-col gap-4">
      <!-- Name -->
      <div class="flex flex-col gap-1">
        <label for="name">Full Name</label>
        <InputText id="name" name="name" placeholder="John Doe" />
        <Message v-if="$form.name?.invalid" severity="error" size="small">
          {{ $form.name.error.message }}
        </Message>
      </div>
      <!---->
      <!-- Address -->
      <div class="flex flex-col gap-1">
        <label for="address">Address</label>
        <InputText id="address" name="address" placeholder="123 Main St, Anytown, USA" />
        <Message v-if="$form.address?.invalid" severity="error" size="small">
          {{ $form.address.error.message }}
        </Message>
      </div>
      <div class="flex flex-col gap-1">
        <label for="address">Num Attendant</label>
        <InputText id="person" name="person" type="number" />
        <Message v-if="$form.address?.invalid" severity="error" size="small">
          {{ $form.person.error.message }}
        </Message>
      </div>
      <!---->
      <div class="flex flex-col gap-1">
        <label for="telp">Phone Number</label>
        <InputMask id="telp" name="telp" mask="+62 999-9999-9999" placeholder="+62 123-4567-8901" />
        <Message v-if="$form.telp?.invalid" severity="error" size="small">
          {{ $form.telp.error.message }}
        </Message>
      </div>
      <!-- Group (Autocomplete) -->
      <div class="flex flex-col gap-1">
        <label for="group">Group</label>
        <AutoComplete id="group" name="group" :suggestions="filteredGroups" @complete="searchGroups"
          placeholder="Friends, Family, etc." :dropdown="true" :forceSelection="true" :createItem="true" />
        <Message v-if="$form.group?.invalid" severity="error" size="small">
          {{ $form.group.error.message }}
        </Message>
      </div>
      <!---->
      <!-- Tags (MultiSelect) -->
      <div class="flex flex-col gap-1">
        <label for="tags">Tags</label>
        <InputText placeholder="Type tags, separated by commas" name="tags" />
        <div class="flex flex-wrap gap-2 mt-2">
          <template v-if="$form?.tags?.value">
            <Chip v-for="(tag, index) in processedTags($form.tags.value)" :key="index" :label="tag"
              @remove="removeTag($form.tags, index)" />
          </template>
        </div>
      </div>

      <div class="flex justify-end gap-2 mt-4">
        <Button type="button" label="Cancel" severity="secondary" @click="closeDialog" />
        <Button type="submit" label="Save" />
      </div>
    </Form>
  </Dialog>
</template>

<script setup>
import { useGuestManagerStore } from '@/stores/guestManagerStore.js';
import { ref, defineModel, computed, reactive } from "vue";
import { useToast } from "primevue/usetoast";
import AutoComplete from "primevue/autocomplete";
import InputMask from "primevue/inputmask";
import Dialog from "primevue/dialog";
const processedTags = (tagString) => {
  console.log("Processing tags:", tagString);
  if (!tagString || tagString.length == 0) return [];
  return tagString.split(',')
    .map(tag => tag.trim())
    .filter(tag => tag.length > 0);
};

const guestManagerStore = useGuestManagerStore()

const removeTag = (tagsField, index) => {
  const currentTags = processedTags(tagsField.value);
  console.log(currentTags, index)
  currentTags.splice(index, 1);
  console.log(currentTags)
  tagsField.value = currentTags.join(', ');
};



const visible = defineModel("visible", { default: false });
const userTemplateID = defineModel("userTemplateID", { default: "" });
const initialData = defineModel("initialData", { default: () => ({}) });
const closeDialog = () => {
  visible.value = false;
}

const toast = useToast();

// Initialize form values (merge with initialData if provided)
const initialValues = ref({
  name: initialData?.name || "",
  address: initialData?.address || "",
  telp: initialData?.telp || "",
  group: initialData?.group || "",
  tags: initialData?.tags || [],
});

// Predefined groups and tags
const predefinedGroups = ["Friends", "Family", "Work", "Other"];
const filteredGroups = ref([]);

// Dynamic group search
const searchGroups = (event) => {
  let text = predefinedGroups.filter((group) =>
    group.toLowerCase().includes(event.query.toLowerCase())
  );

  console.log(text)

  if (text.length === 0) {
    text = [event.query.toLowerCase()]
  }
  filteredGroups.value = text
};

// Form validation
const resolver = ({ values }) => {
  const errors = {};

  if (!values.name) {
    errors.name = [{ message: "Name is required." }];
  }

  if (!values.address) {
    errors.address = [{ message: "Address is required." }];
  }

  if (!values.telp || !values.telp.startsWith("+62")) {
    errors.telp = [{ message: "Valid Indonesian phone number (+62) required." }];
  }

  return { errors };
};

// Form submission
const onFormSubmit = async (e) => {
  const { valid, states } = e
  if (valid) {
    try {
      await guestManagerStore.createGuest({
        "user_template_id": userTemplateID.value,
        "address": states.address.value,
        "group": states.group.value,
        "name": states.name.value,
        "person": Number(states.person.value) || 1,
        "tags": (states.tags.value || "").split(","),
        "telp": states.telp.value
      });
      toast.add({
        severity: "success",
        summary: "Saved!",
        detail: "Person details updated.",
        life: 3000,
      });
      visible.value = false
    } catch (error) {
      toast.add({
        severity: "error",
        summary: "Saved Error!",
        detail: "Error: " + error.message,
        life: 3000,
      });
    }

  }
};
</script>
