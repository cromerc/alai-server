<script setup>
import auth from '../utils/Auth';
import { onMounted } from 'vue';

onMounted(() => {
    checkAuth();
})

const checkAuth = () => {
    auth.checkToken(true);
};
</script>

<template>
    <main>

    </main>
</template>
