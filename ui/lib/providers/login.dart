import 'dart:async';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:ui/providers/APIs.dart';
import 'package:ui/providers/server_response.dart';

var authSettingProvider =
    AsyncNotifierProvider.autoDispose<AuthSettingData, AuthSetting>(
        AuthSettingData.new);

class AuthSettingData extends AutoDisposeAsyncNotifier<AuthSetting> {
  @override
  FutureOr<AuthSetting> build() async {
    final dio = await APIs.getDio();
    var resp = await dio.get(APIs.loginSettingUrl);
    var sp = ServerResponse.fromJson(resp.data);
    if (sp.code != 0) {
      throw sp.message;
    }
    var as = AuthSetting.fromJson(sp.data);
    return as;
  }

  Future<void> updateAuthSetting(
      bool enable, String user, String password) async {
    final dio = await APIs.getDio();
    var resp = await dio.post(APIs.loginSettingUrl,
        data: {"enable": enable, "user": user, "password": password});
    var sp = ServerResponse.fromJson(resp.data);
    if (sp.code != 0) {
      throw sp.message;
    }
    ref.invalidateSelf();
  }

  Future<void> login(String user, String password) async {
    var resp = await Dio()
        .post(APIs.loginUrl, data: {"user": user, "password": password});

    var sp = ServerResponse.fromJson(resp.data);

    if (sp.code != 0) {
      throw sp.message;
    }
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    prefs.setString("token", sp.data["token"]);
  }
}

class AuthSetting {
  bool enable;
  String user;

  AuthSetting({required this.enable, required this.user});

  factory AuthSetting.fromJson(Map<String, dynamic> json) {
    return AuthSetting(enable: json["enable"], user: json["user"]);
  }
}