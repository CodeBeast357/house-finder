import { Module, MutationTree } from "vuex";

interface State {
  shouldShowFilterPanel: boolean;
}

const state: State = {
  shouldShowFilterPanel: false,
};

const mutations: MutationTree<State> = {
  setShouldShowFilterPanel: (
    state: State,
    shouldShowFilterPanel: boolean
  ): void => {
    state.shouldShowFilterPanel = shouldShowFilterPanel;
  },
};

const module: Module<State, unknown> = {
  namespaced: true,
  state,
  mutations,
};

export default module;
