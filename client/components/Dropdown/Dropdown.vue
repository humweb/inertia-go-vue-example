<template>
    <Menu as="div" class="relative">
        <div>
            <MenuButton ref="trigger" v-bind="{ ...$attrs}" class="flex items-center justify-between cursor-pointer">
                <slot/>
            </MenuButton>
        </div>
        <Teleport to="body">


            <MenuItems
                ref="container"
                class="z-50 absolute overflow-hidden w-56 rounded-md shadow-lg bg-white dark:bg-gray-800 dark:border dark:border-gray-800 dark:text-gray-400 ring-1 ring-black ring-opacity-5 focus:outline-none">
                <slot name="menu-items"></slot>
            </MenuItems>
        </Teleport>
    </Menu>
</template>

<script>
import {Menu, MenuButton, MenuItems} from '@headlessui/vue';
import {ChevronDownIcon} from '@heroicons/vue/24/outline';
import {usePopper} from '@/hooks/use-popper';

export default {
    inheritAttrs: false,
    name: 'DropdownMenu',
    components: {
        Menu,
        MenuButton,
        MenuItems,
        ChevronDownIcon,
    },
    props: {
        placement: {
            type: String,
            default: 'bottom-end',
        },
        strategy: {
            type: String,
            default: 'fixed',
        },
    },
    setup(props) {
        let [trigger, container] = usePopper({
            placement: props.placement,
            strategy: props.strategy,
            modifiers: [{name: 'offset', options: {offset: [0, 10]}}],
        });
        return {
            trigger,
            container,
        };
    },
};

</script>

<style scoped>

</style>
