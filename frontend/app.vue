<template>
  <div>
    <NuxtLayout class="bg-inherit">
      <NuxtPage class="bg-inherit" />
    </NuxtLayout>
  </div>
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
