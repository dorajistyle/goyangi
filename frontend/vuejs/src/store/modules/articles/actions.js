import {articleTypes} from '../../mutation-types'
import router from '@/router'

export default {
  initArticleAPI ({commit}) {
    commit(articleTypes.SET_API)
  },
  goArticlePath({state}) {
    router.push({path: `/articles/${state.article.id}`} )
  },
  createArticle ({commit, state, dispatch}, formData) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    console.log(`${state.articleAPI}`)
    state
      .articleAPI
      .create(formData)
      .then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE, response)
          commit(articleTypes.DONE)
          dispatch('goArticlePath')
        }

      })
  },
  updateArticle ({commit, state, dispatch}, formData) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    for (var pair of formData.entries()) {
      console.log(`[updateArticle]: ${pair[0]}: ${pair[1]}`)
    }
    state
      .articleAPI
      .update(state.article.id, formData)
      .then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE, response)
          dispatch('refreshArticleComments')
          dispatch('retrieveArticleLikings')
          commit(articleTypes.DONE)
          dispatch('goArticlePath')
        }
      })
  },
  deleteArticle ({commit, state, dispatch}) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    state
      .articleAPI
      .delete(state.article.id)
      .then((response) => {
        if (response) {
          commit(articleTypes.NEW_ARTICLE, response)
          commit(articleTypes.DONE)
        }
        router.push({path: '/articles'})
      })
  },
  newArticle ({commit}) {
    commit(articleTypes.NEW_ARTICLE)
  },
  listArticle ({commit, state, dispatch}, currentPage) {
    dispatch('initArticleAPI')
    console.log(`dispatch list`)
    state
      .articleAPI
      .list(currentPage)
      .then((response) => {
        if (response) {
          console.log(`retrieved articles`)
          commit(articleTypes.SET_ARTICLE_LIST, response)
        }
      })
  },
  retrieveArticle ({commit, state, dispatch}, id) {
    dispatch('initArticleAPI')

    state
      .articleAPI
      .retrieve(id).then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE, response)
        }
      })
  },

  refreshArticleComments ({commit, state, dispatch}) {
    dispatch('initArticleAPI')
    state
      .articleAPI
      .listComments(state.article.id, 1).then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE_COMMENT_LIST, response)
        }
      })
  },
  retrieveMoreArticleComments ({commit, state, dispatch}, currentPage) {
    if(state.article.commentList.hasNext !== true) return

    dispatch('initArticleAPI')
    state
      .articleAPI
      .listComments(state.article.id, ++currentPage).then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE_COMMENT_LIST, response)
        }
      })
  },

  createArticleComment ({commit, state, dispatch}, formData) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    state
      .articleAPI
      .createComment(state.article.id, formData)
      .then((response) => {
        if (response) {
          dispatch('refreshArticleComments')
          commit(articleTypes.DONE)
        }
      })
  },

  updateArticleComment ({commit, state, dispatch}, formData) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    state
      .articleAPI
    .updateComment(state.article.id, formData.get('commentId'), formData)
      .then((response) => {
        if (response) {
          dispatch('refreshArticleComments')
          commit(articleTypes.DONE)
        }
      })
  },

  deleteArticleComment ({commit, state, dispatch}, commentId) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    state
      .articleAPI
      .deleteComment(state.article.id, commentId)
      .then((response) => {
        if (response) {
          dispatch('refreshArticleComments')
          commit(articleTypes.DONE)
        }
      })
  },

  retrieveArticleLikings ({commit, state, dispatch}) {
    dispatch('initArticleAPI')
    state
      .articleAPI
      .listLikings(state.article.id).then((response) => {
        if (response) {
          commit(articleTypes.SET_ARTICLE_LIKING_LIST, response)
        }
      })
  },

  createArticleLiking ({commit, state, dispatch, rootGetters}) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    let formData = new FormData()
    formData.set('parentId',state.article.id)
    formData.set('userId',rootGetters.currentUserId)
    state
      .articleAPI
      .createLiking(state.article.id, formData)
      .then((response) => {
        if (response) {
          dispatch('retrieveArticleLikings')
          commit(articleTypes.DONE)
        }
      })
  },
  deleteArticleLiking ({commit, state, dispatch, rootGetters}) {
    dispatch('initArticleAPI')
    commit(articleTypes.PENDING)
    state
      .articleAPI
      .deleteLiking(state.article.id, rootGetters.currentUserId)
      .then((response) => {
        if (response) {
          dispatch('retrieveArticleLikings')
          commit(articleTypes.DONE)
        }
      })
  }
}
