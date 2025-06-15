  <script setup>
  import time from '@/pkg/time';
  import { ref, onMounted, computed } from 'vue';
  import AddDialog from '@/pages/admin/user-template/AddDialog.vue';
  import { useRoute } from 'vue-router'

  const route = useRoute()
  import { useUserTemplateManagerStore } from '@/stores/userTemplateManager';
  const templateManagerStore = useUserTemplateManagerStore()

  const dt = ref();
  const templateDialogData = ref({})
  const deleteUserTemplateDialog = ref(false);
  const updateUserTemplateDialog = ref(true);

  const selectedUserTemplates = ref([]);
  const filters = ref({
    global: { value: null, matchMode: 'contains' }
  });

  const updateTemplate = (template = {}) => {
    templateDialogData.value = template
    updateUserTemplateDialog.value = true
  }

  const confirmDeleteUserTemplate = (user = {}) => {
    templateDialogData.value = user
    deleteUserTemplateDialog.value = true
  }

  const templateData = computed(() => {
    return templateManagerStore.getList();
  })

  onMounted(() => {
    const userID = route.params.id;
    templateManagerStore.fetchTemplates(userID)
  })


</script>

  <template>
    <AddDialog v-model:visible="updateUserTemplateDialog" />

    <div class="card overflow-scroll w-full p-10">
      <Toolbar class="mb-6 mt-6">
        <template #start>
          <Button label="New" icon="pi pi-plus" severity="secondary" class="mr-2" size="small"
            @click="updateTemplate" />
        </template>

        <template #end>
          <Button label="Export" icon="pi pi-upload" severity="secondary" @click="exportCSV($event)" size="small" />
        </template>
      </Toolbar>

      <DataTable ref="dt" v-model:selection="selectedUserTemplates" :value="templateData.data" dataKey="id"
        :paginator="true" size="small" :rows="10" :filters="filters"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5, 10, 25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
        <template #header>
          <div class="flex flex-wrap gap-2 items-center justify-between">
            <h4 class="m-0">User Template Manager</h4>
            <IconField>
              <InputIcon size="small">
                <i class="pi pi-search" />
              </InputIcon>
              <InputText size="small" v-model="filters['global'].value" placeholder="Search..." />
            </IconField>
          </div>
        </template>

        <Column field="id" header="ID" sortable style="min-width: 2rem"></Column>
        <Column field="name" header="Name" sortable style="min-width: 10rem"></Column>
        <Column field="slug" header="Slug" style="min-width: 10rem"></Column>
        <Column field="type" header="Type" sortable style="min-width: 5rem"></Column>
        <Column field="state" header="Status" sortable style="min-width: 5rem">
          <template #body="slotProps">
            {{ slotProps.data.state == 1 ? "Active" : "Inactive" }}
          </template>
        </Column>
        <Column field="created_at" header="Created At" sortable style="min-width: 10rem">
          <template #body="slotProps">
            {{ time(slotProps.data.created_at).format("YYYY-MM-DD") }}
          </template>
        </Column>
        <Column field="expire_at" header="Expire At" sortable style="min-width: 10rem">
          <template #body="slotProps">
            {{ time(slotProps.data.expire_at).format("YYYY-MM-DD") }}
          </template>
        </Column>
        <Column :exportable="false" style="min-width: 10rem">
          <template #body="slotProps">
            <Button icon="pi pi-trash" outlined rounded severity="danger"
              @click="confirmDeleteUserTemplate(slotProps.data)" />
          </template>
        </Column>
      </DataTable>
    </div>
  </template>
