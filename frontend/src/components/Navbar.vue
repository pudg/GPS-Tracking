<script>
import Button from './Button.vue';
import { ref } from 'vue';
export default {
    components: {
        Button
    },
    setup() {
        const open = ref(false);
        const links = [
            {name: 'Home', link: 'home'},
            {name: 'Tracking', link: 'tracking'},
            {name: 'Login', link: 'login'}
        ];
        
        const toggleOpen = () => {
            open.value = !open.value;
        };

        return { links, open, toggleOpen };
    },

};
</script>

<template>
    <div class="text-gray-300 py-3.5 px-6
                shadow md:flex justify-between items-center"
                >
        <div class="flex items-center cursor-pointer">
            <span class="text-indigo-700 text-xl mr-1 mt-0">
                <i class="bi bi-messenger"></i>
            </span>
            <h1 style="font-size: large;" class="inline-blockh-full text-indigo-700 rounded font-extrabold">One Step</h1>
        </div>
        <span @click="toggleOpen()" class="absolute md:hidden right-6 top-1.5 cursor-pointer text-4xl">
            <i :class="[open ? 'bi bi-x': 'bi bi-filter-left']"></i>
        </span>
        <ul class="md:flex md:items-center md:px-0 px-4 md:pb-0 pb-6 md:static absolute
        bg-gray-100 md:w-auto w-full top-14 duration-700 ease-in"
        :class="[open ? 'left-0': 'left-[-100%]']">
            <li class="md:mx-4" v-for="link in links">
                <router-link :to="{name: link.link}" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    {{ link.name }}
                </router-link>
            </li>
        </ul>

    </div>
</template>

<style>
</style>