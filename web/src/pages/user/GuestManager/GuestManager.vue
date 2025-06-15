  <script setup>
  import { storeToRefs } from 'pinia'
  import time from '@/pkg/time';
  import { useToast } from 'primevue/usetoast';
  import { ref, onMounted, reactive, watchEffect, watch, onUnmounted, computed, onBeforeMount } from 'vue';
  import { useGuestManagerStore } from '@/stores/guestManagerStore.js';
  import { useUserTemplateManagerStore } from '@/stores/userTemplateManager.js';
  import { useRoute } from 'vue-router'
  import AddGuestManager from './AddGuestManager.vue';


  const route = useRoute()
  const guestManagerStore = useGuestManagerStore()
  const userTemplateStore = useUserTemplateManagerStore()
  const toast = useToast();
  const dt = ref()
  const userTemplateData = ref({})

  const selectedUserManagers = ref([])
  const filters = ref({
    global: { value: null, matchMode: 'contains' }
  });


  const gustManagerList = computed(() => {
    if (isLoading.value) return [];
    return guestManagerStore.getList();
  });

  const isAddModalVisible = ref(false);
  let isLoading = ref(true)
  const userTemplateID = ref("")

  onMounted(async () => {
    isLoading.value = false
    const templateID = route.params.id;

    if (!templateID) {
      return
    }

    await guestManagerStore.fetchGuests(templateID)
    await userTemplateStore.fetchTemplates()

    userTemplateData.value = userTemplateStore.getTemplate(templateID);
    userTemplateID.value = templateID
  })


  const renderTemplate = (text, guest, url) => {
    return text.replace(/{{name}}/gi, guest.name)
      .replace(/{{address}}/gi, guest.address)
      .replace(/{{url}}/gi, url)
  }

  function getMenu(guest) {
    return [
      {
        label: 'Open Website',
        icon: 'pi pi-globe',
        command: () => {
          console.log(`${userTemplateData.value.url}?=guset_id=${guest.id}`)
          window.open(`${userTemplateData.value.url}?=guset_id=${guest.id}`, '_blank', 'noopener,noreferrer')
        }
      },
      {
        separator: true
      },
      {
        label: 'Send Via Whatsapp',
        icon: 'pi pi-whatsapp',
        command: () => {
          const text = renderTemplate(userTemplateData.value.message_template["whatsapp"].text, guest, userTemplateData.value.url)
          const targetURL = `https://api.whatsapp.com/send?phone=${guest.telp}&text=${encodeURIComponent(text)}`
          window.open(targetURL, '_blank', 'noopener,noreferrer')
        }
      },
      {
        label: 'Copy Whatsapp Message',
        icon: 'pi pi-whatsapp pi-copy',
        command: async () => {

          navigator.clipboard.writeText(renderTemplate(userTemplateData.value.message_template["whatsapp"].text, guest, userTemplateData.value.url)).then(() => {
            toast.add({ severity: 'success', summary: 'Text Copied', life: 3000 })
          }).catch(err => {
            toast.add({ severity: 'error', summary: 'Text Copied Error ' + err.Error(), life: 3000 })
          })
        }
      }
    ]
  }


</script>
  <template>
    <Toast />
    <AddGuestManager v-model:visible="isAddModalVisible" v-model:userTemplateID="userTemplateID" />
    <div class="card overflow-scroll w-full p-10">
      <Toolbar class="mb-6 mt-6">
        <template #start>
          <Button label="Add Guest" icon="pi pi-plus" severity="secondary" class="mr-2" size="small"
            @click="isAddModalVisible = true" />
        </template>

        <template #end>
          <Button label="Export" icon="pi pi-upload" severity="secondary" @click="exportCSV($event)" size="small" />
        </template>
      </Toolbar>
      <DataTable ref="dt" v-model:selection="selectedUserManagers" :value="gustManagerList.data" dataKey="id"
        :paginator="true" size="small" :rows="10" :filters="filters"
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
        <Column field="name" header="Name" sortable style="min-width: 10rem"></Column>
        <Column field="address" header="Address" sortable style="min-width: 12rem"></Column>
        <Column field="telp" header="Telp" sortable style="min-width: 15rem" />
        <Column field="group" header="Group" sortable style="min-width: 5rem" />
        <Column field="tags" header="Tags" sortable style="min-width: 5rem">
          <template #body="slotProps">
            {{ (slotProps?.data.tags || []).join(",") }}
          </template>
        </Column>
        <Column field="view_at" header="Last View" sortable>
          <template #body="slotProps">
            {{ slotProps.data.view_at ? time(slotProps.data.view_at).format("YYYY-MM-DD HH:mm") : "none" }}
          </template>
        </Column>
        <Column field="created_at" header="Created" sortable>
          <template #body="slotProps">
            {{ time(slotProps.data.created_at).format("YYYY-MM-DD") }}
          </template>
        </Column>
        <Column field="attend" header="Attend?" sortable style="min-width: 5rem">
          <template #body="slotProps">
            ({{ slotProps.data.person }}) {{ slotProps.data.attend ? "Yes" : "No" }}
          </template>
        </Column>
        <Column field="message" header="Message" sortable style="min-width: 8rem" />
        <Column :exportable="false" style="min-width: 5rem">
          <template #body="slotProps">
            <div class="flex flex-row gap-2">
              <SplitButton label="" icon="pi pi-cog" :model="getMenu(slotProps.data)" size="small" class="m-1" />
              <Button icon="pi pi-trash" outlined rounded severity="danger"
                @click="confirmDeleteUserManager(slotProps.data)" />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>
  </template>
