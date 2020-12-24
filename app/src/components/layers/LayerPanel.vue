<template>
  <v-slide-x-transition>
    <div v-if="shouldShowLayerPanel" class="layer-overaly">
      <LayerToggles />
    </div>
  </v-slide-x-transition>
</template>

<script lang="ts">
import { mapState } from "vuex";
import LayerToggles from "./LayerToggles.vue";

export default {
  name: "LayerPanel",
  components: { LayerToggles },
  computed: {
    ...mapState("ui", ["shouldShowLayerPanel"]),
  },
};
</script>

<style>
.layer-overaly {
  position: absolute;
  width: 23rem;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
}
</style>
