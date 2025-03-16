import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/api/user_api.dart';
import 'package:flutter_2fa/app/core/service/user_service.dart';
import 'package:flutter_2fa/app/core/utils/tool/common_util.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';

import '../../../routes/app_pages.dart';

class LoginController extends GetxController {
  final checked = false.obs;
  final sendBtnDisable = true.obs;
  final submitBtnDisable = true.obs;
  final TextEditingController emailController = TextEditingController();
  final TextEditingController codeController = TextEditingController();
  final count = 0.obs;
  Timer? _timer;
  @override
  void onReady() {
    super.onReady();
    Future.delayed(const Duration(milliseconds: 500), () {
      emailController.text = "623186518@qq.com";
    });
    emailController.addListener(() {
      if (emailController.text.isNotEmpty) {
        //正则判断是否为邮箱
        if (CommonUtil.isEmail(emailController.text)) {
          sendBtnDisable.value = false;
          changeSubmitBtnDisable();
        } else {
          sendBtnDisable.value = true;
        }
      } else {
        sendBtnDisable.value = true;
      }
    });
    codeController.addListener(() {
      if (codeController.text.isNotEmpty) {
        changeSubmitBtnDisable();
      } else {
        submitBtnDisable.value = true;
      }
    });
  }

  Future<void> sendCode() async {
    EasyLoading.show(status: "发送中...");
    var res = await UserApi.sendCode(
        {"email": CommonUtil.rsaEncrypt(emailController.text)});
    EasyLoading.dismiss();
    if (res['code'] == 200) {
      EasyLoading.showToast("发送成功");
      _startTimer();
    } else {
      EasyLoading.showToast(res['msg']);
    }
  }

  void _startTimer() {
    count.value = 120;
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (count.value == 0) {
        _timer?.cancel();
      } else {
        count.value--;
      }
    });
  }

  void changeSubmitBtnDisable() {
    if (CommonUtil.isEmail(emailController.text) &&
        codeController.text.isNotEmpty &&
        checked.value) {
      submitBtnDisable.value = false;
    } else {
      submitBtnDisable.value = true;
    }
  }

  void login() async {
    EasyLoading.show(status: "登录中...");
    var email = CommonUtil.rsaEncrypt(emailController.text);
    var code = CommonUtil.rsaEncrypt(codeController.text);
    var isLogin = await UserService.to.login(email, code);
    if (isLogin) {
      EasyLoading.showToast("登录成功");
      Get.offAllNamed(Routes.HOME);
    } else {
      EasyLoading.showToast("登录失败");
    }
    EasyLoading.dismiss();
  }
}
