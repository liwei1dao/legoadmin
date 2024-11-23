import 'package:flutter/material.dart';
import 'package:flutter_admin/app_router.dart';
import 'package:flutter_admin/generated/l10n.dart';
import 'package:flutter_admin/providers/app_preferences_provider.dart';
import 'package:flutter_admin/providers/user_data_provider.dart';
import 'package:flutter_admin/theme/themes.dart';
import 'package:flutter_admin/utils/app_focus_helper.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';

class RootApp extends StatefulWidget {
  const RootApp({super.key});
  @override
  State<RootApp> createState() => _RootAppState();
}

class _RootAppState extends State<RootApp> {
  GoRouter? _appRouter;
  Future<bool>? _future;

  Future<bool> _getScreenDataAsync(
      AppPreferencesProvider appPreferencesProvider,
      UserDataProvider userDataProvider) async {
    await appPreferencesProvider.loadAsync();
    await userDataProvider.loadAsync();

    return true;
  }

  @override
  Widget build(Object context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (context) => AppPreferencesProvider()),
        ChangeNotifierProvider(create: (context) => UserDataProvider()),
      ],
      child: Builder(
        builder: (BuildContext context) {
          return GestureDetector(
            onTap: () {
              AppFocusHelper.instance.requestUnfocus();
            },
            child: FutureBuilder<bool>(
                initialData: null,
                future: (_future ??= _getScreenDataAsync(
                    context.read<AppPreferencesProvider>(),
                    context.read<UserDataProvider>())),
                builder: (context, snapshot) {
                  if (snapshot.hasData && snapshot.data!) {
                    return Consumer<AppPreferencesProvider>(
                      builder: (context, provider, child) {
                        _appRouter ??=
                            appRouter(context.read<UserDataProvider>());

                        return MaterialApp.router(
                          debugShowCheckedModeBanner: false,
                          routeInformationProvider:
                              _appRouter!.routeInformationProvider,
                          routeInformationParser:
                              _appRouter!.routeInformationParser,
                          routerDelegate: _appRouter!.routerDelegate,
                          supportedLocales: Lang.delegate.supportedLocales,
                          localizationsDelegates: const [
                            Lang.delegate,
                            GlobalMaterialLocalizations.delegate,
                            GlobalWidgetsLocalizations.delegate,
                            GlobalCupertinoLocalizations.delegate,
                          ],
                          locale: provider.locale,
                          onGenerateTitle: (context) =>
                              Lang.of(context).appTitle,
                          theme: AppThemeData.instance.light(),
                          darkTheme: AppThemeData.instance.dark(),
                          themeMode: provider.themeMode,
                        );
                      },
                    );
                  }
                  return const SizedBox.shrink();
                }),
          );
        },
      ),
    );
  }
}
