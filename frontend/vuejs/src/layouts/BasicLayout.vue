<template>
  <div id="basic-layout">
    <div class="hero-head">
      <div class="container">
        <nav class="nav has-shadow">
          <div class="container">
            <div class="nav-left">
              <a class="nav-item" href="../index.html">
              </a>
            </div>
            <div class="nav-right nav-menu">
              <router-link to="/" :class="{'nav-item is-tab':true, 'is-active': $route.path === '/'}">
                {{ $t('navbar.home') }}
              </router-link>
              <router-link to="/articles" :class="{'nav-item is-tab':true, 'is-active': $route.path === '/articles'}">
                {{ $t('navbar.examples.articles') }}
              </router-link>
              <span class="nav-item">
                <div v-if="isAuthenticated">
                  <a href="#" @click="logout" class="button">
                    {{ $t('navbar.logout') }}
                  </a>
                </div>
                <div v-else>
                  <router-link to="/login" class="button">
                    {{ $t('navbar.login') }}
                  </router-link>
                  <router-link to="/registration" class="button is-info">
                    {{ $t('navbar.registration') }}
                  </router-link>
                </div>
              </span>
              <a class="nav-item is-tab">
                <b-field>
                    <b-select v-model="locale" @placeholder="$t('navbar.languages.title')" icon="language" icon-pack="fas">
                        <option value="ko">{{ $t('navbar.languages.korean') }}</option>
                        <option value="en">{{ $t('navbar.languages.english') }}</option>
                    </b-select>
                </b-field>
              </a>
            </div>
          </div>
        </nav>
      </div>
    </div>
    <div class="hero-body">
      <div class="container">
        <router-view></router-view>
      </div>
    </div>

    <div id="footer" class="hero-foot">
      <div class="container">
        <div class="tabs is-centered">
          <ul>
            <li>
              <a target="_blank" title="Goyangi Github" href="https://github.com/dorajistyle/goyangi">
                <i class="fab fa-github"></i>
                Goyangi Github
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'basic-layout',
  computed: {
    ...mapGetters([
      'isAuthenticated',
      'currentLocale'
    ]),
    locale : {
      get(){ return this.currentLocale },
      set( locale ){ this.changeLocale( locale )}
    }
  },
  methods: {
    ...mapActions([
      'logout',
      'changeLocale'
    ])
  }
}
</script>

<style>
#basic-layout {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.pre-line {
  white-space: pre-line !important;
}
</style>
