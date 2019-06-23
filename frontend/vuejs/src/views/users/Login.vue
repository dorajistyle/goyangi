<template>
  <div class="login">
    <section class="is-fullheight is-dark is-bold">
      <div class="container">
        <div class="columns is-vcentered">
          <div class="column is-4 is-offset-4">
            <h1 class="title">
              Login
            </h1>
            <form @submit.prevent="login({ loginEmail, loginPassword })">
              <div class="box">
                <label class="label">Email</label>
                <p class="control">
                  <input class="input" type="text" placeholder="jsmith@example.org" v-model="loginEmail">
                </p>
                <hr>
                <label class="label">Password</label>
                <p class="control">
                  <input class="input" type="password" placeholder="●●●●●●●" v-model="loginPassword">
                </p>
                <hr>
                <p class="control">
                  <button class="button is-primary">Login</button>
                  <button class="button is-default">Cancel</button>
                </p>
              </div>
            </form>
            <p class="has-text-centered">
              <router-link to="/registration">
                Registration
              </router-link>
              |
              <a href="#">Need help?</a>
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {
      loginEmail: '',
      loginPassword: ''
    }
  },
  methods: {
    login () {
      this.$store.dispatch('login', {
        loginEmail: this.loginEmail,
        loginPassword: this.loginPassword
      }).then(() => {
        this.$router.push('/')
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
