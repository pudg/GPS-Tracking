<script setup>
const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_638790418.jpeg';
import Authenticate from '../components/Authenticate.vue';
</script>

<template>
    <div class="login-background-container">
        <Authenticate />
    </div>

</template>

<style>
.login-centered-text {
    max-width: 33vh;
    max-height: 25vh;
    position: absolute;
    top: 45%;
    left: 50%;
    border-radius: 16px;
    transform: translate(-50%, -50%);
}
.login-background-container {
	background-image: url('https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_638790418.jpeg');
	background-size: cover;
	background-repeat: no-repeat;
	height: 95vh;
	max-width: 100%;
}
</style>

