<script>
import { ref, onMounted } from 'vue';
import store from '../store'
export default {
    setup() {
        let ws = ref(null);
        const sendMsg = () => {
            if (!ws) {
                return;
            }
            ws.send(JSON.stringify({ "nelson": "hey" }));
            return;
        };

        // const handleDeviceClick = () => {
        //     console.log('Getting all devices...');
        //     store.dispatch('searchDevices');
        //     console.log("Final Devices: ", store.state.devices);
        // };

        // const handleTrackClick = () => {
        //     console.log('Tracking all devices...');
        // };

        const handleSortClick = () => {
            console.log("Changing sort...");
        };

        const savePreferences = () => {
            console.log("Saving preferences...");
        };

        onMounted((event) => {
            ws = new WebSocket("ws://localhost:8000/devices");
            if (event) {
                console.log(event.target.name);
            }
            ws.onopen = (event) => {
                console.log("ws open");
            };
            ws.onclose = (event) => {
                console.log("ws closed");
                ws = null;
            };
            ws.onmessage = (event) => {
                console.log("ws response: ", event.data);
            };
            ws.onerror = (event) => {
                console.log("err: ", event.data);
            };
        });
        return {
            sendMsg, savePreferences, handleSortClick };
    },
    props: {
        trackDevices: Function,
        allDevices: Function
    }
}
</script>

<template>
    <div class="flex flex-row flex-wrap items-center justify-center w-full p-2 mb-1 rounded shadow">
        <button @click="allDevices" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Devices</button>
        <button @click="trackDevices" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Track</button>
        <button @click="handleSortClick" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Sort</button>
        <button @click="savePreferences" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Save</button>
    </div>
</template>

<style>
</style>