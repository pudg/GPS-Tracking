<script>
import GoogleMapLoader from '../components/GoogleMapLoader.vue'
import DeviceList from '../components/DeviceList.vue';
import Preferences from '../components/Preferences.vue';
const mapConfig = {};
export default {
    components: {
    GoogleMapLoader,
    DeviceList,
    Preferences
},
    computed: {
        mapConfig() {
            return {
                center: {lat: 0, lng: 0}
            }
        },
    },
}
</script>

<template>
    <div class="flex flex-col md:flex-row justify-center items-center bg-red-500">
        <div class="
            bg-green-500 p-4 text-center
            w-full md:w-1/3 h-96 md:h-screen
            justify-center items-center
            ">
            <Preferences />
            <!-- <DeviceList /> -->
        </div>
        <div class="bg-pink-500 p-4 text-center w-full md:w-2/3 h-full md:h-screen">
            MAP
            <!-- <GoogleMapLoader :mapConfig="mapConfig" api-key="AIzaSyAEzJXuznJHAQcdCBA_HcxdRYNuA3MJuDo">
                <template slot-scope="{ google, map }">
                    {{ map }}
                    {{ google }}
                </template>
            </GoogleMapLoader> -->
        </div>
    </div>
</template>

<style>
.one-third {
  width: 25%; /* 2/3 of the screen width */
  height: 95vh; /* 100% of the viewport height */
  /* background-color: #f0f0f0; Background color, change as needed */
}

.two-thirds {
  width: 75%; /* 2/3 of the screen width */
  height: 95vh; /* 100% of the viewport height */
  /* background-color: #f0f0f0; Background color, change as needed */
}
</style>