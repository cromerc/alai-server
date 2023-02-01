import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';

import Button from 'primevue/button';
import PrimeVue from 'primevue/config';
import Chart from 'primevue/chart';
import Dropdown from 'primevue/dropdown';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Ripple from 'primevue/ripple';
import StyleClass from 'primevue/styleclass';
import '@/assets/styles.scss';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(PrimeVue, { ripple: true });

app.directive('ripple', Ripple);
app.directive('styleclass', StyleClass);

app.component('Button', Button);
app.component('Chart', Chart);
app.component('Dropdown', Dropdown);
app.component('InputText', InputText);
app.component('Password', Password);
app.mount('#app');


