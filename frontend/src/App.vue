<template>
  <div class="bg-gray-100 h-auto">
	<Navbar />
	<main>
		<RouterView />
	</main>
	<Footer />
  </div>
</template>

<script>
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue'
import { useStore } from 'vuex';
export default {
	components: {
		Navbar, Footer, RouterView
	},
	setup() {
		const store = useStore();
		const user = localStorage.getItem("userID");
		if (user) {
			store.commit("setUser", JSON.parse(user));
		}
		const devices = localStorage.getItem("userDevices");
		if (devices) {
			store.commit("reloadDeviceList", JSON.parse(devices));
		}
	},
}
</script>