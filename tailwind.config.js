// tailwind.config.js

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./web/**/*.{go,js,templ,html}'],
  theme: {
    extend: {},
  },
  plugins: ['@tailwindcss/postcss'],
}
