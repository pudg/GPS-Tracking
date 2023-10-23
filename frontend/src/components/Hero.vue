<script setup>
</script>

<template>
    <div class="slogan-styles w-full flex flex-col text-white bg-green-0 text-center">
        <h1 class="font-extrabold bg-red-0 mt-24 p-2 text-white">
            Navigate with Confidence, Powered by OneStep
        </h1>
        <p class="text-lg my-4">
            No Contracts   |   Cancel Anytime   |   Use It As You Need It
        </p>
        <ul class="text-lg">
            <li>
                Fleets of all sizes.
            </li>
            <li class="p-1">
                Affordable Live GPS Tracking Solutions.
            </li>
            <li class="p-1">
                Powered and non-powered assets and equipment.
            </li>
            
        </ul>
        <div class="top-3/4 my-4 flex flex-col md:flex-row justify-center items-center">
            <!-- <h1>
                Real-time GPS Tracking
            </h1>
            <p>
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
                
            </ul> -->
            <div>
                <router-link to="/signup" style="font-size: large;"
                class="inline-block font-extrabold p-4 m-1 text-lg
                text-indigo-100 hover:text-white  transition-colors
                duration-150 bg-indigo-700 rounded-lg focus:shadow-outline
                hover:bg-indigo-800 text-center">Sign Up</router-link>
            </div>
            <div>
                <router-link to="/login" style="font-size: large;"
                class="inline-block font-extrabold p-4 m-1 text-lg
                text-indigo-100 hover:text-white  transition-colors
                duration-150 bg-indigo-700 rounded-lg focus:shadow-outline
                hover:bg-indigo-800 text-center">Login</router-link>
            </div>
        </div>
    </div>
</template>

<style>
.slogan-styles {
    /* position: absolute; */
    top: 20%
}
</style>