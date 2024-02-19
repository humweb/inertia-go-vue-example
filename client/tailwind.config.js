import defaultTheme from 'tailwindcss/defaultTheme';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    '../server/index.html',
    './**/*.vue',
    './node_modules/inertia-vue-table/**/*.vue',
    './node_modules/inertia-vue-table/**/*.js',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: [...defaultTheme.fontFamily.sans],
      },
    },
  },

  plugins: [
    forms,
    typography
  ],
};
