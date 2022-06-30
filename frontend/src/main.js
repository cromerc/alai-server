import { createApp } from 'vue';
import App from './App.vue';
import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/mdc-light-indigo/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Menubar from 'primevue/menubar';
import InputText from 'primevue/inputtext';

const app = createApp(App);

app.use(PrimeVue);
app.component('MenuBar', Menubar);
app.component('InputText', InputText);

app.mount('#app');
