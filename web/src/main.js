import "./assets/styles/main.css";
import { createPinia } from 'pinia'
import { createApp } from "vue";
import PrimeVue from "primevue/config";
import App from "./App.vue";
import Aura from "@primeuix/themes/aura";
import ToastService from 'primevue/toastservice';
import router from './router/router'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'




const pinia = createPinia()
const app = createApp(App);
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(ToastService)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: false || 'none',
    }
  }
});
app.use(router)
app.mount("#app");
