<template>
  <div
    class="flex flex-col space-y-20 w-full max-w-screen-xl items-center pt-14"
  >
    <h2
      ref="board"
      class="text-6xl p-10 rounded-full bg-green-100 shadow-2xl shadow-lime-50"
      :style="{ translate: boardTranslate }"
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
    const uid: number = res.Uid
    loginStore.updateLogin(true)
    loginStore.updateUid(uid)
    loginStore.updateUsername(req.Username)
  }
  isLogining.value = false
}

const board: Ref<HTMLElement | null> = ref(null)
const mouseRect: Ref<{ x: number; y: number } | null> = ref(null)

const mouseEvent = (ev: MouseEvent) => {
  mouseRect.value = {
    x: ev.clientX,
    y: ev.clientY,
  }
}
if (process.browser) {
  window.addEventListener('mousemove', mouseEvent)
  onUnmounted(() => {
    window.removeEventListener('mousemove', mouseEvent)
  })
}

const boardRect = computed(() => {
  return board.value?.getBoundingClientRect()
})

const boardTranslate = computed(() => {
  if (!boardRect.value || !mouseRect.value) {
    return 'none'
  }
  const x = boardRect.value.left + boardRect.value.width / 2
  const y = boardRect.value.top + boardRect.value.height / 2
  const odx = mouseRect.value.x - x
  const ody = mouseRect.value.y - y
  const dx = (Math.sign(odx) * Math.sqrt(Math.abs(odx))) / 4
  const dy = (Math.sign(ody) * Math.sqrt(Math.abs(ody))) / 4
  console.log(`${dx}px ${dy}px`)
  return `${dx}px ${dy}px`
})
</script>
