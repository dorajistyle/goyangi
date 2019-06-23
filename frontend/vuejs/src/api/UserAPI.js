import BaseAPI from './BaseAPI'

class UserAPI extends BaseAPI {
  constructor () {
    super('/users')
  }

  create (cred) {
    return this.post('', cred)
  }

}

export default UserAPI
