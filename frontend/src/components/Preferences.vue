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
        return { sendMsg };
    },
}
</script>

<template>
    <div class="bg-red-500 w-full p-2 mb-4 h-1/6">
        Preferences
    </div>
</template>

<style>
</style>