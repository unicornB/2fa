import 'package:flutter_easyloading/flutter_easyloading.dart';

class Tips {
  /// tosat常规提示
  static Future<void> info(
    String text, {
    EasyLoadingToastPosition? toastPosition,
  }) async {
    EasyLoading.showToast(
      text,
      toastPosition: toastPosition ?? EasyLoadingToastPosition.center,
    );
  }
}
