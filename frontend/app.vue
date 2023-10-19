<template>
  <NuxtLayout class="bg-inherit">
    <NuxtPage class="bg-inherit" />
  </NuxtLayout>
</template>

<script setup lang="ts">
useRoute()
useHead({
  title: 'Rewired 宠物领养平台',
  meta: [
    {
      name: 'viewport',
      content: 'width=device-width, initial-scale=1',
    },
  ],
})
</script>

<style lang="scss">
.layout-enter-active,
.layout-leave-active,
.page-enter-active,
.page-leave-active {
  transition: all 0.3s;
}

.layout-enter-from,
.layout-leave-to,
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(1em);
}
</style>
