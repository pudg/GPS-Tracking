<script>
import GoogleMapLoader from '../components/GoogleMapLoader.vue'
import DeviceList from '../components/DeviceList.vue';
import Preferences from '../components/Preferences.vue';
import { mapState, mapActions } from 'vuex'
import store from '../store';
import Login from './Login.vue'
const mapConfig = {};
export default {
    components: {
        GoogleMapLoader,
        DeviceList,
        Preferences
    },
    data() {
        return {
            devices_track: 0,
            all_devices: [],
            track_devices: [],
            localStore: store,
        }
    },
    computed: {
        ...mapState(['user']),
        mapConfig() {
            return {
                center: {lat: 0, lng: 0}
            }
        },
    },
    methods: {
        ...mapActions(['userAuthenticate']),
        trackDevices() {
            this.track_devices = store.state.devices;
        },
        allDevices() {
            store.dispatch('searchDevices');
            this.all_devices = store.state.devices;
        }
    },
}
</script>

<template>
    <div v-if="user" class="flex flex-col md:flex-row justify-center items-center">
        <div class="
            p-0 text-center
            w-full md:w-1/3 h-96 md:h-screen
            justify-center items-center
            ">
            <Preferences :allDevices="allDevices" :trackDevices="trackDevices" />
            <DeviceList :info="localStore.state.devices"/>
        </div>
        <div class="p-1 text-center w-full md:w-2/3 h-full md:h-screen">
            <GoogleMapLoader
            :trackDevices="trackDevices"
            :markers="track_devices"
            api-key="AIzaSyAEzJXuznJHAQcdCBA_HcxdRYNuA3MJuDo">
                <template slot-scope="{ google, map }">
                    {{ map }}
                    {{ google }}
                </template>
            </GoogleMapLoader>
        </div>
    </div>
</template>

<style>
.one-third {
  width: 25%;
  height: 95vh;
}

.two-thirds {
  width: 75%;
  height: 95vh;
}
</style>