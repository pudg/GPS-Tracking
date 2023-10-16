<script>
const URL = "https://gps-imagines.s3.us-east-2.amazonaws.com/gps.png"
const items = new Array(5)
import VirtualList from 'vue-virtual-scroll-list'
import Item from './Item.vue'
import { getData } from './data'
export default {
  name: 'App',
  data() {
    return {
      item: Item,
      items: getData(1000)
    }
  }
}
</script>

<template>
    <div id="app">
        <div class="wrapper">
			<virtual-list
				class="list"
				style="height: 95vh; overflow-y: auto;"
				:data-key="'id'"
				:data-sources="items"
				:data-component="item"
				:estimate-size="50"
			/>
		</div>
    </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 1em; */
  /* padding: 1em; */
}
.list {
  border: 2px solid steelblue;
  border-radius: 3px;
}
</style>