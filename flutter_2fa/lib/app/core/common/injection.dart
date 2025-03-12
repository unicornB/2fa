import 'package:get/get.dart';

import 'package:shared_preferences/shared_preferences.dart';

import '../theme/color_palettes.dart';

///初始化注入对象
class Injection extends GetxService {
  Future<void> init() async {
    await Get.putAsync(() => SharedPreferences.getInstance());
    ColorPalettes.instance.init();
  }
}
