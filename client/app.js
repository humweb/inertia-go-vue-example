import { createApp, h } from "vue";
import { createInertiaApp, Link, Head} from "@inertiajs/vue3";
import {
    vTooltip,
} from 'floating-vue'
import "./css/app.css"
import 'floating-vue/dist/style.css'
import store from '@/store'

createInertiaApp({
    resolve: name => {
        const pages = import.meta.glob('./Pages/**/*.vue', { eager: true })
        return pages[`./Pages/${name}.vue`]
    },
    setup({el, App, props, plugin}) {
        const app = createApp({
            render: () => h(App, props)
        })
        .use(plugin)
        .component("Link", Link)
        .component("Head", Head)
        .directive('tooltip', vTooltip)
        .provide('store', store)
        .mount(el);
        // app.config.performance = true;
    }
})
