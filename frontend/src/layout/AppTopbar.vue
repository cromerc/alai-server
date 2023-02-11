<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import axios from 'axios';

const { onMenuToggle } = useLayout();

const outsideClickListener = ref(null);
const topbarMenuActive = ref(false);
const changePasswordDialog = ref(false);
const router = useRouter();
const user = ref({});
const submitted = ref(false);
const toast = useToast();
const menu = ref();
const items = ref([
    {
        label: 'Log out',
        icon: 'pi pi-sign-out',
        command: () => {
            toast.add({ severity: 'warn', summary: 'Delete', detail: 'Data Deleted', life: 3000 });
            localStorage.removeItem('token');
            router.push('/');
        }
    },
    {
        label: 'Change password',
        icon: 'pi pi-key',
        command: () => {
            changePasswordDialog.value = true;
            submitted.value = false;
            user.value = {};
        }
    }
]);

const toggle = (event) => {
    menu.value.toggle(event);
};

onMounted(() => {
    bindOutsideClickListener();
});

onBeforeUnmount(() => {
    unbindOutsideClickListener();
});


async function onClickPasswordChange() {
    submitted.value = true;
    console.log(user.value);
    var pass =

    {
        password: user.value.current_password,
        new_password: user.value.new_password
    };
    /* try {
        const response = await axios.patch(`http://localhost:3001/user/` + user.value.ID, auth.getTokenHeader());
        if (response.status !== 204) {
            console.error(response);
        }
    }
    catch (error) {
        console.error(error);
    } */
}



const onTopBarMenuButton = () => {
    topbarMenuActive.value = !topbarMenuActive.value;
};
const topbarMenuClasses = computed(() => {
    return {
        'layout-topbar-menu-mobile-active': topbarMenuActive.value
    };
});

const bindOutsideClickListener = () => {
    if (!outsideClickListener.value) {
        outsideClickListener.value = (event) => {
            if (isOutsideClicked(event)) {
                topbarMenuActive.value = false;
            }
        };
        document.addEventListener('click', outsideClickListener.value);
    }
};
const unbindOutsideClickListener = () => {
    if (outsideClickListener.value) {
        document.removeEventListener('click', outsideClickListener);
        outsideClickListener.value = null;
    }
};
const isOutsideClicked = (event) => {
    if (!topbarMenuActive.value) return;

    const sidebarEl = document.querySelector('.layout-topbar-menu');
    const topbarEl = document.querySelector('.layout-topbar-menu-button');

    return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
};
</script>

<template>
    <div class="layout-topbar">
        <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle()">
            <i class="pi pi-bars"></i>
        </button>

        <button class="p-link layout-topbar-menu-button layout-topbar-button" @click="onTopBarMenuButton()">
            <i class="pi pi-ellipsis-v"></i>
        </button>

        <div class="layout-topbar-menu" :class="topbarMenuClasses">
            <button @click="toggle" class="p-link layout-topbar-button">
                <i class="pi pi-user"></i>
                <span>Profile</span>
            </button>
            <Menu id="overlay_menu" ref="menu" :model="items" :popup="true" />
        </div>

        <Dialog v-model:visible="changePasswordDialog" :style="{ width: '450px' }" header="Change Password"
            :modal="true" class="p-fluid">

            <div class="field">
                <label for="current_password">Current Password</label>
                <InputText id="current_password" v-model.trim="user.current_password" required="true" autofocus
                    :class="{ 'p-invalid': submitted && !user.current_password }" />
                <small class="p-invalid" v-if="submitted && !user.current_password">Actual password is required.</small>
            </div>
            <div class="field">
                <label for="new_password">New Password</label>
                <InputText id="new_password" v-model.trim="user.new_password" required="true" autofocus
                    :class="{ 'p-invalid': submitted && !user.new_password }" />
                <small class="p-invalid" v-if="submitted && !user.new_password">New password is required.</small>
            </div>
            <div class="field">
                <label for="confirm_new_password">Confirm new Password</label>
                <InputText id="confirm_new_password" v-model.trim="user.confirm_new_password" required="true" autofocus
                    :class="{ 'p-invalid': submitted && !user.confirm_new_password }" />
                <small class="p-invalid" v-if="submitted && !user.confirm_new_password">Confirm the new password is
                    required.</small>
            </div>
            <template #footer>
                <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
                <Button label="Save" icon="pi pi-check" class="p-button-text" @click="onClickPasswordChange" />
            </template>
        </Dialog>
    </div>
</template>

<style lang="scss" scoped>

</style>
