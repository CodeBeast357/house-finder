import { Module, MutationTree } from "vuex";

interface State {
  panelValue: string;
  shouldShowFilterPanel: boolean;
  shouldShowLayerPanel: boolean;
}

const state: State = {
  panelValue: undefined,
  shouldShowFilterPanel: false,
  shouldShowLayerPanel: false,
};

const mutations: MutationTree<State> = {
  setPanelValue: (state: State, panelValue: string): void => {
    state.panelValue = panelValue;
    state.shouldShowFilterPanel = panelValue == "filter";
    state.shouldShowLayerPanel = panelValue == "layer";
  },
};

const module: Module<State, unknown> = {
  namespaced: true,
  state,
  mutations,
};

export default module;
