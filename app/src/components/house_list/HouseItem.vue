<template>
  <v-card class="mx-auto" :height="height">
    <v-img height="50%" :src="house.thumbnailLink">
      <v-app-bar flat color="rgba(0, 0, 0, 0)">
        <v-spacer></v-spacer>
        <v-icon color="#B9F6CA" v-if="house.isInSweetSpot"> star </v-icon>
      </v-app-bar>
    </v-img>

    <v-card-subtitle>
      <v-row>
        <v-col cols="10" class="text-h5">
          {{ price }}
        </v-col>
        <v-col cols="2">
          <a :href="house.link" target="_blank" class="text-decoration-none"
            ><v-img :src="providerIcon"></v-img></a
        ></v-col>
      </v-row>
    </v-card-subtitle>

    <v-card-text class="text--primary">
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <div v-bind="attrs" v-on="on" class="text-truncate">
            {{ house.address }}
          </div>
        </template>
        <span>{{ house.address }}</span>
      </v-tooltip>
      <div class="my-4 caption">
        Found the {{ new Date(house.creationDatetime).toLocaleString() }}
      </div>
    </v-card-text>
    <v-card-actions>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            icon
            v-on:click="toggleIsBlackListed(house)"
            :color="house.isBlackListed ? 'red' : ''"
          >
            <v-icon>mdi-cancel</v-icon>
          </v-btn>
        </template>
        <span>Put house on blacklist</span>
      </v-tooltip>

      <v-spacer />
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            v-bind="attrs"
            v-on="on"
            icon
            v-on:click="toggleIsFavorite(house)"
            :color="house.isFavorite ? 'pink' : ''"
          >
            <v-icon>mdi-heart</v-icon>
          </v-btn>
        </template>
        <span>Put house as favorite</span>
      </v-tooltip>
      <v-btn icon :href="googleMapsLink" target="_blank">
        <v-icon>map</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions } from "vuex";
import { getHouseLogo, getGoogleMapsLink } from "./../../selectors/house";
import { formatPrice } from "./../../services/format";

export default Vue.extend({
  name: "HouseItem",
  props: ["house", "width", "height"],
  computed: {
    price: function () {
      return formatPrice(this.house.price);
    },
    providerIcon: function () {
      const houseLogo = getHouseLogo(this.house.providerName);
      return require(`@/${houseLogo}`);
    },
    googleMapsLink: function () {
      return getGoogleMapsLink(this.house);
    },
  },
  methods: {
    ...mapActions("house", ["toggleIsBlackListed", "toggleIsFavorite"]),
  },
});
</script>
