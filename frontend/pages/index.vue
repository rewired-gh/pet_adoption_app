<template>
  <div
    class="flex flex-col space-y-20 w-full max-w-screen-xl items-center pt-20"
  >
    <h2
      class="text-6xl p-10 rounded-full bg-green-100 shadow-2xl shadow-lime-50"
    >
      从现在开始，领养一只宠物
    </h2>
    <section
      class="flex space-x-16 underline underline-offset-8 text-green-400 text-2xl"
    >
      <NuxtLink class="transition-all hover:scale-110" to="/">
        我要登录
      </NuxtLink>
      <NuxtLink class="transition-all hover:scale-110" to="/register">
        现在注册
      </NuxtLink>
    </section>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'

const loginStore = useLoginStore()
const { isLogin } = storeToRefs(loginStore)

if (isLogin.value) {
  navigateTo('/explore')
}
</script>
