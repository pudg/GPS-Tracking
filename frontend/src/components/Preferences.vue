<script>
import { ref, onMounted } from 'vue';
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

        const handleDeviceClick = () => {
            console.log('Getting all devices...');
        };

        const handleTrackClick = () => {
            console.log('Tracking all devices...');
        };

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
            sendMsg, handleDeviceClick,
            handleTrackClick, savePreferences,
            handleSortClick };
    },
}
</script>

<template>
    <div class="flex flex-row flex-wrap items-center justify-center w-full p-2 mb-1 rounded shadow">
        <button @click="handleDeviceClick" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Devices</button>
        <button @click="handleTrackClick" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Track</button>
        <button @click="handleSortClick" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Sort</button>
        <button @click="savePreferences" class="rounded border border-gray-400 text-white bg-indigo-700 p-3">Save</button>
    </div>
    <!-- <div class="flex flex-row justify-center items-center bg-red-500 w-full p-2 mb-1 h-1/6">
        <div class="border bg-green-500 p-2">Preferences1</div>
        <div class="border bg-green-500 p-2">Preferences2</div>
    </div> -->
</template>

<style>
</style>