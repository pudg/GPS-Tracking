<script setup>
</script>

<template>
    <div>
        <!-- <div class="slogan-styles">
            <h1 class="font-extrabold">
                Navigate with Confidence, Powered by OneStep
            </h1>
        </div> -->
        <div class="info-styles flex flex-col">
            TEST
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
            <div class="text-gray-300 py-3.5 px-7 bg-red-500
                md:flex justify-between items-center">
                <div>
                    <router-link to="/login" style="font-size: large;"
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
    </div>

</template>

<style>

.info-styles {
    position: absolute;
    top: 65%;
    left: 50%;
    transform: translate(-50%, -50%);
}
</style>