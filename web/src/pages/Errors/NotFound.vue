<template>
  <div>
    <p>Not found</p>
    <RouterLink to="/">Home</RouterLink>
  </div>

</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {RouterLink} from "vue-router";

export default defineComponent({
  name: "NotFound",
  components: {RouterLink}
})
</script>

<style scoped>

</style>