<script setup>
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue'
// const URL = 'https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_403261140.jpeg'
</script>

<template>
  <div class="bg-gray-100 h-auto">
	<Navbar />
	<main>
		<RouterView />
	</main>
	<Footer />
  </div>
</template>


<style>
</style>