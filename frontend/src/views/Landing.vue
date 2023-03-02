<script setup>
import { saveAs } from 'file-saver';
const smoothScroll = (id) => {
    document.querySelector(id).scrollIntoView({
        behavior: 'smooth'
    });
};

const onClickWindows = () => {
    saveAs("/game/alai-1.0.0.zip", "alai-1.0.0.zip");
}
const onClickLinux = () => {
    saveAs("/game/alai-1.0.0.tar.gz", "alai-1.0.0.tar.gz");
}

</script>

<template>
    <div class="surface-0 flex justify-content-center">
        <div id="home" class="landing-wrapper overflow-hidden">
            <div
                class="py-4 px-4 mx-0 md:mx-6 lg:mx-8 lg:px-8 flex align-items-center justify-content-between relative lg:static mb-3">

                <a class="cursor-pointer block lg:hidden text-700 p-ripple" v-ripple
                    v-styleclass="{ selector: '@next', enterClass: 'hidden', leaveToClass: 'hidden', hideOnOutsideClick: true }">
                    <i class="pi pi-bars text-4xl"></i>
                </a>
                <div class="align-items-center surface-0 flex-grow-1 justify-content-between hidden lg:flex absolute lg:static w-full left-0 px-6 lg:px-0 z-2"
                    style="top: 120px">
                    <ul
                        class="list-none p-0 m-0 flex lg:align-items-center select-none flex-column lg:flex-row cursor-pointer">
                        <li>
                            <router-link to="/"
                                class="flex m-0 md:ml-5 px-0 py-3 text-900 font-medium line-height-3 p-ripple" v-ripple>
                                <span>Home</span>
                            </router-link>
                        </li>
                        <li>
                            <a @click="smoothScroll('#features')"
                                class="flex m-0 md:ml-5 px-0 py-3 text-900 font-medium line-height-3 p-ripple" v-ripple>
                                <span>Features</span>
                            </a>
                        </li>
                        <li>
                            <a @click="smoothScroll('#moreinfo')"
                                class="flex m-0 md:ml-5 px-0 py-3 text-900 font-medium line-height-3 p-ripple" v-ripple>
                                <span>More info</span>
                            </a>
                        </li>
                        <li>
                            <a @click="smoothScroll('#downloads')"
                                class="flex m-0 md:ml-5 px-0 py-3 text-900 font-medium line-height-3 p-ripple" v-ripple>
                                <span>Downloads</span>
                            </a>
                        </li>

                    </ul>
                    <div
                        class="flex justify-content-between lg:block border-top-1 lg:border-top-none surface-border py-3 lg:py-0 mt-3 lg:mt-0">
                        <router-link to="/login">
                            <Button label="Login"
                                class="p-button-text p-button-rounded border-none font-light text-white line-height-2 bg-blue-500">
                            </Button>
                        </router-link>
                    </div>

                </div>
            </div>

            <div id="hero" class="flex flex-column pt-4 px-4 lg:px-8 overflow-hidden"
                style="background: linear-gradient(0deg, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.2)), radial-gradient(77.36% 256.97% at 77.36% 57.52%, rgb(238, 239, 175) 0%, rgb(195, 227, 250) 100%); clip-path: ellipse(150% 87% at 93% 13%)">
                <div class="mx-4 md:mx-8 mt-0 md:mt-4">
                    <h1 class="text-6xl font-bold text-gray-900 line-height-2"><span class="font-light block">A.I.
                            Development</span>Alai &
                        Obelisk</h1>
                    <p class="font-normal text-2xl line-height-3 md:mt-3 text-gray-700 mb-8">Artificial Intelligence
                        implementation in a 2D-based platform videogame using a self-made declarative language</p>
                </div>
            </div>


            <div id="features" class="py-4 px-4 lg:px-8 mx-0 my-6 lg:mx-8">
                <div class="text-center">
                    <h2 class="text-900 font-normal mb-2">Features</h2>

                </div>

                <div class="grid mt-8 pb-2 md:pb-8">
                    <div class="flex justify-content-center col-12 lg:col-6  p-0 flex-order-1 lg:flex-order-0"
                        style="border-radius: 8px">
                        <img src="@/assets/images/alai.png" class="w-11" alt="alai screenshot" />
                    </div>

                    <div class="col-12 lg:col-6 my-auto flex flex-column lg:align-items-end text-center lg:text-right">
                        <div class="flex align-items-center justify-content-center bg-purple-200 align-self-center lg:align-self-end"
                            style="width: 4.2rem; height: 4.2rem; border-radius: 10px">
                            <i class="pi pi-fw pi-palette text-5xl text-purple-700"></i>
                        </div>
                        <h2 class="line-height-1 text-900 text-4xl font-normal">Alai</h2>
                        <span class="text-700 text-2xl line-height-3 ml-0 md:ml-2" style="max-width: 650px">Alai is a
                            2D-based platform videogame made with a friendly and easy-to-use game engine called
                            "Godot". Alai contains within the game a data harvest trait
                            for each playthrough, with the purpose to analyze the information for future use</span>
                    </div>
                </div>

                <div class="grid my-8 pt-2 md:pt-8">
                    <div class="col-12 lg:col-6 my-auto flex flex-column text-center lg:text-left lg:align-items-start">
                        <div class="flex align-items-center justify-content-center bg-blue-200 align-self-center lg:align-self-start"
                            style="width: 4.2rem; height: 4.2rem; border-radius: 10px">
                            <i class="pi pi-fw pi-globe text-5xl text-blue-700"></i>
                        </div>
                        <h2 class="line-height-1 text-900 text-4xl font-normal">Obelisk</h2>
                        <span class="text-700 text-2xl line-height-3 mr-0 md:mr-2" style="max-width: 650px">A compiled
                            and declarative self-made language with the purpose of the creation of an Artifitial
                            Intelligence, inspired by the logical language "Prolog"</span>
                    </div>

                    <div class="flex align-items-center justify-content-end flex-order-1 sm:flex-order-2 col-12 lg:col-6  p-0"
                        style="border-radius: 8px">
                        <img src="@/assets/images/obelisk.jpg" class="flex w-11" alt="obelisk" width="100" height="100" />
                    </div>
                </div>
            </div>

            <div id="downloads" class="py-4 px-4 lg:px-8 mx-0 my-6 lg:mx-8">
                <div class="text-center">
                    <h2 class="text-900 font-normal mb-6">Downloads</h2>
                </div>

                <div class="flex justify-content-center p-4 flex-order-1 lg:flex-order-0" style="border-radius: 8px">

                    <Button @click="onClickWindows()" class="flex justify-content-center p-button-lg px-0 w-2 mr-4">
                        <i class="pi pi-microsoft px-1"></i>
                        <span class="px-2">
                            Windows
                        </span>
                    </Button>

                    <Button @click="onClickLinux()" class="flex justify-content-center p-button-lg px-0 w-2 ml-4">
                        <i class="pi pi-globe px-1"></i>
                        <span class="px-2">
                            Linux
                        </span>
                    </Button>

                </div>

            </div>



            <div class="py-4 px-4 mx-0 mt-8 lg:mx-8">
                <div class="grid justify-content-between">
                    <div class="col-12 md:col-2" style="margin-top: -1.5rem">

                    </div>

                    <div class="col-12 md:col-10 lg:col-2">
                        <div class="grid text-center md:text-left">
                            <div class="col-12 md:col-3">
                                <h4 class="font-medium text-2xl line-height-3 mb-3 text-900">Project</h4>
                                <a class="line-height-3 text-xl block cursor-pointer mb-2 text-700">About Us</a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<!-- <style scoped>
#hero {
    background: linear-gradient(0deg, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.2)), radial-gradient(77.36% 256.97% at 77.36% 57.52%, #eeefaf 0%, #c3e3fa 100%);
    height: 700px;
    overflow: hidden;
}

@media screen and (min-width: 768px) {
    #hero {
        -webkit-clip-path: ellipse(150% 87% at 93% 13%);
        clip-path: ellipse(150% 87% at 93% 13%);
        height: 530px;
    }
}

@media screen and (min-width: 1300px) {
    #hero > img {
        position: absolute;
    }

    #hero > div > p {
        max-width: 450px;
    }
}

@media screen and (max-width: 1300px) {
    #hero {
        height: 600px;
    }

    #hero > img {
        position: static;
        transform: scale(1);
        margin-left: auto;
    }

    #hero > div {
        width: 100%;
    }

    #hero > div > p {
        width: 100%;
        max-width: 100%;
    }
}
</style> -->
