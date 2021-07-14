import 'package:get/get.dart';
import 'package:flutterapp/app/core/values/urls.dart';
import 'package:flutterapp/app/data/model/article_model.dart';
import 'package:flutterapp/app/data/model/comment_model.dart';

const baseUrl = API_BASE_URL + '/articles';

class ArticleProvider extends GetConnect {
  Future<Response<List<ArticleModel>>> getArticles(int currentPage) {
    final uri = Uri.https(baseUrl, '', {'currentPage': currentPage});
    return get(uri.toString());
  }

  Future<Response> getArticle(int id) => get('$baseUrl/$id');

  Future<Response<ArticleModel>> createArticle(Map formData) {
    return post(baseUrl, formData);
  }

  Future<Response<ArticleModel>> updateArticle(int id, Map formData) {
    return put("$baseUrl/$id", formData);
  }

  Future<Response<ArticleModel>> deleteArticle(int id) {
    return delete("$baseUrl/$id");
  }

  Future<Response<List<CommentModel>>> getComments(int id, int currentPage) {
    final uri =
        Uri.https(baseUrl, '/$id/comments', {'currentPage': currentPage});
    return get(uri.toString());
  }

  Future<Response> getComment(int id, int commentId) =>
      get('$baseUrl/$id/comments/$commentId');

  Future<Response<CommentModel>> createComment(int id, Map formData) {
    return post('$baseUrl/$id/comments', formData);
  }

  Future<Response<CommentModel>> updateComment(
      int id, int commentId, Map formData) {
    return put("$baseUrl/$id/comments/$commentId", formData);
  }

  Future<Response<CommentModel>> deleteComment(int id, int commentId) {
    return delete("$baseUrl/$id/comments/$commentId");
  }

  Future<Response<List<CommentModel>>> getLikings(int id, int currentPage) {
    final uri =
        Uri.https(baseUrl, '/$id/likings', {'currentPage': currentPage});
    return get(uri.toString());
  }

  Future<Response<CommentModel>> createLiking(int id, Map formData) {
    return post('$baseUrl/$id/likings', formData);
  }

  Future<Response<CommentModel>> deleteLiking(int id, int userId) {
    return delete("$baseUrl/$id/comments/$userId");
  }
}
