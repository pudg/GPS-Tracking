<template>
    <div class="bg-gray-100" >
        <GMapMap
            :center="center"
            :zoom="7"
            map-type-id="terrain"
        >
        <GMapMarker
            :key="marker.id"
            v-for="marker in markers"
            :position="marker.position"
        />
        </GMapMap>
    </div>

</template>

<script>
import store from '../store';

export default {
  name: 'App',
  setup() {
  },
  data() {
    return {
      center: {lat: 34.049229, lng: -118.456320},
    }
  },
  props: {
    markers: []
  },
}
</script>

<style>
.vue-map-container .vue-map {
    /* height: 99.5vh; */
    height: 85vh;
}
</style>
