<script>
import { ref } from 'vue';
import { useInfiniteScroll, useVirtualList } from '@vueuse/core'
import store from '../store';

export default {
	setup() {
		const data = ref([]);
		const { list, containerProps, wrapperProps } = useVirtualList(data, {
		itemHeight: 96,
		});
		console.log(data, list);
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
		info: []
    }
}

</script>

<template>
	<div v-bind="containerProps" class=" h-[93.5%] p-2 rounded">
		<div v-bind="wrapperProps" class="max-w-sm mx-auto">
			<!-- <div 
				v-for="{index, data} in list"
				:key="index"
				class="rounded-lg h-[120px] flex flex-col px-4 justify-center bg-white mb-2
				border-neutral-600 shadow"			>
				<h2 class="mb-2 text-2xl" style="width: 100%;">
					Device: {{ index }} {{ data }}
				</h2>
				<p class="text-lg text-red-500 ">
					info: {{ info }}
				</p>
			</div> -->
			<div 
				v-for="item in info"

				class="rounded-lg h-[150px] flex flex-col px-4 justify-center mb-2
				border-neutral-600 shadow"			>
				<h2 class="mb-2 text-lg" style="width: 100%;">
					{{ item.id }}
				</h2>
				<p>
					Lat:{{ item.position.lat }}
				</p>
				<p>
					Long:{{ item.position.lng }}
				</p>
			</div>
		</div>
	</div>
</template>

<style>

</style>