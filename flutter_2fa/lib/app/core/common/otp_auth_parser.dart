import 'dart:math';
import 'package:crypto/crypto.dart';
import 'package:base32/base32.dart';

class OTPAuthParser {
  late String type;
  late String label;
  late Map<String, String> parameters;

  OTPAuthParser(String url) {
    final uri = Uri.parse(url);
    if (uri.scheme != 'otpauth') {
      throw ArgumentError('Invalid OTPAuth URL: must start with otpauth://');
    }

    type = uri.host;
    label = uri.path.substring(1);
    parameters = uri.queryParameters;
  }

  String? get secret => parameters['secret'];
  int get digits => int.tryParse(parameters['digits'] ?? '6') ?? 6;
  int get period => int.tryParse(parameters['period'] ?? '30') ?? 30;
  String get algorithm => parameters['algorithm'] ?? 'SHA1';

  String generateOTP() {
    final now = DateTime.now().millisecondsSinceEpoch ~/ 1000;
    final counter = now ~/ period;
    final counterBytes = _intToBytes(counter);
    final secretBytes = base32.decode(secret!.replaceAll(' ', ''));

    final hmac = Hmac(sha1, secretBytes);
    final digest = hmac.convert(counterBytes);

    final offset = digest.bytes[19] & 0x0F;
    final binary = ((digest.bytes[offset] & 0x7F) << 24) |
        ((digest.bytes[offset + 1] & 0xFF) << 16) |
        ((digest.bytes[offset + 2] & 0xFF) << 8) |
        (digest.bytes[offset + 3] & 0xFF);

    final otp = binary % pow(10, digits).toInt();
    return otp.toString().padLeft(digits, '0');
  }

  List<int> _intToBytes(int value) {
    final bytes = List<int>.filled(8, 0);
    for (int i = 7; i >= 0; i--) {
      bytes[i] = value & 0xFF;
      value >>= 8;
    }
    return bytes;
  }
}
