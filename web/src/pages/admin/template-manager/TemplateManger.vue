  <script setup>
  import time from '@/pkg/time';
  import { ref, onMounted, computed } from 'vue';
  import AddDialog from '@/pages/admin/template-manager/AddDialog.vue';

  import { useTemplateManagerStore } from '@/stores/TemplateManagerStore';
  const templateManagerStore = useTemplateManagerStore()

  const dt = ref();
  const templateDialogData = ref({})
  const deleteTemplateManagerDialog = ref(false);
  const updateTemplateManagerDialog = ref(false);

  const selectedTemplateManagers = ref([]);
  const filters = ref({
    global: { value: null, matchMode: 'contains' }
  });

  const updateTemplate = (template = {}) => {
    templateDialogData.value = template
    updateTemplateManagerDialog.value = true
  }

  const confirmDeleteTemplateManager = (user = {}) => {
    templateDialogData.value = user
    deleteTemplateManagerDialog.value = true
  }

  const templateData = computed(() => {
    return templateManagerStore.getList();
  })

  onMounted(() => {
    templateManagerStore.fetchTemplates()
  })


</script>

  <template>
    <AddDialog v-model:visible="updateTemplateManagerDialog" />

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

      <DataTable ref="dt" v-model:selection="selectedTemplateManagers" :value="templateData.data" dataKey="id"
        :paginator="true" size="small" :rows="10" :filters="filters"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5, 10, 25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
        <template #header>
          <div class="flex flex-wrap gap-2 items-center justify-between">
            <h4 class="m-0">User Manager</h4>
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
        <Column field="description" header="Description" style="min-width: 20rem"></Column>
        <Column header="Rate" sortable style="min-width: 7rem">
          <template #body="slotProps">
            {{ slotProps.data.price }} / {{ slotProps.data.price_interval }}
          </template>
        </Column>
        <Column field="type" header="Type" sortable style="min-width: 5rem"></Column>
        <Column field="state" header="Status" sortable style="min-width: 5rem">
          <template #body="slotProps">
            {{ slotProps.data.state == 1 ? "Active" : "Inactive" }}
          </template>
        </Column>
        <Column header="Tags" style="min-width: 10rem">
          <template #body="slotProps">
            <Badge v-for="tag in slotProps.data.tags" :key="tag" :value="tag" class="ml-1" severity="warn" />
          </template>
        </Column>
        <Column field="created_at" header="Created At" sortable style="min-width: 10rem">
          <template #body="slotProps">
            {{ time(slotProps.data.created_at).format("YYYY-MM-DD") }}
          </template>
        </Column>
        <Column :exportable="false" style="min-width: 10rem">
          <template #body="slotProps">
            <Button icon="pi pi-pencil" saverity="warning" outlined rounded class="mr-2"
              @click="updateTemplate(slotProps.data)" />
            <Button icon="pi pi-trash" outlined rounded severity="danger"
              @click="confirmDeleteTemplateManager(slotProps.data)" />
          </template>
        </Column>
      </DataTable>
    </div>
  </template>
