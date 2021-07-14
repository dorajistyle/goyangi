import 'package:get/get.dart';
import 'package:flutterapp/app/core/values/urls.dart';

const baseUrl = API_BASE_URL + '/authentications';

class AuthenticationProvider extends GetConnect {
  Future<Response> login(String cred) => post(baseUrl, cred);
}
