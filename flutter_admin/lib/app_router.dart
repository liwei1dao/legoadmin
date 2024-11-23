import 'package:flutter_admin/providers/user_data_provider.dart';
import 'package:flutter_admin/views/screens/error_screen.dart';
import 'package:go_router/go_router.dart';

class RouteUri {
  static const String home = '/';
  static const String dashboard = '/dashboard';
  static const String myProfile = '/my-profile';
  static const String logout = '/logout';
  static const String form = '/form';
  static const String generalUi = '/general-ui';
  static const String colors = '/colors';
  static const String text = '/text';
  static const String buttons = '/buttons';
  static const String dialogs = '/dialogs';
  static const String error404 = '/404';
  static const String login = '/login';
  static const String register = '/register';
  static const String crud = '/crud';
  static const String crudDetail = '/crud-detail';
  static const String iframe = '/iframe';
}

GoRouter appRouter(UserDataProvider userDataProvider) {
  return GoRouter(
      initialLocation: RouteUri.home,
      errorPageBuilder: (context, state) => NoTransitionPage<void>(
            key: state.pageKey,
            child: const ErrorScreen(),
          ),
      routes: []);
}
