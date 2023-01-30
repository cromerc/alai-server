import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';


import PrimeVue from 'primevue/config';
import Chart from 'primevue/chart';
import Dropdown from 'primevue/dropdown';
import '@/assets/styles.scss';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(PrimeVue, { ripple: true });

app.component('Chart', Chart);
app.component('Dropdown', Dropdown);

app.mount('#app');


