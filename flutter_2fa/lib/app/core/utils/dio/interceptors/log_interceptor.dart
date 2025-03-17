import 'dart:convert';
import 'dart:developer';

import 'package:dio/dio.dart';
import 'package:flutter_2fa/app/core/utils/tool/tips_util.dart';
import 'package:flutter_2fa/app/routes/app_pages.dart';
import 'package:get/route_manager.dart';
import '../../../config/app_config.dart';

/*
 * Log 拦截器
 */
class LogsInterceptors extends InterceptorsWrapper {
  // 请求拦截
  @override
  onRequest(RequestOptions options, handler) async {
    if (AppConfig.DEBUG) {
      print(
          """请求url：${options.baseUrl + options.path}\n请求类型：${options.method}\n请求头：${options.headers.toString()}\nparams参数: ${options.queryParameters}""");
      if (options.data != null) {
        print("""data参数: ${options.data}""");
      }
    }
    // if (AppConfig.contentType == "application/json; charset=UTF-8") {
    //   options.data = jsonEncode(options.data);
    // }
    return handler.next(options);
  }

  // 响应拦截
  @override
  onResponse(response, handler) async {
    if (AppConfig.DEBUG) {
      print('返回参数: $response');
    }
    log("code$response");
    var json = jsonDecode(response.data.toString());
    if (json['code'] == 401) {
      //未登录
      Tips.info("未登录");
      Get.offAndToNamed(Routes.LOGIN);
    }
    return handler.next(response);
  }

  // 请求失败拦截
  @override
  onError(DioException err, handler) async {
    if (AppConfig.DEBUG) {
      print('请求错误: $err');
    }
    return handler.next(err);
  }
}
