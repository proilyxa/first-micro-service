<template>
  <div class="w-96">
    <h1 class="text-center text-2xl">Register</h1>
    <form @submit.prevent="authStore.register(userReg)" class="mt-2" action="">
      <div class="flex flex-col gap-2">
        <div>
          <div class="p-inputgroup flex">
            <span class="p-inputgroup-addon">
              <i class="pi pi-user"></i>
            </span>
            <InputText v-model="userReg.firstName" type="text" placeholder="Name"/>
          </div>
          <small
              v-for="(error, index) in authStore.authErrors?.errors?.get('firstName')"
              :key="index"
              class="p-error"
              id="text-error"
          >{{ error }}
          </small>
        </div>

        <div>
          <div class="p-inputgroup flex">
            <span class="p-inputgroup-addon">
              <i class="pi pi-user"></i>
            </span>
            <InputText v-model="userReg.lastName" type="text" placeholder="Last Name"/>
          </div>

          <small
              v-for="(error, index) in authStore.authErrors?.errors?.get('lastName')"
              :key="index"
              class="p-error"
              id="text-error"
          >{{ error }}
          </small>
        </div>

        <div>
          <div class="p-inputgroup flex">
            <span class="p-inputgroup-addon">
              <i class="pi pi-user"></i>
            </span>
            <InputText v-model="userReg.email" type="text" placeholder="Email"/>
          </div>

          <small
              v-for="(error, index) in authStore.authErrors?.errors?.get('email')"
              :key="index"
              class="p-error"
              id="text-error"
          >{{ error }}
          </small>
        </div>

        <div>
          <div class="p-inputgroup flex">
            <span class="p-inputgroup-addon">
              <i class="pi pi-lock"></i>
            </span>
            <InputText v-model="userReg.password" type="password" placeholder="Password"/>
          </div>
          <small
              v-for="(error, index) in authStore.authErrors?.errors?.get('password')"
              :key="index"
              class="p-error"
              id="text-error"
          >{{ error }}
          </small>
        </div>

        <div>
          <div class="p-inputgroup flex">
            <span class="p-inputgroup-addon">
              <i class="pi pi-lock"></i>
            </span>
            <InputText
                v-model="userReg.passwordConfirm"
                type="password"
                placeholder="Confirm Password"
            />
          </div>

          <small
              v-for="(error, index) in authStore.authErrors?.errors?.get('passwordConfirm')"
              :key="index"
              class="p-error"
              id="text-error"
          >{{ error }}
          </small>
        </div>
      </div>

      <div class="flex gap-3">
        <Button type="submit" :size="'small'" class="mt-2">Register</Button>
        <a href="#" class="self-end text-blue-400">
          <router-link :to="{ name: 'login' }"> Sign In</router-link>
        </a>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import Button from 'primevue/button'
import {useAuthStore} from '@/stores/auth'
import {ref} from 'vue'
import type {UserRegistrationDto} from '@/types/user'

const userReg = ref<UserRegistrationDto>({
  firstName: null,
  lastName: null,
  email: null,
  password: null,
  passwordConfirm: null
})

const authStore = useAuthStore()
</script>

<style scoped></style>
