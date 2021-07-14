class CommentModel {
  int id = -1;
  String name = '';

  CommentModel({required this.id, required this.name});

  CommentModel.fromJson(Map<String, dynamic> json) {
    this.id = json['id'];
    this.name = json['name'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['name'] = this.name;
    return data;
  }
}
