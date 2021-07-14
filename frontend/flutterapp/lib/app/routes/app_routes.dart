import 'package:get/get.dart';
import 'package:flutterapp/app/modules/home/home_page.dart';
import 'package:flutterapp/app/modules/about/about_page.dart';
import 'package:flutterapp/app/modules/login/login_page.dart';

class AppRoutes {
  static final routes = [
    GetPage(name: '/', page: () => HomePage(title: 'Flutter Demo Home Page')),
    GetPage(name: '/about', page: () => AboutPage()),
    GetPage(name: '/login', page: () => LoginPage()),
  ];
}
