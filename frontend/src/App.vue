<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue';
</script>

<template>
  <div class="bg-gray-100 h-full">
	<Navbar />
	<main>
		<RouterView />
	</main>
  </div>
</template>

<style scoped>
</style>
