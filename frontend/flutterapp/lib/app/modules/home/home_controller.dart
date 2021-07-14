import 'package:get/get.dart';

class HomePageController extends GetxController {
  final count = 0.obs;
  increment() => count.value++;
}
