import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from '@/store'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import mdiVue from 'mdi-vue/v3';
import * as mdiJs from '@mdi/js';

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .use(mdiVue, { icons: mdiJs })
  .mount('#app')
