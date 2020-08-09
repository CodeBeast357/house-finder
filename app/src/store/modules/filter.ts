import { Module, MutationTree } from "vuex";
import {
  BLACKLIST_ALL_FILTER,
  BLACKLIST_BLACK_FILTER,
  BLACKLIST_WHITE_FILTER,
  CREATION_DATE_3DAYS_FILTER,
  CREATION_DATE_7DAYS_FILTER,
  CREATION_DATE_LAST_FOUND_FILTER,
  CREATION_DATE_MORE_7DAYS_FILTER,
} from "../constants";

interface State {
  blackListStatus: string;
  creationDate: string;
  favoritesOnly: boolean;
  minPriceFilter: number;
  maxPriceFilter: number;
  sweetSpotOnly: boolean;
}

export const creationDateFilter = [
  {
    text: "Last found houses",
    value: CREATION_DATE_LAST_FOUND_FILTER,
  },
  {
    text: "Last 3 days",
    value: CREATION_DATE_3DAYS_FILTER,
  },
  {
    text: "Last 7 days",
    value: CREATION_DATE_7DAYS_FILTER,
  },
  {
    text: "More than 7 days",
    value: CREATION_DATE_MORE_7DAYS_FILTER,
  },
];

export const blackListStatusFilter = [
  {
    text: "White listed",
    value: BLACKLIST_WHITE_FILTER,
  },
  {
    text: "Black listed",
    value: BLACKLIST_BLACK_FILTER,
  },
  {
    text: "All",
    value: BLACKLIST_ALL_FILTER,
  },
];

const state: State = {
  blackListStatus: BLACKLIST_WHITE_FILTER,
  creationDate: null,
  favoritesOnly: false,
  minPriceFilter: null,
  maxPriceFilter: null,
  sweetSpotOnly: false,
};

const mutations: MutationTree<State> = {
  setCreationDate: (state: State, creationDate: string): void => {
    state.creationDate = creationDate;
  },
  setBlackListStatus: (state: State, blackListStatus: string): void => {
    state.blackListStatus = blackListStatus;
  },
  setFavoritesOnly: (state: State, favoritesOnly: boolean): void => {
    state.favoritesOnly = favoritesOnly;
  },
  setSweetSpotOnly: (state: State, sweetSpotOnly: boolean): void => {
    state.sweetSpotOnly = sweetSpotOnly;
  },
  setPriceFilter: (
    state: State,
    [minPriceFilter, maxPriceFilter]: number[]
  ): void => {
    state.minPriceFilter = minPriceFilter;
    state.maxPriceFilter = maxPriceFilter;
  },
};

const module: Module<State, unknown> = {
  namespaced: true,
  state,
  mutations,
};

export default module;
