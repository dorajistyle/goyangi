<template>
  <div class="container">
    <p class="subtitle is-5">
    {{$t('comment.list.count', {count: commentList.count,  context: (commentList.count !== 0) ? 'many' : 'zero'}) }}
    </p>
    <div class="br"></div>
    <div class="box" v-for="(comment, idx) in commentList.comments" v-bind:key="idx">
      <article class="media">
        <div class="media-left">
          <figure class="image is-64x64">
            <img src="https://bulma.io/images/placeholders/128x128.png" alt="Image">
          </figure>
        </div>
        <div class="media-content">
          <div class="content">
            <span><strong>{{(comment.user||{}).username }}</strong> &nbsp; <small>{{comment.createdAt | moment("from", "now")}}</small></span>
            <p  class="pre-line">{{comment.content}}</p>
          </div>
          <!-- TODO comment reactions -->
          <nav v-if="false" class="level is-mobile"> 
            <div class="level-left">
              <a class="level-item" aria-label="reply">
                <span class="icon is-small">
                  <i class="fas fa-reply" aria-hidden="true"></i>
                </span>
              </a>
              <a class="level-item" aria-label="retweet">
                <span class="icon is-small">
                  <i class="fas fa-retweet" aria-hidden="true"></i>
                </span>
              </a>
              <a class="level-item" aria-label="like">
                <span class="icon is-small">
                  <i class="fas fa-heart" aria-hidden="true"></i>
                </span>
              </a>
            </div>
            <!-- TODO comment reactions END -->
            <div v-if="canWrite(comment.userId)" class="level-right">
                <a @click="editAction(comment)" class="level-item" aria-label="edit">
                  <span class="icon is-small">
                    <i class="fas fa-edit" aria-hidden="true"></i>
                  </span>
                </a>
                <a @click="confirmDeleteComment(comment.id)" class="level-item" aria-label="delete">
                  <span class="icon is-small">
                    <i class="fas fa-trash" aria-hidden="true"></i>
                  </span>
                </a>
            </div>
          </nav>
        </div>
      </article>
    </div>

    <div class="br is-clearfix"></div>
</div>
</template>

<script>
import { mapGetters} from 'vuex'
import i18next from 'i18next'

export default {
  name: 'list',
  computed: mapGetters({
    canWrite: 'canWrite'
  }),
  props: {
      parentId: {
        requred: true
      },
      editAction: {
        required: true
      },
      deleteAction: {
        required: true
      },
      loadMoreAction: {
        required: true
      },
      commentList: {
        required: true,
        default () {
          return {
            hasPrev: false,
            hasNext: false,
            count: 0,
            currentPage: 1,
            comments: []
          }
        }
      }
  },
  data () {
    return {
    }
  },
  created () {
  },
  methods: {
    confirmDeleteComment(commentId) {
        this.$dialog.confirm({
            title: i18next.t('comment.confirmDelete.title'),
            message: i18next.t('comment.confirmDelete.content'),
            cancelText: i18next.t('comment.confirmDelete.cancelText'),
            confirmText: i18next.t('comment.confirmDelete.confirmText'),
            type: 'is-danger',
            hasIcon: true,
            iconPack: 'fas',
            icon: 'minus-circle',
            onConfirm: () => this.deleteAction(commentId)
        })
    },
    scroll () {
      window.onscroll = () => {
        let bottomOfWindow = Math.ceil(Math.max(window.pageYOffset, document.documentElement.scrollTop, document.body.scrollTop)) + window.innerHeight === document.documentElement.offsetHeight
        if (bottomOfWindow) {
          this.loadMoreAction(this.commentList.currentPage)
        }
      }
    }
  }, mounted() {
    this.scroll()
  },
}
</script>

<style scoped>

</style>
