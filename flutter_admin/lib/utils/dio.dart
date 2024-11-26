import 'dart:convert';

import 'package:awesome_dialog/awesome_dialog.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

class DioUtils {
  // 创建 Dio 实例
  static final Dio _dio = Dio(BaseOptions(
    baseUrl: 'http://127.0.0.1:7891/api', // 设置基础 URL
    connectTimeout: Duration(milliseconds: 5000), // 连接超时
    receiveTimeout: Duration(milliseconds: 5000), // 响应超时
  ));

  // 设置请求头
  static void setHeaders(Map<String, String> headers) {
    _dio.options.headers.addAll(headers);
  }

  // GET 请求
  static void get(
    BuildContext context,
    String path, {
    Map<String, dynamic>? params,
    required Function(Map<String, dynamic>) onSuccess, // 成功回调
    required Function(Exception) onError, // 异常回调
    bool showLoading = true, // 是否显示加载弹窗
  }) async {
    AwesomeDialog? dialog;
    try {
      if (showLoading) {
        dialog = AwesomeDialog(
          context: context,
          dialogType: DialogType.noHeader,
          animType: AnimType.scale,
          body: Center(
            child: CircularProgressIndicator(),
          ),
          dismissOnTouchOutside: false,
          dismissOnBackKeyPress: false,
        )..show();
      }
      final response = await _dio.get(path, queryParameters: params);

      if (response.statusCode == 200) {
        onSuccess(json.decode(response.data));
      } else {
        throw Exception('请求失败，状态码: ${response.statusCode}');
      }
    } catch (e) {
      onError(e as Exception);
    } finally {
      if (showLoading && dialog != null) {
        dialog.dismiss();
      }
    }
  }

  // POST 请求
  static void post(
    BuildContext context,
    String path, {
    Map<String, dynamic>? params,
    Map<String, dynamic>? data,
    required Function(Map<String, dynamic>) onSuccess, // 成功回调
    required Function(Exception) onError, // 异常回调
    bool showLoading = true, // 是否显示加载弹窗
  }) async {
    AwesomeDialog? dialog;

    try {
      if (showLoading) {
        dialog = AwesomeDialog(
          context: context,
          dialogType: DialogType.noHeader,
          animType: AnimType.scale,
          body: Center(
            child: CircularProgressIndicator(),
          ),
          dismissOnTouchOutside: false,
          dismissOnBackKeyPress: false,
        )..show();
      }

      final response =
          await _dio.post(path, queryParameters: params, data: data);

      if (response.statusCode == 200) {
        onSuccess(json.decode(response.data));
      } else {
        throw Exception('请求失败，状态码: ${response.statusCode}');
      }
    } catch (e) {
      onError(e as Exception);
    } finally {
      if (showLoading && dialog != null) {
        dialog.dismiss();
      }
    }
  }
}
