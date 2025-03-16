import 'package:flutter/material.dart';
import 'package:flutter_2fa/app/core/extensions/rpx_int_extendsion.dart';
import 'package:flutter_2fa/app/core/extensions/string_extension.dart';
import 'package:flutter_2fa/app/core/theme/color_palettes.dart';

import 'package:get/get.dart';

import '../controllers/login_controller.dart';

class LoginView extends GetView<LoginController> {
  const LoginView({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ColorPalettes.instance.background,
      body: SingleChildScrollView(
        child: Obx(() => _formBody()),
      ),
    );
  }

  Widget _formBody() {
    return Container(
      padding: EdgeInsets.all(40.rpx),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            height: 300.rpx,
          ),
          Row(
            children: [
              "assets/images/logo.png".toAssetImage(width: 70.rpx),
              SizedBox(
                width: 10.rpx,
              ),
              Text(
                "星辰验证器",
                style: TextStyle(
                  fontSize: 50.rpx,
                  fontWeight: FontWeight.bold,
                  color: ColorPalettes.instance.firstText,
                ),
              ),
            ],
          ),
          SizedBox(
            height: 20.rpx,
          ),
          Text(
            "  欢迎使用星辰验证器",
            style: TextStyle(
              fontSize: 30.rpx,
              color: ColorPalettes.instance.secondText,
            ),
          ),
          SizedBox(
            height: 120.rpx,
          ),
          SizedBox(
            height: 90.rpx,
            child: TextField(
              controller: controller.emailController,
              keyboardType: TextInputType.emailAddress,
              decoration: InputDecoration(
                contentPadding:
                    EdgeInsets.symmetric(vertical: 10.rpx, horizontal: 20),
                hintText: '请输入邮箱',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20.rpx),
                ),
                focusedBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20.rpx),
                  borderSide: BorderSide(
                    color: ColorPalettes.instance.firstText,
                  ),
                ),
              ),
            ),
          ),
          SizedBox(
            height: 50.rpx,
          ),
          SizedBox(
            height: 90.rpx,
            child: TextField(
              controller: controller.codeController,
              decoration: InputDecoration(
                contentPadding:
                    EdgeInsets.symmetric(vertical: 20.rpx, horizontal: 20),
                hintText: '请输入验证码',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20.rpx),
                ),
                focusedBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20.rpx),
                  borderSide: BorderSide(
                    color: ColorPalettes.instance.firstText,
                  ),
                ),
                suffixIcon: Transform.scale(
                  scale: 0.8, // 缩放比例
                  child: Padding(
                    padding: EdgeInsets.only(right: 6.rpx),
                    child: ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: Size(100.rpx, 40.rpx),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(20.rpx),
                        ),
                        backgroundColor: controller.sendBtnDisable.value
                            ? ColorPalettes.instance.btnDisableBackground
                            : ColorPalettes.instance.primary,
                      ),
                      onPressed: () {
                        if (controller.sendBtnDisable.value) return;
                        controller.sendCode();
                      },
                      child: Text(
                        controller.count.value > 0
                            ? "${controller.count.value}s"
                            : "发送验证码",
                        style: TextStyle(
                          color: ColorPalettes.instance.btnDisableText,
                        ),
                      ),
                    ),
                  ),
                ),
              ),
            ),
          ),
          SizedBox(
            height: 50.rpx,
          ),
          SizedBox(
            height: 90.rpx,
            width: double.infinity,
            child: ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: controller.submitBtnDisable.value
                    ? ColorPalettes.instance.btnDisableBackground
                    : ColorPalettes.instance.primary,
              ),
              onPressed: () {
                if (controller.submitBtnDisable.value) return;
                controller.login();
              },
              child: Text(
                "登录",
                style: TextStyle(color: ColorPalettes.instance.btnDisableText),
              ),
            ),
          ),
          SizedBox(
            height: 30.rpx,
          ),
          //[]我已阅读并同意星辰验证器用户协议
          Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Checkbox(
                activeColor: ColorPalettes.instance.primary,
                value: controller.checked.value,
                onChanged: (value) {
                  controller.checked.value = value!;
                  controller.changeSubmitBtnDisable();
                },
              ),
              Text(
                "我已阅读并同意",
                style: TextStyle(
                  fontSize: 28.rpx,
                  color: ColorPalettes.instance.secondText,
                ),
              ),
              SizedBox(
                width: 10.rpx,
              ),
              Text(
                "服务协议",
                style: TextStyle(
                  fontSize: 28.rpx,
                  color: ColorPalettes.instance.primary,
                ),
              ),
              SizedBox(
                width: 10.rpx,
              ),
              Text(
                "和",
                style: TextStyle(
                  fontSize: 28.rpx,
                  color: ColorPalettes.instance.secondText,
                ),
              ),
              SizedBox(
                width: 10.rpx,
              ),
              Text(
                "隐私政策",
                style: TextStyle(
                  fontSize: 28.rpx,
                  color: ColorPalettes.instance.primary,
                ),
              )
            ],
          )
        ],
      ),
    );
  }
}
