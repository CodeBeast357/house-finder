<template>
  <v-card height="100%" raised class="filter-card">
    <v-container>
      <v-row>
        <v-col md="10" class="text-h5"> Filters </v-col>
        <v-col md="2">
          <v-btn icon v-on:click="setPanelValue()">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-row><v-divider /></v-row>

      <v-row>
        <v-col md="12">
          <v-select
            clearable
            label="Found date"
            :value="creationDate"
            :items="creationDateFilter"
            v-on:change="setCreationDate"
          ></v-select>
        </v-col>
      </v-row>
      <v-row>
        <v-col md="9" align-self="center">Favorites only</v-col>
        <v-col md="3">
          <v-switch
            :value="favoritesOnly"
            v-on:change="setFavoritesOnly"
          ></v-switch>
        </v-col>
      </v-row>
      <v-row>
        <v-col md="9" align-self="center">Sweet spot only</v-col>
        <v-col md="3">
          <v-switch
            :value="sweetSpotOnly"
            v-on:change="setSweetSpotOnly"
          ></v-switch>
        </v-col>
      </v-row>
      <v-row>
        <v-col md="12" align-self="center">Price</v-col>
        <v-col md="12" align-self="center"><div id="price-bucket"></div></v-col>
        <v-col md="12">
          <v-range-slider
            class="align-center"
            v-on:end="setPriceFilter"
            :value="priceFilter"
            :max="maxPrice"
            :min="minPrice"
            step="10000"
            thumb-size="40"
            thumb-label="always"
            hide-details
          >
            <template v-slot:thumb-label="{ value }">
              {{ formatRoundPrice(value) }}
            </template>
          </v-range-slider>
        </v-col>
      </v-row>
      <v-row>
        <v-col md="12" align-self="center">Black list status</v-col>
        <v-col md="12">
          <v-radio-group
            row
            mandatory
            :value="blackListStatus"
            v-on:change="setBlackListStatus"
          >
            <v-radio
              v-for="status in blackListStatusFilter"
              :key="status.value"
              :label="status.text"
              :value="status.value"
            ></v-radio>
          </v-radio-group>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script lang="ts">
import { mapGetters, mapMutations, mapState } from "vuex";
import { formatRoundPrice } from "../../services/format";
import {
  creationDateFilter,
  blackListStatusFilter,
} from "../../store/modules/filter";
import embed from "vega-embed";

export default {
  name: "FilterForm",
  data() {
    return {
      blackListStatusFilter,
      creationDateFilter,
    };
  },
  computed: {
    ...mapState("filter", [
      "blackListStatus",
      "creationDate",
      "favoritesOnly",
      "sweetSpotOnly",
    ]),
    ...mapState("house", ["houseList"]),
    ...mapGetters("house", ["minPrice", "maxPrice"]),
    priceFilter(): number[] {
      return [this.minPrice, this.maxPrice];
    },
    formatRoundPrice(): (number) => string {
      return formatRoundPrice;
    },
  },
  methods: {
    ...mapMutations("filter", [
      "setCreationDate",
      "setBlackListStatus",
      "setFavoritesOnly",
      "setSweetSpotOnly",
      "setPriceFilter",
    ]),
    ...mapMutations("ui", ["setPanelValue"]),
  },
  mounted(): void {
    embed(
      "#price-bucket",
      {
        $schema: "https://vega.github.io/schema/vega-lite/v4.json",
        description: "Price bucketing",
        mark: "bar",
        width: 290,
        height: 50,
        padding: 0,
        data: {
          values: this.houseList.map(({ price }) => ({ price })),
        },
        encoding: {
          x: {
            bin: { step: 50000 },
            field: "price",
            type: "ordinal",
            axis: null,
          },
          y: { aggregate: "count", axis: null },
          color: { value: "#808080" },
        },
        config: {
          view: {
            stroke: "transparent",
          },
        },
      },
      {
        actions: false,
      }
    );
  },
};
</script>

<style>
.filter-card {
  overflow: scroll;
}
</style>
