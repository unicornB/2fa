import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_2fa/app/core/extensions/rpx_int_extendsion.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class OtpItem extends StatefulWidget {
  const OtpItem({super.key});

  @override
  State<OtpItem> createState() => _OtpItemState();
}

class _OtpItemState extends State<OtpItem> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(
        vertical: 10.rpx,
        horizontal: 20.rpx,
      ),
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(
          bottom: BorderSide(
            color: Color(0xffDBDEDB),
            width: 1.rpx,
          ),
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text("Google yeyyeyeyyeyey@gmail.com"),
          SizedBox(
            height: 10,
          ),
          Row(
            children: [
              Text(
                "767677",
                style: TextStyle(
                  fontSize: 50.rpx,
                  fontWeight: FontWeight.bold,
                ),
              )
            ],
          )
        ],
      ),
    );
  }
}
