import Vue from "vue";
import Vuex, { StoreOptions } from "vuex";
import filter from "./modules/filter";
import house from "./modules/house";
import map from "./modules/map";
import ui from "./modules/ui";

Vue.use(Vuex);

const storeOptions: StoreOptions<unknown> = {
  modules: { filter, house, map, ui },
  plugins: [Vuex.createLogger()],
};

export const store = new Vuex.Store<unknown>(storeOptions);
