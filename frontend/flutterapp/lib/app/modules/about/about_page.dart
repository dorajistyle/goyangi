import 'package:flutter/material.dart';
import 'package:flutterapp/app/modules/home/home_page.dart';
import 'package:get/get.dart';
import 'package:flutterapp/app/modules/home/home_controller.dart';
import 'package:flutterapp/app/global_widgets/top_bar.dart';

class AboutPage extends StatelessWidget {
  // final HomePageController controller = Get.put(HomePageController());
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: TopBar(title: 'About GetX'),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: Text(
                'GetX is an extra-light and powerful solution for Flutter. It combines high performance state management, intelligent dependency injection, and route management in a quick and practical way.',
              ),
            ),
            TextButton(
                onPressed: () {
                  Get.to(HomePage(title: 'Flutter Home Again!'));
                },
                child: Text('Go Home'))
          ],
        ),
      ),
    );
  }
}
