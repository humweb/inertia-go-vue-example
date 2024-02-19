import { reactive, readonly } from "vue";

const state = reactive({
    layout: 'stacked',
});

export default { state: readonly(state) };