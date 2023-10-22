<script setup>
import Hero from '../components/Hero.vue'
</script>

<template>
	<div class="home-background-container h-screen">
		<Hero />
	</div>
</template>

<style>
.home-background-container {
	background-image: url('https://gps-imagines.s3.us-east-2.amazonaws.com/AdobeStock_403261140.jpeg'); /* Replace with the actual path or URL to your image */
	background-size: cover;
	background-repeat: no-repeat;
	/* height: 100%; */
	max-width: 100%;
}
</style>