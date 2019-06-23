import Vue from 'vue'
import Router from 'vue-router'
import BasicLayout from '@/layouts/BasicLayout'
import Home from '@/views/Home'
import usersRoute from './users'
import articlesRoute from './articles'
import store from '@/store'
Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: BasicLayout,
      children: [
        {
          path: '',
          component: Home
        }, ...usersRoute,
        ...articlesRoute
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth) && !store.state.authentications.isAuthenticated) return next("/login")
   next()
})

export default router


