export default {
  articleList: state => state.articleList,
  canWrite: (_, rootGetters) => userId => userId == rootGetters.currentUserId,
  article: state => state.article,
}
