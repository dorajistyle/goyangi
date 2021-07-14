import 'package:get/get.dart';
import 'repository.dart';

class AuthService extends GetxService {
  final AuthRepository _authRepository = Get.find<AuthRepository>();
  Future<AuthService> init() async {
    return this;
  }

  Future<AuthService> login(cred) async {
    return _authRepository.login(cred);
  }
}
