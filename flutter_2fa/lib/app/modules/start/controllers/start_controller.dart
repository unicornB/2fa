import 'package:flutter_2fa/app/core/service/user_service.dart';
import 'package:flutter_2fa/app/routes/app_pages.dart';
import 'package:get/get.dart';

class StartController extends GetxController {
  //TODO: Implement StartController

  final count = 0.obs;

  void increment() => count.value++;

  void toHome() {
    //延时跳转到首页
    Future.delayed(const Duration(seconds: 2), () {
      Get.offNamed(Routes.HOME);
      UserService.to.getUserInfo();
    });
  }
}
