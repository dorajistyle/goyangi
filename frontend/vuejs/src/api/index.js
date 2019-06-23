import axios from 'axios'
import baseURI from '../settings/api_base'
import { Toast } from 'buefy/dist/components/toast'
import i18next from 'i18next'

axios.interceptors.response.use((response) => {
// Do something with response data
  if (response.data) {
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

export default (method, uri, params = null) => {
  if (!method) {
    console.error('API function call requires method argument')
    return
  }

  if (!uri) {
    console.error('API function call requires uri argument')
    return
  }

  var url = baseURI + uri
  return axios({ method, url, params })
}
