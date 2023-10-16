<script setup>
const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_638790418.jpeg'
// const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_300774689.jpeg'
import Auth from '../components/Auth.vue';
import VLazyImage from "v-lazy-image";

</script>

<template>
    <div class="login-overlay-text-position flex justify-center">
        <v-lazy-image
            :src="URL"
            style="width: 100%; height: 95vh; contain: content;"
        />
        <div class="login-centered-text" style="background-color: aliceblue;">
            <Auth />
        </div>
    </div>

</template>

<style>
.login-style {
    width: 25vh;
    height: 20vh;
    border-radius: 50px;
    background-color: rgb(238, 236, 236);
}
.login-overlay-text-position {
    position: relative;
    text-align: center;
    color:black;
    font-size: xx-large;
}
.login-centered-text {
    max-width: 33vh;
    max-height: 25vh;
    position: absolute;
    top: 45%;
    left: 50%;
    border-radius: 4px;
    transform: translate(-50%, -50%);
}
</style>

