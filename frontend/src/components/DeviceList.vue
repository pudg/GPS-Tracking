<template>
	<div v-bind="containerProps" class=" h-[93.5%] p-2 rounded">
		<div v-bind="wrapperProps" class="max-w-sm mx-auto">
			<div 
				v-for="item in info"

				class="rounded-lg flex flex-col px-4 py-2 justify-center mb-2
				border-neutral-600 shadow"			>
				<div class="flex flex-col justify-center items-center">
					<h2 class="text-lg" style="width: 100%;">
						{{ item.make }}: {{ item.model }}
					</h2>
					<p>
						Active: {{ item.active }}
					</p>
					<p>
						Lat: {{ item.position.lat.toFixed(3) }}
						Long: {{ item.position.lng.toFixed(3) }}
					</p>
					<button @click="toggleHide(item.id)" class="border bg-indigo-500 hover:bg-indigo-700 text-white p-1">
						{{ !item.hide ? "Hide" : "Show"  }}
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import { ref } from 'vue';
import { mapActions, useStore, mapState } from 'vuex';
import { useInfiniteScroll, useVirtualList } from '@vueuse/core'
import store from '../store';

export default {
	setup() {
		const data = ref([]);
		const store = useStore();
		const { list, containerProps, wrapperProps } = useVirtualList(data, {
		itemHeight: 96,
		});
		useInfiniteScroll(
			containerProps.ref,
			() => {},
			{
				distance: 10
			}
		)
		return {data, list, containerProps, wrapperProps}
	},
	components: {
        useInfiniteScroll,
        useVirtualList,
    },
	props: {
		info: [],
		updateTrackedDevices: Function,
    },
	methods: {
		toggleHide(id) {
			store.dispatch('hideDevice', id);
		},
		...mapActions(['hideDevice']),
	},
	computed: {
		...mapState(['user', 'loginError']),
	},
}

</script>

<style>

</style>