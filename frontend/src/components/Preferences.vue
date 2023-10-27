<template>
    <div class="flex flex-row flex-wrap items-center justify-center w-full p-2 mb-1 rounded shadow">
        <button @click="allDevices" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Devices</button>
        <button @click="trackDevices" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Track</button>
        <button @click="handleSortClick" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Sort</button>
        <button @click="savePreferences" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Save</button>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import store from '../store'
import { useStore, mapState, mapActions } from 'vuex'
export default {
    setup() {
        const store = useStore();
        const handleSortClick = () => {
            store.dispatch('deviceSort');
        };
        const savePreferences = () => {
            const preferences = {
                user: store.state.user,
                devices: store.state.devices,
                sort: store.state.sortAsc
            };
            store.dispatch('saveUserPreferences', preferences);
        };
        return {
            savePreferences, handleSortClick };
    },
    props: {
        trackDevices: Function,
        allDevices: Function,
    }
}
</script>