<template>
    <div v-if="user" class="flex flex-col md:flex-row justify-center items-center">
        <div class="
            p-0 text-center
            w-full md:w-1/3 h-96 md:h-screen
            justify-center items-center
            ">
            <Preferences :allDevices="allDevices" :trackDevices="trackDevices"/>
            <DeviceList :info="devices" :updateTrackedDevices="updateTrackedDevices"/>
        </div>
        <div class="p-1 text-center w-full md:w-2/3 h-full md:h-screen">
            <GoogleMapLoader
            :trackDevices="trackDevices"
            :markers="trackedDevices"
            >
                <template slot-scope="{ google, map }">
                    {{ map }}
                    {{ google }}
                </template>
            </GoogleMapLoader>
        </div>
    </div>
</template>

<script>
import GoogleMapLoader from '../components/GoogleMapLoader.vue'
import DeviceList from '../components/DeviceList.vue';
import Preferences from '../components/Preferences.vue';
import { mapState, mapActions } from 'vuex'
import store from '../store';
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
        ...mapState(['user', 'trackedDevices', 'devices']),
        mapConfig() {
            return {
                center: {lat: 0, lng: 0}
            }
        },
    },
    methods: {
        ...mapActions(['userAuthenticate']),
        trackDevices() {
            store.commit('setTrackedDevices', store.state.devices)
        },
        allDevices() {
            store.dispatch('searchDevices');
            this.all_devices = store.state.devices;
        },
        updateTrackedDevices(id) {}
    },
}
</script>

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