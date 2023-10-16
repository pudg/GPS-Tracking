<script>
export default {
  name: 'item',
  props: {
    source: {
      type: Object,
      default() {
        return {}
      }
    }
  }
}
</script>
<template>
  <div class="item">
	<div class="id">{{ source.id }} - {{ source.text }}</div>
		{{ console.log("MFA") }}
  	</div>
</template>
  
<style scoped>
.item {
	display: flex;
	flex-direction: column;
	border-bottom: 1px solid lightgrey;
	padding: 1em;
}
</style>
  