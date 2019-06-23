import AuthenticationAPI from '../../../api/AuthenticationAPI'

import { authenticationTypes } from '../../mutation-types'

export default {
  [authenticationTypes.SET_API] (state) {
    if (state.authenticationAPI === null) {
      state.authenticationAPI = new AuthenticationAPI()
    }
  },
  [authenticationTypes.LOGIN] (state) {
    state.pending = true
  },
  [authenticationTypes.LOGIN_SUCCESS] (state) {
    state.isAuthenticated = true
    state.pending = false
  },
  [authenticationTypes.LOGOUT] (state) {
    state.isAuthenticated = false
  }
}
