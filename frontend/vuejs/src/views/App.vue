<template>
  <div id="app">
    <router-view></router-view>
    <b-loading :active.sync="pending" :can-cancel="false"></b-loading>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  export default {
    name: 'App',
    computed: mapGetters({
      pending: 'pending'
    }),
    data () {
      return {
        section: 'Head'
      }
    },
    methods: {
      logout () {
        this.$store.commit('SET_USER', null)
        this.$router.push('/login')
      }
    }
  }
</script>