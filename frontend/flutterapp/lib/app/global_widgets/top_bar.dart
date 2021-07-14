import 'package:flutter/material.dart';
import 'package:flutterapp/app/modules/home/home_page.dart';
import 'package:flutterapp/app/modules/login/login_page.dart';
import 'package:get/get.dart';

// See that I have added: 'implements PreferredSizeWidget'.
class TopBar extends StatelessWidget with PreferredSizeWidget {
  // You also need to override the preferredSize attribute.
  // You can set it to kToolbarHeight to get the default appBar height.

  final String title;
  TopBar({required this.title});

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(BuildContext context) {
    return AppBar(
        leading: IconButton(
          icon: Icon(Icons.home_rounded),
          onPressed: () {
            Get.to(HomePage(title: 'Flutter Home by Icon!'));
          },
        ),
        // backgroundColor: Theme.of(context).scaffoldBackgroundColor,
        title: Text(title),
        actions: <Widget>[
          TextButton(
              style: TextButton.styleFrom(
                  padding:
                      EdgeInsets.only(top: 20, bottom: 20, right: 20, left: 20),
                  minimumSize: Size(50, 30),
                  alignment: Alignment.centerLeft),
              onPressed: () {},
              child: Row(
                // Replace with a Row for horizontal icon + text
                children: <Widget>[
                  Icon(Icons.add),
                  Text('Home',
                      style: TextStyle(fontSize: 17.0, color: Colors.black87))
                ],
              )),
          TextButton(
              onPressed: () {},
              child: Text('Articles',
                  style: TextStyle(fontSize: 17.0, color: Colors.black87))),
          ElevatedButton(
              onPressed: () {},
              child: Text('Login',
                  style: TextStyle(fontSize: 17.0, color: Colors.black87))),
          OutlinedButton(
              style: OutlinedButton.styleFrom(
                tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                padding:
                    EdgeInsets.only(top: 20, bottom: 20, right: 20, left: 20),
                shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(12.0)),
                side: BorderSide(
                    width: 0.5, color: Theme.of(context).primaryColor),
                primary: Colors.white.withOpacity(0.9),
              ),
              onPressed: () {
                Get.to(LoginPage());
              },
              child: Text('Registration',
                  style: TextStyle(fontSize: 17.0, color: Colors.black87))),
          IconButton(
            icon: const Icon(Icons.add_alert),
            tooltip: 'Show Snackbar',
            onPressed: () {
              ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('This is a snackbar')));
            },
          ),
        ]);
  }
}
