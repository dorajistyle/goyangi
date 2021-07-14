class LikingModel {
  int id = -1;
  String name = '';

  LikingModel({required this.id, required this.name});

  LikingModel.fromJson(Map<String, dynamic> json) {
    this.id = json['id'];
    this.name = json['name'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['name'] = this.name;
    return data;
  }
}
