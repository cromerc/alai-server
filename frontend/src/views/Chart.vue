<script setup>
import { ref, watch, onMounted } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import auth from '../utils/Auth';
import { useToast } from 'primevue/usetoast';
import axios from 'axios';

const { layoutConfig } = useLayout();
const toast = useToast();
let documentStyle = getComputedStyle(document.documentElement);
let textColor = documentStyle.getPropertyValue('--text-color');
let textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
let surfaceBorder = documentStyle.getPropertyValue('--surface-border');
const time_elapsed = ref([]);
const coins = ref([]);

const lineData = ref(null);
const lineOptions = ref(null);

const setColorOptions = () => {
    documentStyle = getComputedStyle(document.documentElement);
    textColor = documentStyle.getPropertyValue('--text-color');
    textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    surfaceBorder = documentStyle.getPropertyValue('--surface-border');
};
async function setChart() {
    try {
        const response = await axios.get(`http://localhost:3001/frame?limit=5000&game_id=4`);
        time_elapsed.value = response.data.map(frame => {
            return Math.floor(frame.elapsed_time / 1000);
        });
        coins.value = response.data.map(frame => {
            return frame.coins;
        });
    }
    catch (error) {
        console.error(error);
    }


    /////////////////////////////////
    lineData.value = {
        labels: time_elapsed.value,
        datasets: [
            {
                label: 'Game 4',
                data: coins.value,
                fill: false,
                backgroundColor: documentStyle.getPropertyValue('--primary-500'),
                borderColor: documentStyle.getPropertyValue('--primary-500'),
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
                title: {
                    display: true,
                    text: "Time(s)"
                },
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder,
                    drawBorder: false
                }
            },
            y: {
                title: {
                    display: true,
                    text: "Coins"
                },
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
        //setChart();
    },
    { immediate: true }
);

onMounted(() => {
    checkAuth();
    setChart();
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