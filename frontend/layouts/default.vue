<template>
  <div class="flex flex-col text-slate-800">
    <Navbar title-name="Rewired 宠物领养平台" class="fixed top-0 left-0">
      <template #logo>
        <Icon class="text-3xl" name="fluent-emoji:baby-chick" />
      </template>
    </Navbar>
    <div
      class="flex flex-col justify-between pt-12 px-12 min-h-screen min-w-fit bg-inherit"
    >
      <main class="flex justify-around px-auto py-12">
        <slot />
      </main>
      <Footer class="mt-6" />
    </div>
  </div>
</template>
