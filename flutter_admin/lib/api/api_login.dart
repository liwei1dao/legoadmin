import 'package:flutter/material.dart';
import 'package:flutter_admin/utils/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiLogin {
  static Future<void> login(
    BuildContext context,
    String username,
    String password, {
    required Function onSuccess,
    required Function onError,
  }) async {
    DioUtils.post(
      context,
      '/login', // 假设登录接口路径为 /login
      data: {
        'username': username,
        'password': password,
      },
      onSuccess: (data) async {
        // 保存登录信息到本地存储
        SharedPreferences prefs = await SharedPreferences.getInstance();
        await prefs.setString('token', data['token']);
        onSuccess();
      },
      onError: (e) {
        onError(e);
      },
    );
  }

  static Future<void> logout() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.remove('token');
  }

  static Future<String?> getToken() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.getString('token');
  }
}
