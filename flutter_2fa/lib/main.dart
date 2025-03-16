import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/service/user_service.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import 'package:get/get.dart';

import 'app/core/common/injection.dart';

import 'app/core/utils/app_setup/index.dart';
import 'app/core/utils/size_fit/size_fit.dart';
import 'app/routes/app_pages.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initService();
  runApp(
    GetMaterialApp(
      title: "星宸验证器",
      initialRoute: AppPages.INITIAL,
      getPages: AppPages.routes,
      builder: EasyLoading.init(),
    ),
  );
}

Future<void> initService() async {
  await ScreenUtil.ensureScreenSize();
  SizeFit.initialize();
  await Injection().init();
  await appSetupInit();
  await Get.putAsync(() => UserService().init());
  //UserService.to.getUserInfo();
}
