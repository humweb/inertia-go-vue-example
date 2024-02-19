import { createApp, h } from "vue";
import { createInertiaApp, Link, Head} from "@inertiajs/vue3";
import "./css/app.css"

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
        .mount(el);
        // app.config.performance = true;
    }
})
