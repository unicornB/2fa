import 'dart:convert';
import 'dart:developer';

import 'package:flutter_2fa/app/core/api/user_api.dart';
import 'package:get/get.dart';

import '../utils/tool/preference_utils.dart';

class UserService extends GetxService {
  static UserService to = Get.find<UserService>();
  Future<UserService> init() async {
    return this;
  }

  final isLogined = false.obs; //是否登录
  final loginToken = ''.obs; //token

  Future<void> getUserInfo() async {
    var res = await UserApi.getMe();
    log(jsonEncode(res));
  }

  Future<bool> login(String email, String code) async {
    var res = await UserApi.login({"email": email, "code": code});
    log(jsonEncode(res));
    if (res['code'] == 200) {
      isLogined.value = true;
      loginToken.value = res['data']['token'];
      PreferenceUtils.instance.putString("token", res['data']['token']);
      return true;
    }
    return false;
  }
}
