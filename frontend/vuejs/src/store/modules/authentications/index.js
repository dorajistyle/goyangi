import actions from './actions'
import getters from './getters'
import mutations from './mutations'

const state = {
  isAuthenticated: !!localStorage.getItem('isAuthenticated'),
  authenticationAPI: null,
  pending: false,
  callingAPI: false,
  searching: '',
  currentUser: {id: 1}
}

export default {
  state,
  actions,
  getters,
  mutations
}
