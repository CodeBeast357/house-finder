<template>
  <v-card height="100%" raised class="layer-card">
    <v-container>
      <v-row>
        <v-col md="10" class="text-h5"> Layers </v-col>
        <v-col md="2">
          <v-btn icon v-on:click="setPanelValue()">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-row><v-divider /></v-row>
      <v-row>
        <v-col md="9" align-self="center">Schools</v-col>
        <v-col md="3">
          <v-switch :value="schools" v-on:change="toggleSchools"></v-switch>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script lang="ts">
import { mapMutations, mapState } from "vuex";
export default {
  name: "LayerToggles",
  computed: {
    ...mapState("layer", ["schools"]),
  },
  methods: {
    ...mapMutations("layer", ["toggleSchools"]),
    ...mapMutations("ui", ["setPanelValue"]),
  },
};
</script>

<style>
.layer-card {
  overflow: scroll;
}
</style>
