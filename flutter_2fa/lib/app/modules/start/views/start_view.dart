import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/theme/color_palettes.dart';

import 'package:get/get.dart';
import 'package:lottie/lottie.dart';

import '../controllers/start_controller.dart';

class StartView extends GetView<StartController> {
  const StartView({super.key});
  @override
  Widget build(BuildContext context) {
    controller.toHome();
    return Scaffold(
      backgroundColor: ColorPalettes.instance.background,
      body: Center(
        child: Lottie.asset('assets/lotties/start.json'),
      ),
    );
  }
}
