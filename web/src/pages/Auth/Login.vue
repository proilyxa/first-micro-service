<template>
  <div class="w-96">
    <h1 class="text-center text-2xl">Sing In</h1>
    <form @submit.prevent="test" class="mt-2" action="">
      <div class="flex flex-col gap-2">
        <div class="p-inputgroup flex">
          <span class="p-inputgroup-addon">
            <i class="pi pi-user"></i>
          </span>
          <InputText placeholder="Email" />
        </div>

        <div class="p-inputgroup flex">
          <span class="p-inputgroup-addon">
            <i class="pi pi-lock"></i>
          </span>
          <InputText placeholder="Password" />
        </div>
      </div>

      <div class="flex gap-3">
        <Button type="submit" :size="'small'" class="mt-2">Sign In</Button>
        <a href="#" class="self-end text-blue-400">
          <router-link :to="{ name: 'register' }"> Register</router-link>
        </a>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useLifecycleLogger } from '@/common/hooks/lifecycleLogger'

const name = 'login'
import Button from 'primevue/button'
import axios from 'axios'
import { AxiosError } from 'axios'

const test = async () => {
  try {
    const res = await axios.get('/api/test')
    console.log(res.data)
  } catch (error) {
    if (error instanceof AxiosError) {
      console.log(error.response?.data)
    }
  }
}

useLifecycleLogger({ name })
</script>

<style></style>
