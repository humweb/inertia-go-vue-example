<script setup>
import Layout from "@/Layouts/Base.vue";
</script>

<script>
import {HasInertiaTable, Table} from "inertia-vue-table";

export default {
  mixins: [HasInertiaTable],
  components: {
    Table,
  },
  props: {},
  computed: {
    breadcrumbs() {
      let crumbs = [];
      crumbs.push({label: 'Users'})
      return crumbs
    }
  }
}
</script>

<template>
  <Layout>
    <template #header>
      <h2 class="leading-tight">
        Users
      </h2>
    </template>
    <div class="card">
      <div class="card-header">
        <div class="text-base font-semibold leading-7">
          Users
        </div>
      </div>
      <div class="card-body">

        <Table
            :records="records"
            :pagination="pagination"
            v-model="tableData"
            class="text-base mb-0"
        />

      </div>
    </div>
  </Layout>
</template>
