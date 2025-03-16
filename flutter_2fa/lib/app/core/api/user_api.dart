import '../utils/dio/request.dart';

class UserApi {
  static Future<dynamic> getMe() async {
    return Request.get('/api/user/getMe', queryParameters: {});
  }

  static Future<dynamic> sendCode(Object data) async {
    return Request.post('/api/user/send_email_code', data: data);
  }

  //登录
  static Future<dynamic> login(Object data) async {
    return Request.post('/api/user/login', data: data);
  }
}
