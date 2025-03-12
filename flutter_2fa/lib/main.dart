import 'package:flutter/material.dart';

import 'package:get/get.dart';
import 'package:jh_debug/jh_debug.dart';

import 'app/core/common/injection.dart';
import 'app/core/config/common_config.dart';
import 'app/core/utils/app_setup/index.dart';
import 'app/routes/app_pages.dart';

void main() {
  jhDebugMain(
    appChild: const MyApp(),
    debugMode: DebugMode.inConsole,
    errorCallback: (details) {},
  );
  // runApp(
  //   GetMaterialApp(
  //     title: "Application",
  //     initialRoute: AppPages.INITIAL,
  //     getPages: AppPages.routes,
  //   ),
  // );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext context) {
    jhDebug.setGlobalKey = commonConfig.getGlobalKey;
    WidgetsFlutterBinding.ensureInitialized();
    initService();
    return GetMaterialApp(
      title: "双因子认证",
      initialRoute: AppPages.INITIAL,
      getPages: AppPages.routes,
    );
  }

  Future<void> initService() async {
    await Injection().init();
    appSetupInit();
  }
}
