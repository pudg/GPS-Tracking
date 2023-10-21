<script>
import { ref, onMounted } from 'vue';
export default {
    setup() {
        let ws = ref(null);
        const sendMsg = () => {
            if (!ws) {
                return;
            }
            ws.send(JSON.stringify({"nelson": "hey"}))
            return;
        }
        onMounted((event) => {
            ws = new WebSocket("ws://localhost:8000/devices");
            if (event) {
                console.log(event.target.name);
            }
            ws.onopen = (event) => {
                console.log("ws open");
            }
            ws.onclose = (event) => {
                console.log("ws closed");
                ws = null;
            }
            ws.onmessage = (event) => {
                console.log("ws response: ", event.data);
            }
            ws.onerror = (event) => {
                console.log("err: ", event.data);
            }
        })
        return { sendMsg };
    },
}
</script>

<template>
    <div class="flex justify-center">
        <div class="inline-flex rounded-md shadow-sm" role="group">
        <button @click="allDevices()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-l-lg hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
            Devices
        </button>
    
        <button @click="sendMsg()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-r-md hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
            Track
        </button>
        </div>
        <div class="inline-flex rounded-md shadow-sm" role="group">
            <button type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-l-lg hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
                Reset
            </button>
            <button type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-r-md hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
                Save
            </button>
        </div>
    </div>

</template>

<style>
.prerences-button-styling {
    max-width: 25%;
}
.preferences-styling {
    max-width: 100%;
    height: 10vh;
}
</style>