class ArticleModel {
  int id = -1;
  String title = "";
  String body = "";

  ArticleModel({
    required this.id,
    required this.title,
    required this.body,
  });

  ArticleModel.fromJson(Map<String, dynamic> json) {
    this.id = json['id'];
    this.title = json['title'];
    this.body = json['body'];
  }

  Map<String, dynamic> toJson() {
    return {
      'title': this.title,
      'body': this.body,
    };
  }
}
