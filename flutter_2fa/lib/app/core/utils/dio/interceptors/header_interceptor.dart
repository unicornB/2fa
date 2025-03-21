import 'dart:convert';
import 'dart:io';

import 'package:dio/dio.dart';
import 'package:flutter_2fa/app/core/utils/tool/preference_utils.dart';

import '../../../config/app_config.dart';

/*
 * header拦截器
 */
class HeaderInterceptors extends InterceptorsWrapper {
  // 请求拦截
  @override
  onRequest(RequestOptions options, handler) async {
    options.baseUrl = AppConfig.host;
    options.contentType = AppConfig.contentType;
    options.headers = {
      HttpHeaders.contentTypeHeader: AppConfig.contentType,
    };
    var token = PreferenceUtils.instance.getString("token", "");
    if (token.isNotEmpty) {
      options.headers.addAll({
        "Authorization": token,
      });
    }

    return handler.next(options);
  }

  // 响应拦截
  @override
  onResponse(response, handler) {
    // Do something with response data
    return handler.next(response); // continue
  }

  // 请求失败拦截
  @override
  onError(err, handler) async {
    return handler.next(err); //continue
  }
}
