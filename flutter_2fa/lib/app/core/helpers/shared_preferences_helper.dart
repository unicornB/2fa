// import 'package:shared_preferences/shared_preferences.dart';

// class SharedPreferencesHelper {
//   static final SharedPreferencesHelper _instance =
//       SharedPreferencesHelper._internal();
//   factory SharedPreferencesHelper() => _instance;

//   late SharedPreferences _prefs;

//   SharedPreferencesHelper._internal();

//   // 初始化 SharedPreferences
//   Future<void> init() async {
//     _prefs = await SharedPreferences.getInstance();
//   }

//   // 获取字符串值
//   String? getString(String key) {
//     return _prefs.getString(key);
//   }

//   // 设置字符串值
//   Future<bool> setString(String key, String value) {
//     return _prefs.setString(key, value);
//   }

//   // 获取布尔值
//   bool? getBool(String key) {
//     return _prefs.getBool(key);
//   }

//   // 设置布尔值
//   Future<bool> setBool(String key, bool value) {
//     return _prefs.setBool(key, value);
//   }

//   // 获取整数值
//   int? getInt(String key) {
//     return _prefs.getInt(key);
//   }

//   // 设置整数值
//   Future<bool> setInt(String key, int value) {
//     return _prefs.setInt(key, value);
//   }

//   // 获取双精度浮点数值
//   double? getDouble(String key) {
//     return _prefs.getDouble(key);
//   }

//   // 设置双精度浮点数值
//   Future<bool> setDouble(String key, double value) {
//     return _prefs.setDouble(key, value);
//   }

//   // 获取字符串列表值
//   List<String>? getStringList(String key) {
//     return _prefs.getStringList(key);
//   }

//   // 设置字符串列表值
//   Future<bool> setStringList(String key, List<String> value) {
//     return _prefs.setStringList(key, value);
//   }

//   // 删除键值对
//   Future<bool> remove(String key) {
//     return _prefs.remove(key);
//   }

//   // 清除所有键值对
//   Future<bool> clear() {
//     return _prefs.clear();
//   }
// }
