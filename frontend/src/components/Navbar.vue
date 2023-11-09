<template>
    <div class="
            text-gray-300 py-3.5 px-6
            shadow-md md:flex justify-between items-center">
        <div class="flex items-center cursor-pointer">
            <span class="text-indigo-700 text-xl mr-1">
                <i class="bi bi-geo-alt-fill"></i>
            </span>
            <h1 style="font-size: large;" class="inline-block h-full text-indigo-700 rounded font-extrabold">One Step</h1>
        </div>
        <span @click="toggleOpen(user)" class="absolute md:hidden right-6 top-1.5 cursor-pointer text-4xl text-indigo-700">
            <i :class="[open ? 'bi bi-x': 'bi bi-filter-left']"></i>
        </span>
        <ul class="md:flex md:items-center md:px-0 px-4 md:pb-0 pb-6 md:static absolute
        bg-gray-100 md:w-auto w-full top-14 duration-700 ease-in"
        :class="[open ? 'right-0': 'right-[-100%]']">
            <li class="md:mx-4" v-for="link in navLinks">
                <router-link :to="{name: link.link}" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    {{ link.name }}
                </router-link>
            </li>
            <li @click="logOut()" v-if="user" class="md:mx-4">
                <router-link to="/" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    Logout
                </router-link>
            </li>
            <li v-else class="md:mx-4">
                <router-link to="/login" style="font-size: large;" class="inline-block p-4 h-full text-indigo-700 hover:bg-indigo-800 rounded hover:text-white font-extrabold">
                    Login
                </router-link>
            </li>
        </ul>

    </div>
</template>

<script>
import { ref } from 'vue';
import { useStore, mapState } from 'vuex';
export default {
    components: {},
    setup() {
        const store = useStore();
        const open = ref(false);
        const toggleOpen = () => {
            open.value = !open.value;
        };

        const logOut = () => {
            store.dispatch('userLogout', store.state.user);
        }

        return { open, toggleOpen, logOut };
    },
    computed: {
		...mapState(['user', 'navLinks'])
	},
};
</script>