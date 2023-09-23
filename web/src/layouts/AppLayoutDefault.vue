<template>
  <main class="container">
    <slot />
  </main>
</template>

<script setup lang="ts">
import { useLifecycleLogger } from '@/common/hooks/lifecycleLogger'

useLifecycleLogger({ name: 'AppLayoutDefault' })
</script>

<style scoped></style>
