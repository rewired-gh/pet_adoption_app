<template>
  <NuxtLink :to="$props.link" :custom="true">
    <template #default="{ isActive }">
      <a
        :href="$props.link"
        :target="target"
        class="transition-colors after:transition-all relative inline-block hover:text-green-400 after:block after:absolute after:left-1/2 after:-bottom-2 after:-translate-x-1/2 after:rounded-full after:-z-10 hover:after:w-full hover:after:h-2 hover:after:bg-slate-100"
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
  isExternalLink?: boolean
}>()

const getButtonStyle = (isActive: boolean) => {
  return isActive
    ? 'text-green-400 after:w-full after:h-2 after:bg-green-400/20'
    : 'text-slate-500 after:bg-green-400/20 after:w-0 after:h-0'
}

const target = computed(() => {
  return props.isExternalLink ? '_blank' : '_self'
})
</script>
