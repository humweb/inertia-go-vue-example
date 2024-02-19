import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {devtools} from "vue";

// https://vitejs.dev/config/
export default defineConfig({
    content: [
        "./index.html",
        "./**/*.{vue,js,ts,jsx,tsx}",
    ],
    define: {
        // __VUE_PROD_DEVTOOLS__: 'true',
    },
    plugins: [
        vue(),
        // devtools
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./', import.meta.url))
        }
    },
    // optimizeDeps: {
    //     include: [
    //         'lodash-es',
    //         'dayjs',
    //         'vue'
    //     ],
    // },
    // build: {
    //     rollupOptions: {
    //         input: {
    //             app: './index.html', // default
    //         },
    //         output: {
    //             entryFileNames: `assets/[name].js`,
    //             chunkFileNames: `assets/[name].js`,
    //             assetFileNames: `assets/[name].[ext]`
    //         }
    //     },
    // },
})
