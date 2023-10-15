<template>
  <div class="flex flex-col text-slate-800">
    <Navbar title-name="Rewired 宠物领养控制台" class="fixed top-0 left-0">
      <template #logo>
        <Icon class="text-3xl" name="fluent-emoji:hammer-and-wrench" />
      </template>
    </Navbar>
    <main class="flex pt-24 h-screen min-w-fit bg-inherit">
      <ul class="pl-4 w-fit">
        <li class="w-full">
          <ConsoleTab link="/console/pets">管理宠物信息</ConsoleTab>
        </li>
        <li class="w-full">
          <ConsoleTab link="/">管理用户</ConsoleTab>
        </li>
      </ul>
      <div class="flex flex-grow justify-around px-10">
        <slot></slot>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'

const loginStore = useLoginStore()
const { isLogin } = storeToRefs(loginStore)

if (!isLogin.value) {
  navigateTo('/')
}
</script>
