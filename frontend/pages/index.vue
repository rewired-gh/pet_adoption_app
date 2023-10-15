<template>
  <div
    class="flex flex-col space-y-20 w-full max-w-screen-xl items-center pt-14"
  >
    <h2
      class="text-6xl p-10 rounded-full bg-green-100 shadow-2xl shadow-lime-50"
    >
      从现在开始，领养一只宠物
    </h2>
    <section class="flex items-end space-x-8">
      <BaseWithLabel>
        <template #label>
          <span> 用户名 </span>
        </template>
        <ElInput v-model="username" size="large" />
      </BaseWithLabel>
      <BaseWithLabel>
        <template #label>
          <span> 密码 </span>
        </template>
        <ElInput v-model="pword" type="password" size="large" />
      </BaseWithLabel>
      <ElButton
        type="primary"
        plain
        size="large"
        round
        :disabled="!isInfoValid || isLogining"
        :loading="isLogining"
        @click="onLogin()"
      >
        登录
      </ElButton>
    </section>
    <NuxtLink
      class="text-slate-400 underline underline-offset-4"
      to="/register"
    >
      还没有账户？现在注册吧！
    </NuxtLink>
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
watch(isLogin, (newVal) => {
  if (newVal) {
    navigateTo('/explore')
  }
})

const username = ref()
const pword = ref()
const isLogining = ref(false)
const isInfoValid = computed(() => {
  return username.value && pword.value
})

const onLogin = async () => {
  isLogining.value = true
  const req = {
    Username: username.value,
    Pword: pword.value,
  }
  const res = await fetchR('/user/auth', req)
  if (res) {
    const uid: number = (await res.json()).Uid
    loginStore.updateLogin(true)
    loginStore.updateUid(uid)
    loginStore.updateUsername(req.Username)
  }
  isLogining.value = false
}
</script>
