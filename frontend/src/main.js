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
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';
import Toolbar from 'primevue/toolbar';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import Textarea from 'primevue/textarea';
import RadioButton from 'primevue/radiobutton';
import InputNumber from 'primevue/inputnumber';
import Dialog from 'primevue/dialog';

import '@/assets/styles.scss';


const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(PrimeVue, { ripple: true });
app.use(ToastService);

app.directive('ripple', Ripple);
app.directive('styleclass', StyleClass);

app.component('Button', Button);
app.component('Chart', Chart);
app.component('Dropdown', Dropdown);
app.component('InputText', InputText);
app.component('Password', Password);
app.component('Toast', Toast);
app.component('Column', Column);
app.component('Toolbar', Toolbar);
app.component('DataTable', DataTable);
app.component('Textarea', Textarea);
app.component('RadioButton', RadioButton);
app.component('InputNumber', InputNumber);
app.component('Dialog', Dialog);


app.mount('#app');


