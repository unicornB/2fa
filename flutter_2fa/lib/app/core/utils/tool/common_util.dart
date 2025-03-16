import 'dart:math';
import 'package:crypto/crypto.dart';
import 'dart:convert';

import 'package:encrypt/encrypt.dart' as encrypt;
import 'package:flutter_2fa/app/core/config/app_config.dart';
import 'package:pointycastle/export.dart';

class CommonUtil {
  //生成随机字符串
  static String generateRandomString(int len) {
    var r = Random();
    const chars =
        'AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890';
    return List.generate(len, (index) => chars[r.nextInt(chars.length)]).join();
  }

  //sha256加密
  static String sha256Hash(String data) {
    var bytes = utf8.encode(data);
    var digest = sha256.convert(bytes);
    return digest.toString();
  }

  //判断是否是邮箱
  static bool isEmail(String email) {
    final emailRegex = RegExp(r'^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$');
    return emailRegex.hasMatch(email);
  }

  //Rsa加密
  static String rsaEncrypt(String plainText) {
    final publicKey = encrypt.RSAKeyParser().parse(AppConfig.rsaPublicKey);
    RSAPublicKey? key = publicKey as RSAPublicKey?;
    final encrypter = encrypt.Encrypter(encrypt.RSA(publicKey: key));
    final encrypted = encrypter.encrypt(plainText);
    return encrypted.base64;
  }
}
