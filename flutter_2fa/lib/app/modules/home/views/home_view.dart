import 'dart:developer';
import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/extensions/rpx_int_extendsion.dart';
import 'package:flutter_2fa/app/core/theme/color_palettes.dart';
import '/app/core/extensions/string_extension.dart';

import 'package:get/get.dart';
import 'package:barcode_scan2/barcode_scan2.dart';
import 'package:remixicon/remixicon.dart';
import '../../../core/common/otp_auth_parser.dart';
import '../controllers/home_controller.dart';

class HomeView extends GetView<HomeController> {
  const HomeView({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorPalettes.instance.background,
      appBar: AppBar(
        backgroundColor: ColorPalettes.instance.background,
        leading: IconButton(
          onPressed: () {},
          icon:
              "assets/svg/menu-2-line.svg".toSvg(width: 60.rpx, height: 60.rpx),
        ),
        title: _searchBar(),
        centerTitle: true,
        actions: [
          IconButton(
            onPressed: () async {
              var result = await BarcodeScanner.scan();
              log(result.rawContent);
              final parser = OTPAuthParser(result.rawContent);
              log('Type: ${parser.type}');
              log('Label: ${parser.label}');
              log('Secret: ${parser.secret}');
              log('Digits: ${parser.digits}');
              log('Period: ${parser.period}');
              log('Algorithm: ${parser.algorithm}');
              log('OTP: ${parser.generateOTP()}');
            },
            icon: "assets/svg/scan.svg".toSvg(width: 60.rpx),
          ),
          SizedBox(
            width: 20.rpx,
          )
        ],
      ),
      body: Container(
        padding: const EdgeInsets.all(20),
        child: ListView(
          shrinkWrap: true,
        ),
      ),
    );
  }

  Widget _searchBar() {
    return SizedBox(
      height: 45,
      child: TextField(
        decoration: InputDecoration(
          contentPadding: EdgeInsets.symmetric(vertical: 0, horizontal: 20),
          hintText: '搜索',
          prefixIcon: Icon(
            Remix.search_line,
            color: Colors.black,
            size: 24,
          ),
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(20),
          ),
        ),
      ),
    );
  }
}
