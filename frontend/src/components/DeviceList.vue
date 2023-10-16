<script setup>
import { ref } from 'vue';
import { useInfiniteScroll, useVirtualList } from '@vueuse/core'
const data = ref(Array.from(Array(50).keys(), () => 'Billy Mays'));
const { list, containerProps, wrapperProps } = useVirtualList(data, {
  itemHeight: 96
});

useInfiniteScroll(
	containerProps.ref,
	() => {
		data.value.push(...Array.from(Array(10).keys(), () => 'Morgan Freeman'))
	},
	{
		distance: 10
	}
)

</script>

<template>
	<div v-bind="containerProps" class="h-screen p-2 rounded">
		<div v-bind="wrapperProps" class="max-w-sm mx-auto" style="max-width: 97%;">
			<div 
				v-for="{index, data} in list"
				:key="index"
				class="rounded-lg h-[120px] flex flex-col px-4 justify-center bg-white mb-2
				border-neutral-600 shadow"
				style="width: 100%;"
			>
			<h2 class="mb-2 text-2xl">
				Device: {{ index }}
			</h2>
			<p class="text-lg text-white ">
				{{ data }}
			</p>
			</div>
		</div>
	</div>
</template>

<style>
.h-screen {
	height: 85vh;
}
</style>