<template>
  <header
    class="whitespace-nowrap flex items-center space-x-12 px-10 h-16 z-50 w-full shadow-sm backdrop-blur-md backdrop-saturate-200 bg-white/70"
  >
    <NuxtLink class="justify-start" to="/">
      <div class="flex items-center">
        <div
          class="flex items-center justify-center mr-3 content-stretch h-10 w-10"
        >
          <slot name="logo" />
        </div>
        <span class="text-xl font-medium text-slate-800">
          {{ $props.titleName }}
        </span>
      </div>
    </NuxtLink>
    <div class="grow" />
    <nav class="justify-end font-medium text-base">
      <ul class="flex flex-row list-none space-x-8">
        <li>
          <NavButton v-if="isLogin" link="/explore"> 主页 </NavButton>
          <NavButton v-else link="/"> 主页 </NavButton>
        </li>
        <li v-if="isLogin">
          <NavButton link="/logout"> 登出 </NavButton>
        </li>
        <li v-if="!isLogin">
          <NavButton link="/register"> 注册 </NavButton>
        </li>
        <li v-if="isLogin && isAdmin">
          <NavButton link="/console/pets"> 控制台 </NavButton>
        </li>
        <li>
          <NavButton link="http://rewired.moe" :is-external-link="true">
            关于我们
          </NavButton>
        </li>
        <li v-if="isLogin">
          <span class="text-slate-400"> 您好，{{ username }} </span>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'
import { useRoleStore } from '~/utils/store/roleStore'

defineProps<{
  titleName: string
}>()

const loginStore = useLoginStore()
const { isLogin, username } = storeToRefs(loginStore)
const roleStore = useRoleStore()
const { roles } = storeToRefs(roleStore)

var adminSet = new Set(['user_admin', 'reviewer'])
const isAdmin = computed(() => {
  return roles.value.filter((role) => adminSet.has(role)).length > 0
})
</script>
~/utils/store/roleStore ~/utils/store/roleStore
