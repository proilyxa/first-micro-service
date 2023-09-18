<template>
  <main class="container">
    <div>AppLayoutDefault</div>
    <slot/>
  </main>
</template>

<script setup lang="ts">
import {useLifecycleLogger} from '@/common/hooks/lifecycleLogger'

useLifecycleLogger({name: 'AppLayoutDefault'})

</script>

<style scoped></style>
