<template>
  <main class="flex pt-72 justify-center h-screen">
    <slot />
  </main>
</template>

<script setup lang="ts">
import { useLifecycleLogger } from '@/common/hooks/lifecycleLogger'

useLifecycleLogger({ name: 'AuthLayout' })
</script>

<style scoped></style>
