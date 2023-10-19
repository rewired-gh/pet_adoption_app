<template><div></div></template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useLoginStore } from '~/utils/store/loginStore'
import { useRoleStore } from '~/utils/store/roleStore'

const loginStore = useLoginStore()
const { isLogin } = storeToRefs(loginStore)
const roleStore = useRoleStore()

if (!isLogin.value) {
  navigateTo('/')
}
watch(isLogin, (newVal) => {
  if (!newVal) {
    navigateTo('/')
  }
})

loginStore.updateLogin(false)
roleStore.updateRole()
</script>
