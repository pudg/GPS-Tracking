<script setup>
const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_300774689.jpeg'
</script>

<template>
    <div class="overlay-text-position">
        <img
        style="width: 100%; height: 100vh; contain: content;"
        :src="URL"
        alt="GPS" />
        <div class="centered-text">
            <p class="font-bold">
                Navigate with Confidence, Powered by NavTrek
            </p>
        </div>
    </div>

</template>

<style>
/* image-style {
    width: 100%;
    height: 100vh;
    background-size: cover;
    border-radius: 50px;
} */
.overlay-text-position {
    position: relative;
    text-align: center;
    color: aliceblue;
    font-size: xx-large;
}
.centered-text {
    position: absolute;
    top: 35%;
    left: 60%;
    transform: translate(-50%, -50%);
}
</style>