import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:flutterapp/app/routes/app_routes.dart';
import 'package:flutterapp/app/core/theme/app_theme.dart';

import 'modules/home/home_page.dart';

class AppWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      initialRoute: '/',
      defaultTransition: Transition.fade,
      debugShowCheckedModeBanner: false,
      getPages: AppRoutes.routes,
      theme: appThemeData,
      home: HomePage(title: 'Flutter Demo Home Page'),
    );
  }
}
