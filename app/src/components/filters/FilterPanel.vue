<template>
  <v-slide-x-transition>
    <div v-if="shouldShowFilterPanel" class="filter-overaly">
      <FilterForm />
    </div>
  </v-slide-x-transition>
</template>

<script lang="ts">
import { mapState } from "vuex";
import FilterForm from "./FilterForm.vue";

export default {
  name: "FilterPanel",
  components: { FilterForm },
  computed: {
    ...mapState("ui", ["shouldShowFilterPanel"]),
  },
};
</script>

<style>
.filter-overaly {
  position: absolute;
  width: 23rem;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
}

.filter-card {
  overflow: scroll;
}
</style>
