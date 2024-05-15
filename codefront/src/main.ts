import { createApp } from 'vue'
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import router from './routers';

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'


import './style.css'
import App from './App.vue'



const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

createApp(App)
  .use(pinia)
  .use(router)
  .use(ElementPlus)
  .mount('#app')
