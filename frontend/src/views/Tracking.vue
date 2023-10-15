<script>
import { computed } from 'vue';
import store from '../store'
import GoogleMapLoader from '../components/GoogleMapLoader.vue'
const GOOGLE_MAPS_API_KEY = 'AIzaSyAEzJXuznJHAQcdCBA_HcxdRYNuA3MJuDo';
const devices = computed(() => store.state.devices);
const mapConfig = {};
export default {
    components: {
        GoogleMapLoader
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
    <div>
        <GoogleMapLoader :mapConfig="mapConfig" api-key="AIzaSyAEzJXuznJHAQcdCBA_HcxdRYNuA3MJuDo">
            <template slot-scope="{ google, map }">
                {{ map }}
                {{ google }}
            </template>
        </GoogleMapLoader>
    </div>
</template>