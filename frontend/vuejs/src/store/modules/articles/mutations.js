import ArticleAPI from '../../../api/ArticleAPI'

import { articleTypes } from '../../mutation-types'

export default {
  [articleTypes.SET_API] (state) {
    if (state.articleAPI === null) {
      state.articleAPI = new ArticleAPI()
    }
  },
  [articleTypes.PENDING] (state) {
    state.pending = true
  },
  [articleTypes.DONE] (state) {
    state.pending = false
  },
  [articleTypes.NEW_ARTICLE] (state) {
    state.article = {}
  },
  [articleTypes.SET_ARTICLE_LIST] (state, response) {
    state.articleList = response.articleList
  },
  [articleTypes.SET_ARTICLE] (state, response) {
    console.log(`SET_ARTICLE response : ${JSON.stringify(response)}`)
    state.article = response.article
    console.log(`state.article.id : ${state.article.id}`)
  },
  [articleTypes.SET_ARTICLE_COMMENT_LIST] (state, response) {
    let comments = state.article.commentList.comments
    state.article.commentList = response.commentList
    if(state.article.commentList.currentPage !== 1 && state.article.commentList.comments != null) {
        state.article.commentList.comments = comments.concat(state.article.commentList.comments)
    }
  },
  [articleTypes.SET_ARTICLE_LIKING_LIST] (state, response) {
        state.article.likingList = response.likingList
  }

}
