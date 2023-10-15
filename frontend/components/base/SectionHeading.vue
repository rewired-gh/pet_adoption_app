<template>
  <h2 class="flex flex-nowrap font-medium space-x-1.5" :class="styleClass">
    <span class="inline-block -translate-y-0.5">
      <slot name="icon" />
    </span>
    <span
      class="transition-colors after:transition-all relative inline-block z-10 after:block after:absolute after:bg-green-200 after:rounded-md after:w-4 after:h-2 after:-right-5 after:bottom-1 after:-z-10 hover:after:w-full hover:after:h-3 hover:after:right-0"
    >
      <slot />
    </span>
  </h2>
</template>

<script setup lang="ts">
export interface Props {
  compact?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  compact: false,
})

const styleClass = computed(() => {
  return props.compact ? 'text-xl' : 'text-3xl'
})
</script>
