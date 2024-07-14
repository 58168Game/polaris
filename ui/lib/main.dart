import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:ui/activity.dart';
import 'package:ui/login_page.dart';
import 'package:ui/navdrawer.dart';
import 'package:ui/providers/APIs.dart';
import 'package:ui/search.dart';
import 'package:ui/system_settings.dart';
import 'package:ui/tv_details.dart';
import 'package:ui/weclome.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    // GoRouter configuration
    final _shellRoute = ShellRoute(
      builder: (BuildContext context, GoRouterState state, Widget child) {
        return SelectionArea(
          child: Scaffold(
              appBar: AppBar(
                // TRY THIS: Try changing the color here to a specific color (to
                // Colors.amber, perhaps?) and trigger a hot reload to see the AppBar
                // change color while the other colors stay the same.
                backgroundColor: Theme.of(context).colorScheme.inversePrimary,
                // Here we take the value from the MyHomePage object that was created by
                // the App.build method, and use it to set our appbar title.
                title: Row(
                  children: [
                    const Text("Polaris 追剧"),
                    const SizedBox(
                      width: 100,
                    ),
                    IconButton(
                        tooltip: "搜索剧集",
                        onPressed: () => context.go(SearchPage.route),
                        icon: const Icon(Icons.search)),
                  ],
                ),
                actions: [
                  IconButton(
                      onPressed: () => context.go(SystemSettingsPage.route),
                      icon: const Icon(Icons.settings))
                ],
              ),
              body: Center(
                  // Center is a layout widget. It takes a single child and positions it
                  // in the middle of the parent.
                  child: Row(children: <Widget>[
                const NavDrawer(),
                const VerticalDivider(thickness: 1, width: 1),
                Expanded(child: child)
              ]))),
        );
      },
      routes: [
        GoRoute(
          path: "/",
          redirect: (context, state) => WelcomePage.route,
        ),
        GoRoute(
          path: WelcomePage.route,
          builder: (context, state) => const WelcomePage(),
        ),
        GoRoute(
          path: SearchPage.route,
          builder: (context, state) => const SearchPage(),
        ),
        GoRoute(
          path: SystemSettingsPage.route,
          builder: (context, state) => const SystemSettingsPage(),
        ),
        GoRoute(
          path: TvDetailsPage.route,
          builder: (context, state) =>
              TvDetailsPage(seriesId: state.pathParameters['id']!),
        ),
        GoRoute(
          path: ActivityPage.route,
          builder: (context, state) => ActivityPage(),
        )
      ],
    );

    final _router = GoRouter(
      navigatorKey: APIs.navigatorKey,
      initialLocation: WelcomePage.route,
      routes: [
        _shellRoute,
        GoRoute(
          path: LoginScreen.route,
          builder: (context, state) => const LoginScreen(),
        )
      ],
    );

    return ProviderScope(
      child: MaterialApp.router(
        title: 'Polaris',
        theme: ThemeData(
          // This is the theme of your application.
          //
          // TRY THIS: Try running your application with "flutter run". You'll see
          // the application has a purple toolbar. Then, without quitting the app,
          // try changing the seedColor in the colorScheme below to Colors.green
          // and then invoke "hot reload" (save your changes or press the "hot
          // reload" button in a Flutter-supported IDE, or press "r" if you used
          // the command line to start the app).
          //
          // Notice that the counter didn't reset back to zero; the application
          // state is not lost during the reload. To reset the state, use hot
          // restart instead.
          //
          // This works for code too, not just values: Most code changes can be
          // tested with just a hot reload.
          colorScheme: ColorScheme.fromSeed(
              seedColor: Colors.blue, brightness: Brightness.dark),
          useMaterial3: true,
        ),
        routerConfig: _router,
      ),
    );
  }
}

CustomTransitionPage buildPageWithDefaultTransition<T>({
  required BuildContext context,
  required GoRouterState state,
  required Widget child,
}) {
  return CustomTransitionPage<T>(
    key: state.pageKey,
    child: child,
    transitionsBuilder: (context, animation, secondaryAnimation, child) =>
        FadeTransition(opacity: animation, child: child),
  );
}
