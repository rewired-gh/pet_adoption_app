<template>
  <div class="flex flex-col space-y-[2px]">
    <label class="text-slate-400 text-sm pl-2">
      <slot name="label"></slot>
    </label>
    <slot></slot>
  </div>
</template>
