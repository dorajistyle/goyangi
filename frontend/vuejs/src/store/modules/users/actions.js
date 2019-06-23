import {userTypes} from '../../mutation-types'

export default {
  initUserAPI ({commit}) {
    commit(userTypes.SET_API)
  },
  createUser ({commit, state, dispatch}, creds) {
    dispatch('initUserAPI')
    commit(userTypes.CREATE_USER) // show spinner

    state
      .userAPI
      .create(creds)
      .then((response) => {
        if (response) {
          localStorage.setItem('token', 'JWT')
          commit(userTypes.CREATE_USER_SUCCESS)
        }
      })
  }
}
