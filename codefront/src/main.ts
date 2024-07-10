import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import router from './routers';

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'


import './style.css';
import App from './App.vue';

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { faLocation, faPeopleGroup, faPrint,faRightFromBracket } from '@fortawesome/free-solid-svg-icons';

/* add icons to the library */
library.add(faLocation,faPeopleGroup,faPrint,faRightFromBracket);

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

createApp(App)
  .use(pinia)
  .use(router)
  .use(ElementPlus)
  .component('font-awesome-icon', FontAwesomeIcon)
  .mount('#app');
