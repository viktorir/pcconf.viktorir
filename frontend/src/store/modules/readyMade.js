import api from '@/api/api';

const state = () => ({
  configs: [],
  loading: false,
  error: null,
  currentPage: 1,
  configsPerPage: 12,
  totalFetched: 0, 
});

const getters = {
  paginatedConfigs: (state) => {
    const startIndex = (state.currentPage - 1) * state.configsPerPage;
    const endIndex = startIndex + state.configsPerPage;
    return state.configs.slice(startIndex, endIndex);
  },
  loading: (state) => state.loading,
  error: (state) => state.error,
  currentPage: (state) => state.currentPage,
  totalPages: (state) => {
    return Math.ceil(state.totalFetched / state.configsPerPage);
  },
};

const mutations = {
  ADD_CONFIGS(state, payload) {
    const newConfigs = payload.filter(config => !state.configs.some(existingConfig => existingConfig.id_config === config.id_config));
    state.configs = [...state.configs, ...newConfigs]; 
  },
  SET_LOADING(state, payload) {
    state.loading = payload;
  },
  SET_ERROR(state, payload) {
    state.error = payload;
  },
  SET_CURRENT_PAGE(state, payload) {
    state.currentPage = payload;
  },
  SET_TOTAL_FETCHED(state, payload) {
    state.totalFetched = payload; 
  },
};

const actions = {
  async fetchConfigs({ commit, state }, page = 1) {
    commit('SET_LOADING', true);
    commit('SET_ERROR', null);

    try {
      const startIndex = (page - 1) * state.configsPerPage;

      const res = await api.get(`/configs?page=${page}&limit=${state.configsPerPage}`);
      const body = res.data;

      if (body.configs && body.configs.length > 0) {
        commit('ADD_CONFIGS', body.configs); 
        commit('SET_TOTAL_FETCHED', state.configs.length + body.configs.length);
      } else {
        console.warn('No more configurations available.');
      }
    } catch (e) {
      commit('SET_ERROR', e.message);
      console.error('Error fetching configs:', e);
    } finally {
      commit('SET_LOADING', false);
    }
  },
  setCurrentPage({ commit, dispatch }, page) {
    commit('SET_CURRENT_PAGE', page);
    dispatch('fetchConfigs', page); 
  },
  loadMoreConfigs({ dispatch, state }) {
    const nextPage = state.currentPage + 1;
    dispatch('fetchConfigs', nextPage); 
    commit('SET_CURRENT_PAGE', nextPage); 
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
