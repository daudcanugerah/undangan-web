  <script setup>
  import time from '@/pkg/time';
  import { useRouter } from 'vue-router';
  import { ref, onMounted, computed } from 'vue';
  import { useUserTemplateManagerStore } from '@/stores/userTemplateManager';
  const userTemplateStore = useUserTemplateManagerStore()

  const router = useRouter();
  const dt = ref();

  const list = computed(() => {
    if (isLoading.value) return [];
    return userTemplateStore.getList;
  });

  const selectedUserManagers = ref([]);
  const filters = ref({
    global: { value: null, matchMode: 'contains' }
  });

  function toGuestManager(data) {
    router.push({
      name: 'user.guestManager',
      params: { id: data.id }
    });

  }

  const isLoading = ref(true);

  onMounted(async () => {
    userTemplateStore.fetch()
    isLoading.value = false;
  });

</script>


  <template>
    <div class="card overflow-scroll w-full p-10">
      <DataTable ref="dt" v-model:selection="selectedUserManagers" :value="list" dataKey="id" :paginator="true"
        size="small" :rows="10" :filters="filters"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5, 10, 25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
        <template #header>
          <div class="flex flex-wrap gap-2 items-center justify-between">
            <h4 class="m-0">My Template</h4>
            <IconField>
              <InputIcon size="small">
                <i class="pi pi-search" />
              </InputIcon>
              <InputText size="small" v-model="filters['global'].value" placeholder="Search..." />
            </IconField>
          </div>
        </template>
        <Column field="name" header="Name" sortable style="min-width: 16rem"></Column>
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
            <Button icon="pi pi-user" saverity="warning" outlined rounded class="mr-2"
              @click="toGuestManager(slotProps.data)" />
            <Button icon="pi pi-trash" outlined rounded severity="danger"
              @click="confirmDeleteUserManager(slotProps.data)" />
          </template>
        </Column>
      </DataTable>
    </div>
  </template>
