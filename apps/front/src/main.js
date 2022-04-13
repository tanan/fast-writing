import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from '@/store'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import mdiVue from 'mdi-vue/v3'
import * as mdiJs from '@mdi/js'
import PrimeVue from 'primevue/config'
import ToastService from 'primevue/toastservice';

import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import "primeflex/primeflex.css"

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .use(PrimeVue)
  .use(ToastService)
  .use(mdiVue, { icons: mdiJs })
  .mount('#app')
