<script setup>
const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_403261140.jpeg'
import VLazyImage from "v-lazy-image";
</script>

<template>
    <div class="overlay-text-position">
        <v-lazy-image
            :src="URL"
            style="width: 100%; height: 95vh; contain: content;"
        />
        <div class="slogan-styles">
            <h1 class="font-extrabold">
                Navigate with Confidence, Powered by OneStep
            </h1>
        </div>
        <!--TODO: restore border property-->
        <div class="info-styles m-2" > 
            <h1 class="font-extrabold m-2">
                Real-time GPS Tracking
            </h1>
            <p class="font-bold m-2">
                No Contracts   |   Cancel Anytime   |   Use It As You Need It
            </p>
            <ul>
                <li>
                    Fleets of all sizes.
                </li>
                <li>
                    Affordable Live GPS Tracking Solutions.
                </li>
                <li>
                    Powered and non-powered assets and equipment.
                </li>
                
            </ul>
            <router-link to="/login" style="font-size: large;" class="inline-block font-extrabold p-4 m-2 text-lg text-indigo-100 hover:text-white  transition-colors duration-150 bg-indigo-700 rounded-lg focus:shadow-outline hover:bg-indigo-800">TODO: Register</router-link>
            <router-link to="/login" style="font-size: large;" class="inline-block font-extrabold p-4 m-2 text-lg text-indigo-100 hover:text-white  transition-colors duration-150 bg-indigo-700 rounded-lg focus:shadow-outline hover:bg-indigo-800">Login</router-link>
        </div>
    </div>

</template>

<style>
.overlay-text-position {
    position: relative;
    text-align: center;
    color: rgb(255, 255, 255);
    /* font-size: xx-large; */
    font-size:xx-large;
}
.slogan-styles {
    position:absolute;
    top: 15%;
    left: 50%;
    max-width: 50%;
    font-size: xx-large;
    transform: translate(-50%, -50%);
}

.info-styles {
    position: absolute;
    top: 50%;
    left: 66.66%;
    max-width: 50%;
    height: 30vh;
    font-size: large;
    transform: translate(-50%, -50%);
}

</style>