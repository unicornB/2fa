import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

extension StringExtension on String {
  SvgPicture toSvg({double height = 15, double width = 15, Color? color}) =>
      SvgPicture.asset(
        this,
        colorFilter: ColorFilter.mode(color ?? Colors.black, BlendMode.srcIn),
        width: width,
        height: width,
        semanticsLabel: 'A red up arrow',
      );
  Image toAssetImage({double height = 50, double width = 50}) =>
      Image.asset(this, height: height, width: width);
}
