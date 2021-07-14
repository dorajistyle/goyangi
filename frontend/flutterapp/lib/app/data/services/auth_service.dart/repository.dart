import 'package:flutterapp/app/data/provider/authentication_provider.dart';

class AuthRepository {
  final AuthenticationProvider authentication_provider;

  AuthRepository(this.authentication_provider);

  login(cred) {
    return authentication_provider.login(cred);
  }
}
