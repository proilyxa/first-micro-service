<template>
  <div class="about">
    <h1>This is an home page</h1>
  </div>
</template>

<script setup lang="ts">
import { useLifecycleLogger } from '@/common/hooks/lifecycleLogger'

const name = 'home'
useLifecycleLogger({ name })
</script>
