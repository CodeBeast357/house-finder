import { Module, MutationTree } from "vuex";

interface State {
  schools: boolean;
}

const state: State = {
  schools: false,
};

const mutations: MutationTree<State> = {
  toggleSchools: (state: State): void => {
    state.schools = !state.schools;
  },
};

const module: Module<State, unknown> = {
  namespaced: true,
  state,
  mutations,
};

export default module;
