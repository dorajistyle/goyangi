import 'package:get/get.dart';
import 'package:flutterapp/app/core/values/urls.dart';

const baseUrl = API_BASE_URL + '/users';

class UserProvider extends GetConnect {
  Future<Response> create(Map data) => post(baseUrl, data);
}
