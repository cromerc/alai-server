<script setup>
import { FilterMatchMode } from 'primevue/api';
import { ref, onMounted, onBeforeMount } from 'vue';
import { useToast } from 'primevue/usetoast';
import axios from 'axios';
import auth from '../utils/Auth';

const toast = useToast();
const products = ref(null);
const productDialog = ref(false);
const deleteProductDialog = ref(false);
const deleteProductsDialog = ref(false);
const user = ref({});
const selectedProducts = ref(null);
const dt = ref(null);
const filters = ref({});
const submitted = ref(false);
const productDialogChange = ref(false);
const checkAuth = () => {
    auth.checkToken(true);
};

async function onShowClick() {
    try {
        const response = await axios.get(`http://localhost:3001/user`, auth.getTokenHeader());
        console.log(response.data);
        products.value = response.data;
    }
    catch (error) {
        console.error(error);
    }
}

async function onCreateClick() {
    var newUser =
    {
        name: user.value.name,
        username: user.value.username,
        email: user.value.email
    };
    try {
        const response = await axios.post(`http://localhost:3001/user`, newUser, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'User Created', life: 3000 });
            productDialog.value = false;
        }
        else {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    }

}

onBeforeMount(() => {
    initFilters();
});
onMounted(() => {
    checkAuth();
});

const openNew = () => {
    user.value = {};
    submitted.value = false;
    productDialog.value = true;
};

const hideDialog = () => {
    productDialog.value = false;
    submitted.value = false;
};

const hideDialogChange = () => {
    productDialogChange.value = false;
    submitted.value = false;
};

const editProduct = (editProduct) => {

    user.value = { ...editProduct };
    console.log(user);
    productDialogChange.value = true;
};

async function onChangeUser() {
    var changeUser =
    {
        name: user.value.name,
        username: user.value.username,
        email: user.value.email
    };
    try {
        const response = await axios.patch(`http://localhost:3001/user/` + user.value.ID, changeUser, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'User Modified', life: 3000 });
            productDialogChange.value = false;
        }
        else {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    }
}

const confirmDeleteProduct = (editProduct) => {
    user.value = editProduct;
    deleteProductDialog.value = true;
};

async function deleteProduct() {
    try {
        const response = await axios.delete(`http://localhost:3001/user/` + user.value.ID, auth.getTokenHeader());
        if (response.status !== 204) {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    }
    products.value = products.value.filter((val) => val.ID !== user.value.ID);
    deleteProductDialog.value = false;
    user.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'User Deleted', life: 3000 });
};


const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
const deleteSelectedProducts = () => {
    selectedProducts.value.forEach(element => {
        try {
            const response = axios.delete(`http://localhost:3001/user/` + element.ID, auth.getTokenHeader());
            if (response.status !== 204) {
                console.error(response);
            }
        }
        catch (error) {
            console.error(error);
        }

    });
    console.log(selectedProducts.value);

    products.value = products.value.filter((val) => !selectedProducts.value.includes(val));
    deleteProductsDialog.value = false;
    selectedProducts.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};

const initFilters = () => {
    filters.value = {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS }
    };
};
</script>

<template>
    <div class="grid">
        <div class="col-12">
            <div class="card">
                <Toast />
                <Toolbar class="mb-4">
                    <template v-slot:start>
                        <div class="my-2">
                            <Button label="New" icon="pi pi-plus" class="p-button-success mr-2" @click="openNew" />
                            <Button label="Show" icon="pi pi-eye" class="p-button-success mr-2" @click="onShowClick" />
                            <Button label="Delete" icon="pi pi-trash" class="p-button-danger"
                                @click="confirmDeleteSelected"
                                :disabled="!selectedProducts || !selectedProducts.length" />
                        </div>
                    </template>


                </Toolbar>

                <DataTable ref="dt" :value="products" v-model:selection="selectedProducts" dataKey="ID"
                    :paginator="true" :rows="10" :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[5, 10, 25]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products"
                    responsiveLayout="scroll">
                    <template #header>
                        <div class="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
                            <h5 class="m-0">Manage Users</h5>
                            <span class="block mt-2 md:mt-0 p-input-icon-left">
                                <i class="pi pi-search" />
                                <InputText v-model="filters['global'].value" placeholder="Search..." />
                            </span>
                        </div>
                    </template>

                    <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
                    <Column field="name" header="Name" :sortable="true" headerStyle="width:14%; Fin-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Name</span>
                            {{ slotProps.data.name }}
                        </template>
                    </Column>
                    <Column field="username" header="Username" :sortable="true"
                        headerStyle="width:14%; Fin-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Username</span>
                            {{ slotProps.data.username }}
                        </template>
                    </Column>
                    <Column field="email" header="Email" :sortable="true" headerStyle="width:14%; Fin-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">Email</span>
                            {{ slotProps.data.email }}
                        </template>
                    </Column>
                    <Column headerStyle="min-width:10rem;">
                        <template #body="slotProps">
                            <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2"
                                @click="editProduct(slotProps.data)" />
                            <Button icon="pi pi-trash" class="p-button-rounded p-button-warning mt-2"
                                @click="confirmDeleteProduct(slotProps.data)" />
                        </template>
                    </Column>
                </DataTable>

                <Dialog v-model:visible="productDialog" :style="{ width: '450px' }" header="New User" :modal="true"
                    class="p-fluid">
                    <div class="field">
                        <label for="name">Name</label>
                        <InputText id="name" v-model.trim="user.name" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.name }" />
                        <small class="p-invalid" v-if="submitted && !user.name">Name is required.</small>
                    </div>
                    <div class="field">
                        <label for="username">Username</label>
                        <InputText id="username" v-model.trim="user.username" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.username }" />
                        <small class="p-invalid" v-if="submitted && !user.username">Username is required.</small>
                    </div>
                    <div class="field">
                        <label for="email">Email</label>
                        <InputText id="email" v-model.trim="user.email" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.email }" />
                        <small class="p-invalid" v-if="submitted && !user.email">Email is required.</small>
                    </div>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onCreateClick" />
                    </template>
                </Dialog>
                <Dialog v-model:visible="productDialogChange" :style="{ width: '450px' }" header="New User"
                    :modal="true" class="p-fluid">
                    <div class="field">
                        <label for="name">Name</label>
                        <InputText id="name" v-model.trim="user.name" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.name }" />
                        <small class="p-invalid" v-if="submitted && !user.name">Name is required.</small>
                    </div>
                    <div class="field">
                        <label for="username">Username</label>
                        <InputText id="username" v-model.trim="user.username" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.username }" />
                        <small class="p-invalid" v-if="submitted && !user.username">Username is required.</small>
                    </div>
                    <div class="field">
                        <label for="email">Email</label>
                        <InputText id="email" v-model.trim="user.email" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !user.email }" />
                        <small class="p-invalid" v-if="submitted && !user.email">Email is required.</small>
                    </div>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialogChange" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onChangeUser" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteProductDialog" :style="{ width: '450px' }" header="Confirm"
                    :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="user">Are you sure you want to delete <b>{{ user.name }}</b>?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text"
                            @click="deleteProductDialog = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteProduct" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteProductsDialog" :style="{ width: '450px' }" header="Confirm"
                    :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="user">Are you sure you want to delete the selected products?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text"
                            @click="deleteProductsDialog = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteSelectedProducts" />
                    </template>
                </Dialog>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
@import '@/assets/badges.scss';
</style>
