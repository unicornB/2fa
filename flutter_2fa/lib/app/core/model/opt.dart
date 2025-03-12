//2fa 信息
class Opt {
  String type;
  String label;
  String secret;
  String digits;
  String period;
  String algorithm;
  Opt({
    required this.type,
    required this.label,
    required this.secret,
    required this.digits,
    required this.period,
    required this.algorithm,
  });
}
