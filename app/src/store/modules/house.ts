import axios from "axios";
import {
  ActionContext,
  ActionTree,
  GetterTree,
  Module,
  MutationTree,
} from "vuex";
import {
  cond,
  find,
  flow,
  get,
  identity,
  isEqual,
  minBy,
  maxBy,
  overEvery,
  partition,
  stubTrue,
} from "lodash/fp";
import {
  BLACKLIST_BLACK_FILTER,
  BLACKLIST_WHITE_FILTER,
  CREATION_DATE_3DAYS_FILTER,
  CREATION_DATE_7DAYS_FILTER,
  CREATION_DATE_LAST_FOUND_FILTER,
} from "../constants";

export type ProviderName = "centris" | "duproprio" | "remax";
export interface House {
  id: number;
  address: string;
  arrondissement: string;
  creationDatetime: number;
  isBlackListed: boolean;
  isFavorite: boolean;
  isInSweetSpot: boolean;
  latitude: number;
  link: string;
  longitude: number;
  price: number;
  providerName: ProviderName;
  thumbnailLink: string;
}

interface State {
  houseList: House[];
  isSyncLoading: boolean;
}

const state: State = {
  houseList: [],
  isSyncLoading: false,
};

const filterIfTrue = cond([
  [identity, () => isEqual(true)],
  [stubTrue, () => stubTrue],
]);

const isLastFound = isEqual(CREATION_DATE_LAST_FOUND_FILTER);
const isLast3Days = isEqual(CREATION_DATE_3DAYS_FILTER);
const isLast7Days = isEqual(CREATION_DATE_7DAYS_FILTER);

const filterBlackList = cond([
  [isEqual(BLACKLIST_WHITE_FILTER), () => isEqual(false)],
  [isEqual(BLACKLIST_BLACK_FILTER), () => isEqual(true)],
  [stubTrue, () => stubTrue],
]);
const getters: GetterTree<State, any> = {
  filteredHouseList: (state, getters, rootState): House[] => {
    const newestCreationDatetime = get(
      "creationDatetime",
      maxBy("creationDatetime", state.houseList)
    );
    const filterCreationDatetime = cond([
      [isLastFound, () => isEqual(newestCreationDatetime)],
      [
        isLast3Days,
        () => (creationDatetime) =>
          newestCreationDatetime - creationDatetime <= 3600 * 24 * 1000 * 3,
      ],
      [
        isLast7Days,
        () => (creationDatetime) =>
          newestCreationDatetime - creationDatetime <= 3600 * 24 * 1000 * 7,
      ],
      [stubTrue, () => stubTrue],
    ]);

    const filterByCreationDatetime = flow(
      get("creationDatetime"),
      filterCreationDatetime(rootState.filter.creationDate)
    );
    const filterByFavoritesOnly = flow(
      get("isFavorite"),
      filterIfTrue(rootState.filter.favoritesOnly)
    );
    const filterByBlackList = flow(
      get("isBlackListed"),
      filterBlackList(rootState.filter.blackListStatus)
    );
    const filterBySweetSpotOnly = flow(
      get("isInSweetSpot"),
      filterIfTrue(rootState.filter.sweetSpotOnly)
    );
    const filterMinPrice = ({ price }) => {
      const minPrice = rootState.filter.minPriceFilter || getters.minPrice;
      return price >= minPrice;
    };
    const filterMaxPrice = ({ price }) => {
      const maxPrice = rootState.filter.maxPriceFilter || getters.maxPrice;
      return price <= maxPrice;
    };

    const filterCondition = overEvery([
      filterByCreationDatetime,
      filterByFavoritesOnly,
      filterByBlackList,
      filterBySweetSpotOnly,
      filterMinPrice,
      filterMaxPrice,
    ]);

    return state.houseList.filter(filterCondition);
  },
  findHouse: (state: State) => (id: number): House => {
    return find({ id }, state.houseList);
  },

  houseListLength: (state, getters): number => {
    return getters.filteredHouseList.length;
  },

  partitionHouseBySweetSpotness: (state: State, getters): House[][] => {
    return partition(
      ({ isInSweetSpot }: House) => isInSweetSpot,
      getters.filteredHouseList
    );
  },

  minPrice: (state): number => {
    return get("price", minBy("price", state.houseList));
  },
  maxPrice: (state): number => {
    return get("price", maxBy("price", state.houseList));
  },
};

const mutations: MutationTree<State> = {
  setHouseList: (state: State, houseList: House[]): void => {
    state.houseList = houseList;
  },
  setIsSyncLoading: (state: State, isSyncLoading: boolean): void => {
    state.isSyncLoading = isSyncLoading;
  },
};

const fetchSelectedHouse = ({
  commit,
}: ActionContext<State, unknown>): Promise<void> => {
  return axios
    .get("/service/houses")
    .then((response) => commit("setHouseList", response.data));
};

const actions: ActionTree<State, State> = {
  fetchSelectedHouse,
  syncHouses: (actionContext: ActionContext<State, unknown>): Promise<void> => {
    actionContext.commit("setIsSyncLoading", true);
    return axios.post("/service/sync-houses").then(
      () => {
        actionContext.commit("setIsSyncLoading", false);
        fetchSelectedHouse(actionContext);
      },
      () => {
        actionContext.commit("setIsSyncLoading", false);
      }
    );
  },
  toggleIsFavorite: (
    actionContext: ActionContext<State, unknown>,
    house: House
  ): Promise<void> => {
    return axios
      .post("/service/house", {
        ...house,
        isFavorite: !house.isFavorite,
      })
      .then(() => {
        fetchSelectedHouse(actionContext);
      });
  },
  toggleIsBlackListed: (
    actionContext: ActionContext<State, unknown>,
    house: House
  ): Promise<void> => {
    return axios
      .post("/service/house", {
        ...house,
        isBlackListed: !house.isBlackListed,
      })
      .then(() => {
        fetchSelectedHouse(actionContext);
      });
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
