import 'package:flutter/material.dart';
import 'package:flutter_admin/constants/values.dart';
import 'package:shared_preferences/shared_preferences.dart';

class UserDataProvider extends ChangeNotifier {
  var _userProfileImageUrl = '';
  var _username = '';

  String get userProfileImageUrl => _userProfileImageUrl;
  String get username => _username;

  Future<void> loadAsync() async {
    final sharedPref = await SharedPreferences.getInstance();

    _username = sharedPref.getString(StorageKeys.username) ?? '';
    _userProfileImageUrl =
        sharedPref.getString(StorageKeys.userProfileImageUrl) ?? '';

    notifyListeners();
  }

  bool isUserLoggedIn() {
    return _username.isNotEmpty;
  }
}
