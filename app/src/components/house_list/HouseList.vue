<template>
  <v-container>
    <v-row>
      <v-spacer />
      <v-btn-toggle tile color="accent-3" group>
        <v-btn v-on:click="setSortOrder('creationDatetime', 'desc')"
          >Found date<v-icon dark small> mdi-arrow-down </v-icon>
        </v-btn>
        <v-btn v-on:click="setSortOrder('creationDatetime', 'asc')"
          >Found date<v-icon dark small> mdi-arrow-up </v-icon>
        </v-btn>
        <v-btn v-on:click="setSortOrder('price', 'desc')"
          >Price<v-icon dark small> mdi-arrow-down </v-icon>
        </v-btn>
        <v-btn v-on:click="setSortOrder('price', 'asc')"
          >Price<v-icon dark small> mdi-arrow-up </v-icon>
        </v-btn>
      </v-btn-toggle>
    </v-row>
    <v-row dense>
      <v-col
        cols="4"
        v-for="house in sortedHouseList"
        v-bind:key="house.address"
      >
        <HouseItem v-bind:house="house" :height="490" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { orderBy } from "lodash/fp";
import Vue from "vue";
import { mapGetters } from "vuex";
import HouseItem from "./HouseItem.vue";

export default Vue.extend({
  name: "HouseList",
  components: { HouseItem },
  data: () => {
    return {
      column: "creationDatetime",
      order: "desc",
    };
  },
  computed: {
    ...mapGetters("house", ["filteredHouseList"]),
    sortedHouseList() {
      return orderBy([this.column], [this.order], this.filteredHouseList);
    },
  },
  methods: {
    setSortOrder(column, order) {
      this.column = column;
      this.order = order;
    },
  },
});
</script>
