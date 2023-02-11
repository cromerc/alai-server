<script setup>
import { FilterMatchMode } from 'primevue/api';
import { ref, onMounted, onBeforeMount } from 'vue';
import { useToast } from 'primevue/usetoast';
import axios from 'axios';
import auth from '../utils/Auth';

const toast = useToast();
const products = ref(null);
const productDialog = ref(false);
const productDialogChange = ref(false);
const deleteProductDialog = ref(false);
const deleteProductsDialog = ref(false);
const os = ref({});
const selectedProducts = ref(null);
const dt = ref(null);
const filters = ref({});
const submitted = ref(false);
const checkAuth = () => {
    auth.checkToken(true);
};

async function onShowClick() {
    try {
        const response = await axios.get(`http://localhost:3001/os`);
        console.log(response.data);
        products.value = response.data;
    }
    catch (error) {
        console.error(error);
    }
}

async function onCreateClick() {
    submitted.value = true;
    if (os.value.name === "") {
        throw new Error("name empty");
    }
    var newOS =
    {
        name: os.value.name
    };
    try {
        const response = await axios.post(`http://localhost:3001/os`, newOS, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'OS Created', life: 3000 });
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
    os.value = {};
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
    os.value = { ...editProduct };
    console.log(os.value.ID);
    productDialogChange.value = true;
};
async function onChangeOS() {
    var changeOS =
    {
        name: os.value.name
    };
    console.log(os.value.ID);
    try {
        const response = await axios.patch(`http://localhost:3001/os/` + os.value.ID, changeOS, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'OS Modified', life: 3000 });
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
    os.value = editProduct;
    deleteProductDialog.value = true;
};

async function deleteProduct() {
    try {
        const response = await axios.delete(`http://localhost:3001/os/` + os.value.ID, auth.getTokenHeader());
        if (response.status !== 204) {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    }
    products.value = products.value.filter((val) => val.ID !== os.value.ID);
    deleteProductDialog.value = false;
    os.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'OS Deleted', life: 3000 });
};


const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
const deleteSelectedProducts = () => {
    selectedProducts.value.forEach(element => {
        try {
            const response = axios.delete(`http://localhost:3001/os/` + element.ID, auth.getTokenHeader());
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
                            <h5 class="m-0">Manage Products</h5>
                            <span class="block mt-2 md:mt-0 p-input-icon-left">
                                <i class="pi pi-search" />
                                <InputText v-model="filters['global'].value" placeholder="Search..." />
                            </span>
                        </div>
                    </template>

                    <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
                    <Column field="name" header="OS Name" :sortable="true" headerStyle="width:14%; Fin-width:10rem;">
                        <template #body="slotProps">
                            <span class="p-column-title">OS Name</span>
                            {{ slotProps.data.name }}
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

                <Dialog v-model:visible="productDialog" :style="{ width: '450px' }" header="New OS" :modal="true"
                    class="p-fluid">
                    <div class="field">
                        <label for="name">Name</label>
                        <InputText id="name" v-model.trim="os.name" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !os.name }" />
                        <small class="p-invalid" v-if="submitted && !os.name">Name is required.</small>
                    </div>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onCreateClick" />
                    </template>
                </Dialog>
                <Dialog v-model:visible="productDialogChange" :style="{ width: '450px' }" header="New OS" :modal="true"
                    class="p-fluid">
                    <div class="field">
                        <label for="name">Name</label>
                        <InputText id="name" v-model.trim="os.name" required="true" autofocus
                            :class="{ 'p-invalid': submitted && !os.name }" />
                        <small class="p-invalid" v-if="submitted && !os.name">Name is required.</small>
                    </div>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialogChange" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onChangeOS" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteProductDialog" :style="{ width: '450px' }" header="Confirm"
                    :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="os">Are you sure you want to delete <b>{{ os.name }}</b>?</span>
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
                        <span v-if="os">Are you sure you want to delete the selected products?</span>
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
