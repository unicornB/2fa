import 'dart:convert';
import 'dart:developer';

import 'package:dio/dio.dart';
import 'package:flutter_2fa/app/core/config/app_config.dart';
import 'package:flutter_2fa/app/core/utils/tool/common_util.dart';

class SignInterceptors extends InterceptorsWrapper {
  // 请求拦截
  @override
  onRequest(RequestOptions options, handler) async {
    //参数签名
    _sign(options);
    return handler.next(options);
  }

  _sign(RequestOptions options) {
    //获取query参数和body参数，进行合并
    var params = options.queryParameters;
    var body = options.data;
    if (options.data is String) {
      body = jsonDecode(options.data);
    }
    if (body != null) {
      params.addAll(body);
    }
    //添加时间戳
    params['timestamp'] = DateTime.now().millisecondsSinceEpoch ~/ 1000;
    options.headers['Timestamp'] = params['timestamp'];
    //添加随机字符串
    params['nonce'] = CommonUtil.generateRandomString(32);
    options.headers['Nonce'] = params['nonce'];
    //对参数进行排序
    var keys = params.keys.toList();
    keys.sort();
    //拼接Json字符串
    var json = StringBuffer();
    json.write('{');
    for (var key in keys) {
      json.write('"$key":"${params[key]}",');
    }
    json.write('}');
    //移除最后一个逗号
    var jsonStr = json.toString();
    jsonStr = jsonStr.substring(0, jsonStr.length - 2);
    jsonStr += "}${AppConfig.signKey}";
    log(jsonStr);
    //对字符串进行sha256加密
    var sign = CommonUtil.sha256Hash(jsonStr);
    //添加签名
    options.headers['Sign'] = sign;
  }
}
