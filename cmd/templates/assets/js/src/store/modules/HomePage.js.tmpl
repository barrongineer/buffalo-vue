import http from 'axios'
import { FETCH_ROUTES } from '../action-types'
import { SET_ROUTES } from '../mutation-types'

const state = {
  buffaloRoutes: []
}

const mutations = {
  [SET_ROUTES] (state, routes) {
    state.buffaloRoutes = routes
  }
}

const actions = {
  [FETCH_ROUTES] ({commit}) {
    return new Promise((resolve, reject) => {
      http.get('/api/app/routes')
        .then(res => {
          commit(SET_ROUTES, res.data)
          resolve(res.data)
        })
        .catch(err => reject(err))
    })
  }
}

export default {
  state,
  mutations,
  actions
}