// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@pinia/nuxt',
    '@nuxtjs/color-mode',
    'nuxt-icon',
    'dayjs-nuxt',
    '@element-plus/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
  ],
  ssr: true,
  colorMode: {
    preference: 'light',
  },
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/assets/css/element/index.scss" as element;`,
        },
      },
    },
  },
  elementPlus: {
    icon: 'ElIcon',
    importStyle: 'scss',
  },
  // ui: {
  // icons: 'all',
  // },
})
