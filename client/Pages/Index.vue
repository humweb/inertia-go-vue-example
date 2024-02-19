<script setup>
import {BuildingLibraryIcon} from '@heroicons/vue/20/solid'
import Layout from "@/Layouts/Base.vue";
</script>
<script>

export default {
  name: "StudentDashboard",
  data() {
    return {
      boxes: [
        {
          title: 'Latest News',
          text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet consectetur adipiscing elit duis tristique.',
          bgColorClass: 'bg-green-600'
        }, {
          title: 'Upcoming Events',
          text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet consectetur adipiscing elit duis tristique.',
          bgColorClass: 'bg-blue-600'
        }
      ],

    }
  },
  methods: {
  },
  computed: {
  }
}
</script>
<template>
  <Layout>
    <template #header>
      <div class="flex w-full justify-between">
        <h2 class="text-gray-800 leading-tight">
          Home
        </h2>
      </div>
    </template>

    <div class="flex items-center justify-between mb-6 ml-6 sm:ml-0">
      <h2 class="text-base flex items-center font-semibold leading-7 text-gray-900">
        Recent information
      </h2>
      <a href="#" class="text-sm font-semibold leading-6 text-indigo-600 dark:text-indigo-500 hover:text-indigo-500">View
        all<span
            class="sr-only">, clients</span></a>
    </div>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 sm:gap-6 xl:grid-cols-3">
      <div v-for="box in boxes" class="card">
        <div class="card-title flex items-center">
          <div class="rounded-full text-white p-2 inline-flex mr-3" :class="box.bgColorClass">
            <BuildingLibraryIcon class="w-5 h-5"/>
          </div>
          {{ box.title }}
        </div>
        <div class="px-6 pb-7 text-gray-600 dark:text-gray-400">
          <ul class="divide-y divide-gray-100 dark:divide-gray-800">
            <li v-for="i in 3" class="py-4">
              {{ box.text }}
              <a href="#" class="block mt-2 text-sm leading-6 text-indigo-500 hover:text-indigo-500">Read more <span
                  aria-hidden="true">&rarr;</span></a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </Layout>
</template>
