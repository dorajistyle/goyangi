<template>
  <div class="registration">
    <section class="is-fullheight is-dark is-bold">
      <div class="registration-body">
        <div class="container">
          <div class="columns is-vcentered">
            <div class="column is-4 is-offset-4">
              <h1 class="title">
                Register an Account
              </h1>
              <div class="box">
                <form @submit.prevent="createUser({ registrationUsername, registrationEmail, registrationPassword})">
                <label class="label">Username</label>
                <p class="control">
                  <input class="input" type="text" placeholder="jsmith" v-model="registrationUsername">
                </p>
                <label class="label">Email</label>
                <p class="control">
                  <input class="input" type="text" placeholder="jsmith@example.org" v-model="registrationEmail">
                </p>
                <hr>
                <label class="label">Password</label>
                <p class="control">
                  <input class="input" type="password" placeholder="●●●●●●●" v-model="registrationPassword">
                </p>
                <label class="label">Confirm Password</label>
                <p class="control">
                  <input class="input" type="password" placeholder="●●●●●●●">
                </p>
                <hr>
                <p class="control">
                  <button class="button is-primary">Register</button>
                  <button class="button is-default">Cancel</button>
                </p>
                </form>
              </div>
              <p class="has-text-centered">
                <router-link to="/login">
                  Login
                </router-link>
                |
                <a href="#">Need help?</a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'registration',
  data () {
    return {
      registrationUsername: '',
      registrationEmail: '',
      registrationPassword: ''
    }
  },
  methods: {
    createUser () {
      this.$store.dispatch('createUser', {
        registrationUsername: this.registrationUsername,
        registrationEmail: this.registrationEmail,
        registrationPassword: this.registrationPassword
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
