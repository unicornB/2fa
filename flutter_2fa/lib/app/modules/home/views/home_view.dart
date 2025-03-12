import 'dart:developer';
import 'package:flutter/material.dart';

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
      appBar: AppBar(
        leading: IconButton(onPressed: () {}, icon: Icon(Remix.menu_2_line)),
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
            icon: Icon(
              Remix.qr_scan_line,
              color: Color(0xff336DF3),
              size: 30,
            ),
          ),
          SizedBox(
            width: 10,
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
