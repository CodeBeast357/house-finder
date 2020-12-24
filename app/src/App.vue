<template>
  <v-app id="app">
    <FilterPanel />
    <LayerPanel />
    <v-toolbar dense max-height="48px">
      <v-btn-toggle v-on:change="setPanelValue" :value="panelValue" tile group>
        <v-btn value="filter"> filters </v-btn>
        <v-btn value="layer"> Layers </v-btn>
      </v-btn-toggle>
      <v-spacer></v-spacer>
      <v-progress-circular
        v-if="isSyncLoading"
        indeterminate
        color="primary"
      ></v-progress-circular>
      <v-btn v-else icon v-on:click="syncHouses()">
        <v-icon>mdi-sync</v-icon>
      </v-btn>
    </v-toolbar>
    <v-tabs color="accent-4" grow>
      <v-tab>Map</v-tab>
      <v-tab>List ({{ houseListLength }}) </v-tab>

      <v-tab-item active-class="fill-height">
        <Map v-if="houseListLength > 0" />
      </v-tab-item>
      <v-tab-item>
        <HouseList />
      </v-tab-item>
    </v-tabs>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import FilterPanel from "./components/filters/FilterPanel.vue";
import LayerPanel from "./components/layers/LayerPanel.vue";
import HouseList from "./components/house_list/HouseList.vue";
import Map from "./components/map/Map.vue";
import { mapActions, mapGetters, mapMutations, mapState } from "vuex";

export default Vue.extend({
  name: "App",
  components: { FilterPanel, HouseList, LayerPanel, Map },
  computed: {
    ...mapState("house", ["isSyncLoading"]),
    ...mapState("ui", ["panelValue"]),
    ...mapGetters("house", ["houseListLength"]),
    ...mapGetters("house", ["houseListLength"]),
  },
  methods: {
    ...mapActions("house", ["fetchSelectedHouse", "syncHouses"]),
    ...mapMutations("ui", ["setPanelValue"]),
  },
  mounted() {
    this.fetchSelectedHouse();
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
}

.fill-height {
  height: 100%;
}

.v-tabs-items {
  height: 100%;
}
</style>
