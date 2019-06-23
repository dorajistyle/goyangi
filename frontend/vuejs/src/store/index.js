import Vue from 'vue'
import Vuex from 'vuex'
import authentications from './modules/authentications'
import users from './modules/users'
import articles from './modules/articles'
import languages from './modules/languages'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    authentications: authentications,
    users: users,
    articles: articles,
    languages: languages
  }
})

export default store
