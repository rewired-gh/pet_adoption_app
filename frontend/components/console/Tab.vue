import { progressProps } from 'element-plus';
<template>
  <NuxtLink :to="$props.link" :custom="true">
    <template #default="{ isActive }">
      <a
        :href="$props.link"
        class="px-8 py-4 w-full transition-colors whitespace-nowrap after:transition-all relative inline-block hover:bg-green-50 after:block after:absolute after:right-0 after:top-0 after:translate-x-1 after:h-full after:w-2 rounded-l-full"
        :class="getButtonStyle(isActive)"
      >
        <slot />
      </a>
    </template>
  </NuxtLink>
</template>

<script setup lang="ts">
const props = defineProps<{
  link: string
}>()

const getButtonStyle = (isActive: boolean) => {
  return isActive
    ? 'text-green-400 after:bg-green-100'
    : 'text-slate-500 after:bg-slate-50'
}
</script>
