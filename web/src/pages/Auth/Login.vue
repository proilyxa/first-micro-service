<template>
  <div>
    <h1>This is an auth page</h1>
  </div>
</template>

<script setup lang="ts">
import { useLifecycleLogger } from '@/common/hooks/lifecycleLogger'

const name = 'login'
useLifecycleLogger({ name })

</script>

<style>

</style>
