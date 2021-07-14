import 'package:flutter/material.dart';
import 'package:flutterapp/app/modules/home/home_page.dart';
import 'package:get/get.dart';
import 'package:flutterapp/app/global_widgets/top_bar.dart';
import 'package:flutter_login/flutter_login.dart';

const users = const {
  'dribbble@gmail.com': '12345',
  'hunter@gmail.com': 'hunter',
};

class LoginPage extends StatelessWidget {
  Duration get loginTime => Duration(milliseconds: 2250);

  Future<String> _authUser(LoginData data) {
    print('Name: ${data.name}, Password: ${data.password}');
    return Future.delayed(loginTime).then((_) {
      if (!users.containsKey(data.name)) {
        return 'User not exists';
      }
      if (users[data.name] != data.password) {
        return 'Password does not match';
      }
      return data.name;
    });
  }

  Future<String> _recoverPassword(String name) {
    print('Name: $name');
    return Future.delayed(loginTime).then((_) {
      if (!users.containsKey(name)) {
        return 'User not exists';
      }
      return name;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: TopBar(title: 'About GetX'),
        body: FlutterLogin(
          // appBar: TopBar(title: 'Registration'),
          title: 'ECORP',
          logo: 'assets/images/ecorp-lightblue.png',
          onLogin: _authUser,
          onSignup: _authUser,
          onSubmitAnimationCompleted: () {
            Navigator.of(context).pushReplacement(MaterialPageRoute(
              builder: (context) =>
                  HomePage(title: "Home page from Registration Page"),
            ));
          },
          onRecoverPassword: _recoverPassword,
        ));
  }
}
