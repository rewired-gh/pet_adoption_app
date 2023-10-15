<template><div></div></template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'

const loginStore = useLoginStore()
const { isLogin } = storeToRefs(loginStore)

if (!isLogin.value) {
  navigateTo('/')
}
watch(isLogin, (newVal) => {
  if (!newVal) {
    navigateTo('/')
  }
})

loginStore.updateLogin(false)
</script>
