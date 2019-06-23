<template>
  <div class="container list">
    <!--<h2>{{i18n 'article.view.list.title'}}</h2>-->
    <div v-if="isAuthenticated" class="text-center">
      <router-link to="/articles/new" class="button is-pulled-right">
        <span class="icon">
          <i class="fas fa-pen-alt"></i>
        </span>
        <span>
          {{ $t('article.view.list.write') }}
        </span>
      </router-link>
    </div>
    <div class="br"></div>
    <div class="">
      <table class="table is-striped is-fullwidth">
          <tbody v-for="(article, idx) in articleList.articles" v-bind:key="idx">
                <tr>
                  <td>{{article.id}}</td>
                  <td><router-link :to="`/articles/${article.id}`">{{article.title}}</router-link></td>
                  <td>{{article.author.username}}</td>
                  <td>{{article.createdAt | moment("MM-DD-YYYY, h:mm:ss a")}}</td>
                </tr>
          </tbody>
      </table>

      <b-pagination
          :total="articleList.count"
          :current.sync="current"
          :order="order"
          :size="size"
          :simple="isSimple"
          :rounded="isRounded"
          :per-page="articleList.perPage"
          icon-pack="fas"
          aria-next-label="Next page"
          aria-previous-label="Previous page"
          aria-page-label="Page"
          aria-current-label="Current page"
          @change="pageChange"
          >
      </b-pagination>

    </div>
    <router-view></router-view>
</div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'list',
  computed: mapGetters({
    articleList: 'articleList',
    isAuthenticated: 'isAuthenticated'
  }),
  props: {
      currentPage: {
        type: String,
        required: false
      }
  },
  data () {
    return {
      current: this.currentPage || 1,
      order: '',
      size: '',
      isSimple: false,
      isRounded: false
    }
  },
  created () {
    this.listArticle(this.current)
  },
  methods: {
    ...mapActions([
      'listArticle'
    ]),
    pageChange(page) {
      this.listArticle(page) // api call
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
