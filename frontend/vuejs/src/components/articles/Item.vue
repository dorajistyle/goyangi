<template>
  <div class="container">
    <nav class="breadcrumb" aria-label="breadcrumbs">
      <ul>
        <li>
          <router-link to="/articles" class="has-text-centered">
            <span>
              {{ $t('article.view.item.list') }}
            </span>
          </router-link>
        </li>
        <li class="is-active"><a href="#" aria-current="page">{{article.title}}</a></li>
      </ul>
    </nav>
    <div class="br is-clearfix"></div>
    <div v-if="canWrite(article.userId)" class="text-center">
      <a @click="confirmDelete" class="button is-pulled-right">
        <span class="icon">
          <i class="fas fa-trash"></i>
        </span>
        <span>
          {{ $t('article.view.item.delete') }}
        </span>
      </a>
      <router-link :to="`/articles/${article.id}/edit`" class="button is-pulled-right">
        <span class="icon">
          <i class="fas fa-edit"></i>
        </span>
        <span>
          {{ $t('article.view.item.edit') }}
        </span>
      </router-link>
    </div>
    <div class="br is-clearfix"></div>
    <div class="content">
      <h1 class="title is-2">
          {{article.title}}
        </h1>
      <div>
        <div class="is-pulled-left" v-if="article.author != null">
          {{article.author.username}}
        </div>
        <div class="is-pulled-right">
          {{article.createdAt | moment("MM-DD-YYYY, h:mm:ss a")}}
        </div>
      </div>
      <div class="br is-clearfix"></div>
      <div class="box">
        <p class="pre-line">
          {{article.content}}
        </p>
      </div>
    </div>
    <likings
    :likingList="article.likingList"
    :createAction="createArticleLiking"
    :deleteAction="deleteArticleLiking"
    componentClass="is-pulled-left"
    />
    <div class="br is-clearfix"></div>
    <div class="is-divider"></div>
    <comments
    :commentList="article.commentList"
    :parentId="article.id"
    :createAction="createArticleComment"
    :updateAction="updateArticleComment"
    :deleteAction="deleteArticleComment"
    :loadMoreAction="retrieveMoreArticleComments"
    />
    <router-view></router-view>
</div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import i18next from 'i18next'
import comments from '@/components/comments'
import likings from '@/components/likings'

export default {
  name: 'item',
  components: {
    comments,
    likings
  },
  computed: mapGetters({
    article: 'article',
    canWrite: 'canWrite'
  }),
  props: {
      articleId: {
        type: String,
        required: true
      }
  },
  created () {
    console.log(`Item this.article: ${JSON.stringify(this.article)}`)
    if(this.article == undefined || Object.keys(this.article).length === 0  || this.article.id != this.articleId) {
      this.retrieveArticle(this.articleId)
    }
  },
  data () {
    return {
    }
  },
  methods: {
    ...mapActions([
      'retrieveArticle',
      'deleteArticle',
      'createArticleComment',
      'updateArticleComment',
      'deleteArticleComment',
      'retrieveMoreArticleComments',
      'createArticleLiking',
      'deleteArticleLiking'
    ]),
    confirmDelete() {
                this.$dialog.confirm({
                    title: i18next.t('article.view.item.confirm.title'),
                    message: i18next.t('article.view.item.confirm.content'),
                    cancelText: i18next.t('article.view.item.confirm.cancelText'),
                    confirmText: i18next.t('article.view.item.confirm.confirmText'),
                    type: 'is-danger',
                    hasIcon: true,
                    iconPack: 'fas',
                    icon: 'minus-circle',
                    onConfirm: () => this.deleteArticle()
                })
            }
  },
  mounted () {
      // if(this.article.id == null) this.$router.push({path: '/articles'})
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
