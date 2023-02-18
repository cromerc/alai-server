<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
const username = ref('');
const password = ref('');
const router = useRouter();
const toast = useToast();

async function onLoginClick() {
    var login =
    {
        username: username.value,
        password: password.value
    };
    try {
        const response = await axios.post(`http://localhost:3001/login`, login);
        if (response.status === 200) {
            localStorage.setItem('token', response.data.token);
            router.push('/home');
        }
    }
    catch (error) {
        if (error.response.status === 400) {
            toast.add({ severity: 'error', summary: 'Error', detail: 'Incorrect Password', life: 3000 });
        }
        else {
            console.log(error);
        }
    }
}

onMounted(() => {
    if (localStorage.getItem("token") != null) {
        router.push('/home');
    }
});

</script>

<template>
    <div class="surface-ground flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden">
        <div class="flex flex-column align-items-center justify-content-center">

            <div
                style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
                <div class="w-full surface-card py-8 px-5 sm:px-8" style="border-radius: 53px">
                    <div class="text-center mb-5">
                        <span class="text-600 font-medium">Log in to continue</span>
                    </div>
                    <Toast />
                    <div>
                        <label for="username1" class="block text-900 text-xl font-medium mb-2">Username</label>
                        <InputText id="username1" type="text" placeholder="Username" class="w-full md:w-30rem mb-5"
                            style="padding: 1rem" v-model="username" />

                        <label for="password1" class="block text-900 font-medium text-xl mb-2">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" class="w-full mb-3"
                            inputClass="w-full" inputStyle="padding:1rem" :feedback="false">
                        </Password>

                        <Button label="Log In" class="w-full p-3 text-xl" @click="onLoginClick()"></Button>

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.pi-eye {
    transform: scale(1.6);
    margin-right: 1rem;
}

.pi-eye-slash {
    transform: scale(1.6);
    margin-right: 1rem;
}
</style>
