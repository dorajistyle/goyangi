// import types from './action_types'
// import api from '../api'
// export default {
//   login ({ commit }, creds) {
//     commit(types.LOGIN) // show spinner
//     return api('post', '/authentications', creds)
//       .then((response) => {
//         if (response) {
//           // console.log('response', response)

//           localStorage.setItem('token', 'JWT')
//           commit(types.LOGIN_SUCCESS)
//         }
//       })
//     // return new Promise(resolve => {
//     //   setTimeout(() => {
//     //     localStorage.setItem('token', 'JWT')
//     //     commit(types.LOGIN_SUCCESS)
//     //     resolve()
//     //   }, 1000)
//     // })
//   },
//   logout ({ commit }) {
//     localStorage.removeItem('token')
//     commit(types.LOGOUT)
//   }
// }
