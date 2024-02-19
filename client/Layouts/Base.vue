<script setup>
import {Link} from '@inertiajs/vue3'
import {Disclosure, DisclosureButton, DisclosurePanel} from '@headlessui/vue'
import {Bars3Icon, BellIcon, FireIcon, UserIcon, XMarkIcon} from '@heroicons/vue/24/outline'

</script>

<script>
import {navigation} from "@/store/navigation.js";

export default {
  props: {
    headerWidth: {
      type: String,
      default: 'full'
    },
    bodyWidth: {
      type: String,
      default: 'centered'
    }
  },
  data() {
    return {
      user: {},
      navigation
    }
  },
  methods: {
    logout() {
      alert('Logout not setup');
    },
    isUrl(...urls) {
      let currentUrl = this.$page.url.substr(1)
      if (urls[0] === '') {
        return currentUrl === ''
      }
      return urls.filter((url) => currentUrl.startsWith(url)).length
    },
  },
  computed: {}
}
</script>

<template>
  <div class="min-h-full bg-gray-100 text-gray-900 dark:bg-gray-900 dark:text-gray-200">
    <Disclosure as="nav" class="border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-900"
                v-slot="{ open }">
      <div :class="[headerWidth == 'full' ? '' : 'max-w-7xl', 'mx-auto px-4 sm:px-6 lg:px-8']">
        <div class="flex h-16 justify-between">
          <div class="flex">
            <div class="flex flex-shrink-0 items-center logo text-lg font-bold">
              <FireIcon class="w-9 h-9 inline mr-1 text-indigo-700 dark:text-indigo-500"/>
              Go + Inertia + VueTables
            </div>
            <div class="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8">
              <Link v-for="item in navigation" :key="item.name" :href="item.href"
                    :class="[isUrl(item.current) ? 'border-indigo-500 text-gray-900 dark:text-indigo-300' : 'border-transparent text-gray-500 dark:text-gray-200 hover:border-gray-300 hover:text-gray-700 dark:hover:text-indigo-300', 'inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium']"
                    :aria-current="isUrl(item.current) ? 'page' : undefined">{{ item.name }}
              </Link>
            </div>
          </div>


          <div class="-mr-2 flex items-center sm:hidden">
            <!-- Mobile menu button -->
            <DisclosureButton
                class="relative inline-flex items-center justify-center rounded-md bg-white dark:bg-gray-800 p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
              <span class="absolute -inset-0.5"/>
              <span class="sr-only">Open main menu</span>
              <Bars3Icon v-if="!open" class="block h-6 w-6" aria-hidden="true"/>
              <XMarkIcon v-else class="block h-6 w-6" aria-hidden="true"/>
            </DisclosureButton>
          </div>
        </div>
      </div>

      <DisclosurePanel class="sm:hidden">
        <div class="space-y-1 pb-3 p-2">
          <DisclosureButton v-for="item in navigation" :key="item.name" as="a" :href="item.href"
                            :class="[isUrl(item.current) ? ' bg-indigo-500 text-indigo-50' : 'text-gray-600 dark:text-gray-400 dark:hover:text-gray-700 hover:bg-indigo-100 hover:text-gray-800', 'block rounded-md py-4 pl-3 pr-4 text-base font-medium']"
                            :aria-current="item.current ? 'page' : undefined">{{ item.name }}
          </DisclosureButton>
        </div>

      </DisclosurePanel>
    </Disclosure>
    <header class="bg-white dark:bg-gray-800 shadow-sm dark:shadow-none">
      <div :class="[headerWidth == 'full' ? '' : 'max-w-7xl', 'mx-auto py-4 px-4 sm:px-6 lg:px-8']">
        <slot name="header"/>
      </div>
    </header>
    <div class="py-10 h-full">
      <main>
        <div :class="[bodyWidth == 'full' ? '' :  'max-w-7xl', 'mx-auto sm:px-6 lg:px-8']">
          <slot/>
        </div>
      </main>
    </div>
  </div>

</template>
