<template>
    <MenuItem :class="styleClass">
        <component
            :is="'href' in $attrs ? 'Link' : 'button'"
            v-bind="{ ...$attrs}"
            class="flex w-full text-left py-2 px-4 rounded-md text-sm hover:bg-gray-100 dark:hover:bg-gray-700">
            <slot/>
        </component>
    </MenuItem>
</template>

<script>
import {MenuItem} from '@headlessui/vue';

export default {
    inheritAttrs: false,
    name: 'DropdownItem',
    props:{
        styleClass:{
            type: String,
            default: 'items-center'
        }
    },
    components: {
        MenuItem
    },
};
</script>

<style scoped>

</style>
