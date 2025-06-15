  <script setup>
  import { storeToRefs } from 'pinia'
  import time from '@/pkg/time';
  import { useToast } from 'primevue/usetoast';
  import { ref, onMounted, computed, reactive } from 'vue';
  import { useUserStore } from '@/stores/userStores';
  import { useRouter } from 'vue-router'
  import DeleteDialog from '@/pages/admin/user-manager/DeleteDialog.vue';

  const userManagerStore = useUserStore()
  const toast = useToast();
  const route = useRouter()
  const dt = ref();
  const userDialogData = ref({})

  const deleteUserManagerDialog = ref(false);

  const selectedUserManagers = ref([]);
  const filters = ref({
    global: { value: null, matchMode: 'contains' }
  });

  const updateUserTemplate = function (user) {
    route.push({
      name: 'admin.userTemplateManager',
      params: { id: user.id }
    });
  }

  const confirmDeleteUserManager = (user = {}) => {
    // userDialogData.value = user
    // deleteUserManagerDialog.value = true
  }

  const userData = computed(() => {
    return userManagerStore.getList();
  })

  onMounted(() => {
    userManagerStore.fetchUser()
  })



</script>
  <template>
    <DeleteDialog v-model:visible="deleteUserManagerDialog" :userData="userDialogData" />

    <div class="card overflow-scroll w-full p-10">
      <Toolbar class="mb-6 mt-6">
        <template #end>
          <Button label="Export" icon="pi pi-upload" severity="secondary" @click="exportCSV($event)" size="small" />
        </template>
      </Toolbar>

      <DataTable ref="dt" v-model:selection="selectedUserManagers" :value="userData.data" dataKey="id" :paginator="true"
        size="small" :rows="10" :filters="filters"
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

        <Column field="name" header="Name" sortable style="min-width: 16rem"></Column>
        <Column field="email" header="Email" sortable style="min-width: 10rem"></Column>
        <Column field="total_active_template" header="Active" sortable></Column>
        <Column field="total_inactive_template" header="Inactive" sortable></Column>
        <Column field="state" header="Status" sortable style="min-width: 5rem">
          <template #body="slotProps">
            {{ slotProps.data.is_active == 1 ? "Active" : "Inactive" }}
          </template>
        </Column>
        <Column field="created_at" header="Created At" sortable style="min-width: 10rem">
          <template #body="slotProps">
            {{ time(slotProps.data.created_at).format("YYYY-MM-DD") }}
          </template>
        </Column>
        <Column :exportable="false" style="min-width: 10rem">
          <template #body="slotProps">
            <Button icon="pi pi-receipt" saverity="warning" outlined rounded class="mr-2"
              @click="updateUserTemplate(slotProps.data)" />
          </template>
        </Column>
      </DataTable>
    </div>
  </template>
