<template>
  <main class="">
    <slot/>
  </main>
</template>

<script setup lang="ts">
import {useLifecycleLogger} from '@/common/hooks/lifecycleLogger'

useLifecycleLogger({name: 'AuthLayout'})
</script>

<style scoped>


</style>
