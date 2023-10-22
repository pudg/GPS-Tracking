<script>
export default {
  name: 'App',
  data() {
    return {
      center: {lat: 34.049229, lng: -118.456320},
      markers: [
        {
          id: 'dfsldjl3r',
          position: {
            lat: 34.049229, lng: -118.456320
          },
        },
        {
          id: 'dfsldjl3t',
          position: {
            lat: 35.049229, lng: -118.456320
          },
        }
      ]
    }
  }
}
</script>

<template>
    <div class="bg-gray-100 h-screen vue-map" >
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

<style>
.vue-map-container .vue-map {
    height: 99.5vh;
}
</style>
