import {authenticationTypes} from '../../mutation-types'

export default {
  initAuthenticationAPI ({commit}) {
    commit(authenticationTypes.SET_API)
  },
  login ({commit, state, dispatch}, creds) {
    dispatch('initAuthenticationAPI')
    commit(authenticationTypes.LOGIN) // show spinner

    state
      .authenticationAPI
      .login(creds)
      .then((response) => {
        if (response) {
          localStorage.setItem('isAuthenticated', true)
          commit(authenticationTypes.LOGIN_SUCCESS)
        }
      })
    // return new Promise(resolve => {
    //   setTimeout(() => {
    //     localStorage.setItem('token', 'JWT')
    //     commit(types.LOGIN_SUCCESS)
    //     resolve()
    //   }, 1000)
    // })
  },
  logout ({ commit }) {
    localStorage.removeItem('isAuthenticated')
    commit(authenticationTypes.LOGOUT)
  }
}
