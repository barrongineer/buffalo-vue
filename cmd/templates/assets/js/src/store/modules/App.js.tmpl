import http from 'axios'
import { FETCH_I18N } from '../action-types'
import { SET_GREETING } from '../mutation-types'

const state = {
  greeting: ''
}

const mutations = {
  [SET_GREETING] (state, greeting) {
    state.greeting = greeting
  }
}

const actions = {
  [FETCH_I18N] ({commit}, key) {
    return new Promise((resolve, reject) => {
      http.get(`/api/i18n/${key}`)
        .then(res => {
          commit(SET_GREETING, res.data)
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