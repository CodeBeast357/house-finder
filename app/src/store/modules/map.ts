import axios from "axios";
import { ActionTree, Module, MutationTree } from "vuex";
import { House } from "./house";

export interface BoundingBox {
  x1: number;
  y1: number;
  x2: number;
  y2: number;
}
interface State {
  selectedHouseId: number | null;
  boundingBox: BoundingBox | null;
}

const state: State = {
  selectedHouseId: null,
  boundingBox: null,
};

const getters = {
  getSelectedHouse: (
    state: State,
    getters: unknown,
    rootState: unknown,
    rootGetters: unknown
  ): House => {
    return rootGetters["house/findHouse"](state.selectedHouseId);
  },
};

const mutations: MutationTree<State> = {
  resetSelectedHouseId: (state: State): void => {
    state.selectedHouseId = null;
  },
  setSelectedHouseId: (
    state: State,
    { houseId }: { houseId: number }
  ): void => {
    state.selectedHouseId = houseId;
  },
  setBoundingBox: (state: State, boundingBox: BoundingBox): void => {
    state.boundingBox = boundingBox;
  },
};

const actions: ActionTree<State, State> = {
  resetSelectedHouseId: ({ commit }): void => {
    commit("resetSelectedHouseId");
  },
  setSelectedHouseId: ({ commit }, { houseId }: { houseId: number }): void => {
    commit("setSelectedHouseId", { houseId });
  },
  fetchBoundingBox: ({ commit }) => {
    return axios
      .get("/service/bounding-box")
      .then((response) => commit("setBoundingBox", response.data.boundingBox));
  },
};

const module: Module<State, unknown> = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};

export default module;
