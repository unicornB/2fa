import 'package:ana_page_loop/ana_page_loop.dart' show anaPageLoop;
import 'package:flutter_2fa/app/routes/app_pages.dart';
import '../../config/app_config.dart' show AppConfig;
import '../tool/log_util.dart';

/// 初始化埋点统计插件
void anaPageLoopInit() {
  anaPageLoop.init(
    beginPageFn: (name) {
      // TODO: 第三方埋点统计开始
      LogUtil.d('待添加：埋点统计开始$name');
    },
    endPageFn: (name) {
      // TODO: 第三方埋点统计结束
      LogUtil.d('待添加：埋点统计结束$name');
    },
    routeRegExp: [Routes.HOME], // 过滤路由
    debug: AppConfig.DEBUG, // 路由调试
  );
}
