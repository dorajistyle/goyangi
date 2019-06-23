import BaseAPI from './BaseAPI'

class ArticleAPI extends BaseAPI {
  constructor () {
    super('/articles')
  }

  create (formData) {
    // for (var pair of formData.entries()) {
    //   console.log(`API ${pair[0]}: ${pair[1]}`)
    // }
    return super.post('', null, formData)
  }

  update (id, formData) {
    console.log(`ID : ${formData.get('id')}`)
    for (var pair of formData.entries()) {
      console.log(`API ${pair[0]}: ${pair[1]}`)
    }
    return super.put(id, null, formData)
  }

  delete (id) {
    return super.delete(id, null)
  }

  list (currentPage = 1) {
    return super.get('', {currentPage: currentPage})
  }

  retrieve (id) {
    return super.get(id)
  }

  listComments(id, currentPage = 1) {
    return super.get(`${id}/comments`, {currentPage: currentPage})
  }

  createComment (id, formData) {
    return super.post(`${id}/comments`, null, formData)
  }

  updateComment (id, commentId, formData) {
    return super.put(`${id}/comments/${commentId}`, null, formData)
  }

  deleteComment (id, commentId) {
    return super.delete(`${id}/comments/${commentId}`, null)
  }

  listLikings(id, currentPage = 1) {
    return super.get(`${id}/likings`, {currentPage: currentPage})
  }

  createLiking (id, formData) {
    return super.post(`${id}/likings`, null, formData)
  }

  deleteLiking (id, userId) {
    return super.delete(`${id}/likings/${userId}`, null)
  }



}

export default ArticleAPI
