// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import { sync } from 'vuex-router-sync'
import App from './views/App'
import router from './router'
import store from './store'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import 'bulma-divider'
import Moment from 'vue-moment'
import i18next from 'i18next'
import i18nOption from './settings/i18next'
import VueI18Next from '@panter/vue-i18next'
import VeeValidate, { Validator } from 'vee-validate'
import ko from 'vee-validate/dist/locale/ko'
import $ from 'jquery'

Vue.use({
  install: function(Vue){
      Vue.prototype.$jQuery = $; // you'll have this.$jQuery anywhere in your vue project
  }
  })
Vue.use(Buefy)
Vue.use(Moment)
Vue.use(VueI18Next)
sync(store, router)

var lang = $('html').attr('lang')

lang !== undefined
    ? i18nOption.lng = lang
    : i18nOption.lng = 'en'
i18next.init(i18nOption)
const i18n = new VueI18Next(i18next)

Vue.config.productionTip = false

Vue.use(VeeValidate, {
  locale: 'ko'
})
Validator.localize('ko', ko)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router: router,
  store: store,
  i18n: i18n,
  render: h => h(App)
})

// Check local storage to handle refreshes
if (window.localStorage) {
  var localUserString = window.localStorage.getItem('user') || 'null'
  var localUser = JSON.parse(localUserString)

  if (localUser && store.state.user !== localUser) {
    store.commit('SET_USER', localUser)
    store.commit('SET_IS_AUTHENTICATED', window.localStorage.getItem('isAuthenticated'))
  }
}
