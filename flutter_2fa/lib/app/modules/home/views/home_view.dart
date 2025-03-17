import 'dart:developer';
import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/extensions/rpx_int_extendsion.dart';
import 'package:flutter_2fa/app/core/theme/color_palettes.dart';
import 'package:flutter_2fa/app/modules/home/views/components/otp_item.dart';
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
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: ColorPalettes.instance.primary,
        leading: IconButton(
          onPressed: () {},
          icon: "assets/svg/menu-2-line.svg".toSvg(
            width: 50.rpx,
            height: 50.rpx,
            color: Colors.white,
          ),
        ),
        title: Text(
          "星辰验证器",
          style: TextStyle(
            color: Colors.white,
            fontSize: 36.rpx,
          ),
        ),
        centerTitle: true,
        actions: [
          IconButton(
            onPressed: () async {},
            icon: "assets/svg/search.svg".toSvg(
              width: 50.rpx,
              color: Colors.white,
            ),
          ),
          IconButton(
            onPressed: () async {
              var result = await BarcodeScanner.scan();
              log(result.rawContent);
              if (result.rawContent == '') {
                return;
              }
              final parser = OTPAuthParser(result.rawContent);
              log('Type: ${parser.type}');
              log('Label: ${parser.label}');
              log('Secret: ${parser.secret}');
              log('Digits: ${parser.digits}');
              log('Period: ${parser.period}');
              log('Algorithm: ${parser.algorithm}');
              log('OTP: ${parser.generateOTP()}');
            },
            icon: "assets/svg/add.svg".toSvg(
              width: 50.rpx,
              color: Colors.white,
            ),
          ),
          SizedBox(
            width: 20.rpx,
          ),
        ],
      ),
      body: Container(
        padding: EdgeInsets.symmetric(horizontal: 30.rpx),
        child: ListView.builder(
          itemBuilder: (context, index) {
            return OtpItem();
          },
          itemCount: 10,
        ),
      ),
    );
  }

  Widget _searchBar() {
    return SizedBox(
      height: 70.rpx,
      child: TextField(
        decoration: InputDecoration(
          contentPadding: EdgeInsets.symmetric(vertical: 0, horizontal: 20),
          hintText: '搜索',
          hintStyle: TextStyle(
            color: Colors.white,
            fontSize: 28.rpx,
          ),
          prefixIcon: Icon(
            Remix.search_line,
            color: Colors.white,
            size: 40.rpx,
          ),
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(40.rpx),
            borderSide: BorderSide(
              color: Colors.white,
            ),
          ),
          focusedBorder: OutlineInputBorder(
            borderRadius: BorderRadius.circular(40.rpx),
            borderSide: BorderSide(
              color: Colors.white,
            ),
          ),
        ),
      ),
    );
  }
}
