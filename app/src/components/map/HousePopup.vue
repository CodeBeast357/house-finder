<template>
  <v-app style="height: 440px; width: 250px" v-if="selectedHouseId">
    <HouseItem :house="getSelectedHouse" :height="440" />
  </v-app>
</template>

<script lang="ts">
import { isEqual } from "lodash/fp";
import { mapGetters, mapMutations, mapState } from "vuex";
import HouseItem from "./../house_list/HouseItem.vue";
import { destroyMapPopup } from "../../services/map/layers/house_layer";

export default {
  name: "HousePopup",
  components: { HouseItem },
  computed: {
    ...mapGetters("map", ["getSelectedHouse"]),
    ...mapState("map", ["selectedHouseId"]),
  },
  methods: {
    handleEscape(event: KeyboardEvent): void {
      if (isEqual(event.key, "Escape")) {
        destroyMapPopup();
        this.resetSelectedHouseId();
      }
    },
    ...mapMutations("map", ["resetSelectedHouseId"]),
  },
  mounted(): void {
    document.addEventListener("keydown", (event) => this.handleEscape(event));
  },
};
</script>

<style></style>
