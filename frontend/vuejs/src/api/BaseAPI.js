import axios from 'axios'
import { Toast } from 'buefy/dist/components/toast'
import i18next from 'i18next'
import humps from 'humps'
import baseURI from '@/settings/api_base'
import store from '@/store'
import router from '@/router'



// const CSRF_TOKEN = {'X-CSRF-Token': document.getElementsByName('csrf-token')[0].getAttribute('content')}
// axios.interceptors.request.use(
//   (config) => {
//     config.withCredentials = true
//     return config
//   },
//   (error) => Promise.reject(error)
// )

// The below line commented to prevent an error: gin has been blocked by CORS policy: The value of the 'Access-Control-Allow-Origin' header in the response must not be the wildcard '*' when the request's credentials mode is 'include'. The credentials mode of requests initiated by the XMLHttpRequest is controlled by the withCredentials attribute.
// axios.defaults.headers['Access-Control-Allow-Headers'] = "Access-Control-Allow-Headers, , access-control-allow-origin, Access-Control-Allow-Origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"
axios.defaults.headers['Access-Control-Allow-Origin'] = ['http://localhost:3001', 'http://localhost']
axios.defaults.withCredentials = true

axios.interceptors.response.use((response) => {
// Do something with response data
  if (response.data && response.data.messageType) {
    Toast.open({
      duration: 5000,
      message: i18next.t(response.data.messageType),
      type: 'is-success'})
  }
  return response
}, (error) => {
  // Do something with response error
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    if (error.response.data) {
      Toast.open({
        duration: 2000,
        message: i18next.t(error.response.data.messageType),
        type: 'is-danger'})
    } else {
      if (error.response.status) {
        console.log('status : ', error.response.status)
      }
      if (error.response.headers) {
        console.log('header : ' + JSON.stringify(error.response.headers))
      }
    }
    if (error.response.status == 401) {
      store.dispatch('logout')
      router.push('/login')
    }

  } else if (error.request) {
    // The request was made but no response was received
    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
    // http.ClientRequest in node.js
    console.log('request : ' + error.request)
  } else {
    // Something happened in setting up the request that triggered an Error
    console.log('Error', error.message)
    console.log('config error : ' + JSON.stringify(error.config))
  }
  // return Promise.reject(error)
})

class BaseAPI {
  constructor (endpoint) {
    console.log(`endpoint ${endpoint}`)
    this.endpoint = baseURI + endpoint
  }

  get (path, params = null, data = null) {
    return this.send('get', path, params, data)
  }

  post (path, params = null, data = null) {
    return this.send('post', path, params, data)
  }

  put (path, params = null, data = null) {
    return this.send('put', path, params, data)
  }

  delete (path, params = null, data = null) {
    return this.send('delete', path, params, data)
  }

  send (method, path, params = null, data = null) {
    var url
    console.log(`send ${method}, ${path}, ${params}, ${data}`)
    if (path) {
      url = `${this.endpoint}/${path}`
    } else {
      url = `${this.endpoint}`
    }

    return axios({
      method: method,
      url: url,
      params: params,
      data: data
      // headers: Object.assign({'X-Requested-With': 'XMLHttpRequest'}, CSRF_TOKEN)
    }).then(response => {
      // console.log(`response: ${response}`)
      var data = {}
      if(response) {
        data = humps.camelizeKeys(response.data)
      }
      return Promise.resolve(data)
      // if (data.meta.code == 200) {
      //   return Promise.resolve(data)
      // } else {

      // }
    })
  }
}

export default BaseAPI
