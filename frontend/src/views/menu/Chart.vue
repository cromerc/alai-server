<script setup>
import { ref, watch, onMounted } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import auth from '../../utils/Auth';

const { layoutConfig } = useLayout();
let documentStyle = getComputedStyle(document.documentElement);
let textColor = documentStyle.getPropertyValue('--text-color');
let textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
let surfaceBorder = documentStyle.getPropertyValue('--surface-border');

const lineData = ref(null);
const lineOptions = ref(null);

const setColorOptions = () => {
    documentStyle = getComputedStyle(document.documentElement);
    textColor = documentStyle.getPropertyValue('--text-color');
    textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    surfaceBorder = documentStyle.getPropertyValue('--surface-border');
};
const setChart = () => {
    lineData.value = {
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
        datasets: [
            {
                label: 'First Dataset',
                data: [65, 59, 80, 81, 56, 55, 40],
                fill: false,
                backgroundColor: documentStyle.getPropertyValue('--primary-500'),
                borderColor: documentStyle.getPropertyValue('--primary-500'),
                tension: 0.4
            },
            {
                label: 'Second Dataset',
                data: [28, 48, 40, 19, 86, 27, 90],
                fill: false,
                backgroundColor: documentStyle.getPropertyValue('--primary-200'),
                borderColor: documentStyle.getPropertyValue('--primary-200'),
                tension: 0.4
            }
        ]
    };

    lineOptions.value = {
        plugins: {
            legend: {
                labels: {
                    fontColor: textColor
                }
            }
        },
        scales: {
            x: {
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder,
                    drawBorder: false
                }
            },
            y: {
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder,
                    drawBorder: false
                }
            }
        }
    };
}
watch(
    layoutConfig.theme,
    () => {
        setColorOptions();
        setChart();
    },
    { immediate: true }
);

onMounted(() => {
    checkAuth();
})

const checkAuth = () => {
    auth.checkToken(true);
};
</script>

<template>
    <div class="card">
        <div class="grid p-fluid">
            <div class="col-12 xl:col-2">
                <span class="p-float-label">
                    <Dropdown id="dropdown" :options="cities" v-model="value8" optionLabel="name"></Dropdown>
                    <label for="dropdown">Dropdown</label>
                </span>
            </div>
        </div>
    </div>
    <div class="card">
        <div class="grid p-fluid">
            <div class="col-12 xl:col-6">
                <div class="card">
                    <h5>Linear Chart</h5>
                    <Chart type="line" :data="lineData" :options="lineOptions"></Chart>
                </div>
            </div>
        </div>
    </div>
</template>