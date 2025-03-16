// 初始化第三方插件

import '../../config/app_env.dart';

Future<void> appSetupInit() async {
  appEnv.init(); // 初始环境
}
