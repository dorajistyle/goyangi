import UserAPI from '../../../api/UserAPI'

import { userTypes } from '../../mutation-types'

export default {
  [userTypes.SET_API] (state) {
    if (state.userAPI === null) {
      state.userAPI = new UserAPI()
    }
  },
  [userTypes.CREATE_USER] (state) {
    state.pending = true
  },
  [userTypes.CREATE_USER_SUCCESS] (state) {
    state.createdUser = true
    state.pending = false
  }
}
