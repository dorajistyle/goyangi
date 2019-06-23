import actions from './actions'
import getters from './getters'
import mutations from './mutations'

const state = {
  articleAPI: null,
  pending: false,
  callingAPI: false,
  searching: '',
  articleList: {},
  article: {}
}

export default {
  state,
  actions,
  getters,
  mutations
}
