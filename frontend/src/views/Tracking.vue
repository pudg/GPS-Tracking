<script>
import GoogleMapLoader from '../components/GoogleMapLoader.vue'
import DeviceList from '../components/DeviceList.vue';
const mapConfig = {};
export default {
    components: {
    GoogleMapLoader,
    DeviceList
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
    <div class="flex">
        <div class="one-third-component">
            <DeviceList />
        </div>
        <div class="two-thirds-component border border-red-600">
            <GoogleMapLoader :mapConfig="mapConfig" api-key="AIzaSyAEzJXuznJHAQcdCBA_HcxdRYNuA3MJuDo">
                <template slot-scope="{ google, map }">
                    {{ map }}
                    {{ google }}
                </template>
            </GoogleMapLoader>
        </div>
    </div>
</template>
<style>
.one-third-component {
  width: 25%; /* 2/3 of the screen width */
  height: 95vh; /* 100% of the viewport height */
  background-color: #f0f0f0; /* Background color, change as needed */
}

.two-thirds-component {
  width: 75%; /* 2/3 of the screen width */
  height: 95vh; /* 100% of the viewport height */
  background-color: #f0f0f0; /* Background color, change as needed */
}
</style>