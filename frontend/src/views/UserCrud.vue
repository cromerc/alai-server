<script setup>
import { FilterMatchMode } from 'primevue/api';
import { ref, onMounted, onBeforeMount } from 'vue';
import { useToast } from 'primevue/usetoast';
import axios from 'axios';
import auth from '../utils/Auth';

const toast = useToast();
const users = ref(null);
const userDialog = ref(false);
const deleteUserDialog = ref(false);
const deleteUsersDialog = ref(false);
const user = ref({});
const selectedUsers = ref(null);
const dt = ref(null);
const filters = ref({});
const submitted = ref(false);
const userDialogChange = ref(false);
const checkAuth = () => {
    auth.checkToken(true);
};
const url = new URL(window.location.href);
const api = (url.port == "5173") ? "http://localhost:3001" : "/api/v1";

async function showTable() {
    try {
        const response = await axios.get(api + `/user`, auth.getTokenHeader());
        users.value = response.data;
    }
    catch (error) {
        console.error(error);
    }
}

async function onCreateClick() {
    submitted.value = true;
    if (user.value.name === "") {
        throw new Error("name empty");
    }
    else if (user.value.username === "") {
        throw new Error("username empty");
    }
    else if (user.value.email === "") {
        throw new Error("email empty");
    }
    var newUser =
    {
        name: user.value.name,
        username: user.value.username,
        email: user.value.email
    };
    try {
        const response = await axios.post(api + `/user`, newUser, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'User Created', life: 3000 });
            userDialog.value = false;
            showTable();
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
    showTable();
});

const openNew = () => {
    user.value = {};
    submitted.value = false;
    userDialog.value = true;
};

const hideDialog = () => {
    userDialog.value = false;
    submitted.value = false;
};

const hideDialogChange = () => {
    userDialogChange.value = false;
    submitted.value = false;
};

const editUser = (editUser) => {

    user.value = { ...editUser };
    console.log(user);
    userDialogChange.value = true;
};

async function onChangeUser() {
    var changeUser =
    {
        name: user.value.name,
        username: user.value.username,
        email: user.value.email
    };
    try {
        const response = await axios.patch(api + `/user/` + user.value.ID, changeUser, auth.getTokenHeader());

        if (response.status === 204) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'User Modified', life: 3000 });
            userDialogChange.value = false;
            showTable();
        }
        else {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    }
}

const confirmDeleteUser = (editUser) => {
    user.value = editUser;
    deleteUserDialog.value = true;
};

async function deleteUser() {
    try {
        const response = await axios.delete(api + `/user/` + user.value.ID, auth.getTokenHeader());
        if (response.status !== 204) {
            console.error(response);
        }
        showTable();
    }
    catch (error) {
        console.error(error);
    }
    users.value = users.value.filter((val) => val.ID !== user.value.ID);
    deleteUserDialog.value = false;
    user.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'User Deleted', life: 3000 });
};


const confirmDeleteSelected = () => {
    deleteUsersDialog.value = true;
};
const deleteSelectedUsers = () => {
    selectedUsers.value.forEach(element => {
        try {
            const response = axios.delete(api + `/user/` + element.ID, auth.getTokenHeader());
            if (response.status !== 204) {
                console.error(response);
            }
        }
        catch (error) {
            console.error(error);
        }

    });
    console.log(selectedUsers.value);

    users.value = users.value.filter((val) => !selectedUsers.value.includes(val));
    deleteUsersDialog.value = false;
    selectedUsers.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Users Deleted', life: 3000 });
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
                            <Button label="Delete" icon="pi pi-trash" class="p-button-danger" @click="confirmDeleteSelected"
                                :disabled="!selectedUsers || !selectedUsers.length" />
                        </div>
                    </template>


                </Toolbar>

                <DataTable ref="dt" :value="users" v-model:selection="selectedUsers" dataKey="ID" :paginator="true"
                    :rows="10" :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[5, 10, 25]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} users" responsiveLayout="scroll">
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
                    <Column field="username" header="Username" :sortable="true" headerStyle="width:14%; Fin-width:10rem;">
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
                                @click="editUser(slotProps.data)" />
                            <Button icon="pi pi-trash" class="p-button-rounded p-button-warning mt-2"
                                @click="confirmDeleteUser(slotProps.data)" />
                        </template>
                    </Column>
                </DataTable>

                <Dialog v-model:visible="userDialog" :style="{ width: '450px' }" header="New User" :modal="true"
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
                <Dialog v-model:visible="userDialogChange" :style="{ width: '450px' }" header="New User" :modal="true"
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
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialogChange" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onChangeUser" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteUserDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="user">Are you sure you want to delete <b>{{ user.name }}</b>?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteUserDialog = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteUser" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteUsersDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="user">Are you sure you want to delete the selected users?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteUsersDialog = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteSelectedUsers" />
                    </template>
                </Dialog>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
@import '@/assets/badges.scss';
</style>
