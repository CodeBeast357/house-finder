<template>
  <div class="map-container">
    <div class="map-icons">
      <img
        id="school-icon"
        crossOrigin="Anonymous"
        src="@/assets/school-15.svg"
        width="20"
        height="20"
      />
    </div>
    <div id="map-container" class="map" />
  </div>
</template>

<script lang="ts">
import { mapState, mapGetters, mapActions } from "vuex";
import { initMapBox } from "../../services/map/init_map";
import {
  addHouseLayer,
  addStarHouseLayer,
  addOnHouseClickHandler,
  updateFeatureState,
  updateHouseLayer,
  updateStarHouseLayer,
} from "../../services/map/layers/house_layer";
import { toggleSchoolsLayer } from "../../services/map/layers/utils";
import { addSchoolLayer } from "../../services/map/layers/schools";
import { House } from "../../store/modules/house";

export default {
  name: "Map",
  computed: {
    ...mapState("map", ["selectedHouseId", "boundingBox"]),
    ...mapState("layer", ["schools"]),
    ...mapGetters("house", ["partitionHouseBySweetSpotness"]),
  },
  methods: {
    ...mapActions("map", [
      "fetchBoundingBox",
      "setSelectedHouseId",
      "resetSelectedHouseId",
    ]),
  },
  mounted(): void {
    this.fetchBoundingBox()
      .then(() => initMapBox(this.boundingBox))
      .then((mapInstance) => (this.mapInstance = mapInstance))
      .then(() =>
        addHouseLayer(this.partitionHouseBySweetSpotness[1], this.mapInstance)
      )
      .then(() =>
        addStarHouseLayer(
          this.partitionHouseBySweetSpotness[0],
          this.mapInstance
        )
      )
      .then(() => addSchoolLayer(this.mapInstance))
      .then(() =>
        addOnHouseClickHandler(
          this.setSelectedHouseId,
          this.resetSelectedHouseId
        )(this.mapInstance)
      );
  },
  watch: {
    selectedHouseId: function (newHouseId: number, oldHouseId: number): void {
      updateFeatureState(this.mapInstance, newHouseId, oldHouseId);
    },
    partitionHouseBySweetSpotness: function (
      newHouseList: House[][],
      oldHouseList: House[][]
    ): void {
      if (this.mapInstance && newHouseList && newHouseList != oldHouseList) {
        updateHouseLayer(newHouseList[1], this.mapInstance);
        updateStarHouseLayer(newHouseList[0], this.mapInstance);
      }
    },
    schools: function (schools: boolean): void {
      toggleSchoolsLayer(schools, this.mapInstance);
    },
  },
};
</script>

<style>
.map-container {
  width: 100%;
  height: 90%;
}
.map {
  width: 100%;
  height: 100%;
}
.map-icons {
  display: none;
}
</style>
