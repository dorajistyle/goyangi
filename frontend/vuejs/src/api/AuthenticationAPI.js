import BaseAPI from './BaseAPI'

class AuthenticationAPI extends BaseAPI {
  constructor () {
    super('/authentications')
  }

  login (cred) {
    return this.post('', cred)
  }

}

export default AuthenticationAPI
