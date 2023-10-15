/** @type {import('tailwindcss').Config} */

const plugin = require('tailwindcss/plugin')

module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: [
          'YuFeiTe',
          '-apple-system',
          'system-ui',
          'Segoe UI',
          'Roboto',
          'Ubuntu',
          'Cantarell',
          'Noto Sans',
          'sans-serif',
          'BlinkMacSystemFont',
          'Helvetica Neue',
          'PingFang SC',
          'Hiragino Sans GB',
          'Microsoft YaHei',
          'Arial',
        ],
        yufeite: ['YuFeiTe', 'sans-serif'],
      },
    },
  },
  content: [
    './components/**/*.{js,vue,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './nuxt.config.{js,ts}',
  ],
  plugins: [
    plugin(function ({ addVariant }) {
      addVariant('hocus', ['&:hover', '&:focus'])
    }),
  ],
}
