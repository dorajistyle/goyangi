import actions from './actions'
import getters from './getters'
import mutations from './mutations'

const state = {
  createdUser: !!localStorage.getItem('token'),
  userAPI: null,
  pending: false,
  callingAPI: false,
  searching: '',
  user: null
}

export default {
  state,
  actions,
  getters,
  mutations
}
