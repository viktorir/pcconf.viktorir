import { createStore } from 'vuex';
import api from '@/api/api'
import settings from './modules/settings';
import readyMade from './modules/readyMade';

export default createStore({
  state() {
    return {
      configs: [],
      loading: false,
      error: null
    }
  },
  getters: {
    configs: (state) => state.configs,
    loading: (state) => state.loading,
    error: (state) => state.error,
  },
  mutations: {
    SET_CONFIGS(state, payload) {
      state.configs = payload
    },
    SET_LOADING(state, payload) {
      state.loading = payload
    },
    SET_ERROR(state, payload) {
      state.error = payload
    }
  },
  actions: {
    async fetchConfigs({ commit }) {
      commit('SET_LOADING', true)
      commit('SET_ERROR', null)

      try {
        const res = await api.get('/configs?page=1&limit=6');
        let body = res.data
        commit('SET_CONFIGS', body.configs)
      } catch(e) {
        commit('SET_ERROR', e.message)
        console.error('Error fetching configs:', e);
      } finally {
        commit('SET_LOADING', false)
      }
    }
  },
  modules: {
    settings,
    readyMade
  }
})
