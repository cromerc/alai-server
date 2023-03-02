<script setup>
import { ref, watch, onMounted } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import auth from '../utils/Auth';
import axios from 'axios';

const { layoutConfig } = useLayout();
let documentStyle = getComputedStyle(document.documentElement);
let textColor = documentStyle.getPropertyValue('--text-color');
let textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
let surfaceBorder = documentStyle.getPropertyValue('--surface-border');
const time_elapsed = ref([]);
const coins = ref([]);
const games = ref([]);
const ms = ref(true);
const selectedGame = ref();

const lineData = ref(null);
const lineOptions = ref(null);

const setColorOptions = () => {
    documentStyle = getComputedStyle(document.documentElement);
    textColor = documentStyle.getPropertyValue('--text-color');
    textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    surfaceBorder = documentStyle.getPropertyValue('--surface-border');
};
const url = new URL(window.location.href);
const api = (url.port == "5173") ? "http://localhost:3001" : "/api/v1";


async function getGames() {
    try {
        const response = await axios.get(api + `/game?limit=5000`);
        games.value = response.data.map(game_id => {
            return game_id.ID;
        });
    }
    catch (error) {
        console.error(error);
    }
}

async function getFrames() {
    try {
        if (selectedGame.value === undefined) {
            return;
        }
        const response = await axios.get(api + `/frame?limit=5000&game_id=` + selectedGame.value);
        if (response.data[response.data.length - 1].elapsed_time >= 10000) {
            ms.value = false;
        }
        else {
            ms.value = true;
        }
        time_elapsed.value = response.data.map(frame => {
            return (ms.value == true) ? (Math.round(frame.elapsed_time * 10) / 10).toFixed(1) : (Math.floor(frame.elapsed_time / 1000 * 10) / 10).toFixed(1);
        });
        coins.value = response.data.map(frame => {
            return frame.coins;
        });
    }
    catch (error) {
        console.error(error);
    }

    lineData.value.labels = time_elapsed.value;
    lineData.value.datasets[0].label = 'Game ' + selectedGame.value;
    lineData.value.datasets[0].data = coins.value;

    lineOptions.value.scales.x.title.text = (ms.value) ? "Time(ms)" : "Time(s)";

}

async function setChart() {
    lineData.value = {
        labels: time_elapsed.value,
        datasets: [
            {
                label: 'Game ' + selectedGame.value,
                data: coins.value,
                fill: false,
                backgroundColor: documentStyle.getPropertyValue('--primary-500'),
                borderColor: documentStyle.getPropertyValue('--primary-500'),
                tension: 0.4
            }
        ]
    };

    lineOptions.value = {
        reponsive: true,
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
                    text: (ms.value) ? "Time(ms)" : "Time(s)"
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
    },
    { immediate: true }
);

onMounted(() => {
    setChart();
    checkAuth();
    getGames();
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
                    <Dropdown v-model="selectedGame" :options="games" optionLabel="" optionValue="" @change="getFrames()" />
                    <label for="dropdown">Game session</label>
                </span>
            </div>
        </div>
    </div>
    <div class="card">
        <div class="grid p-fluid">
            <div class="col-12 xl:col-6">
                <Chart type="line" :data="lineData" :options="lineOptions"></Chart>
            </div>
        </div>
    </div>
</template>

<style>
canvas {
    width: 500px !important;
    height: 310px !important;
}
</style>