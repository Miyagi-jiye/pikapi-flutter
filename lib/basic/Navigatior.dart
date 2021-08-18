import 'dart:async';

import 'package:flutter/material.dart';

Future<dynamic> navPushOrReplace(
    BuildContext context, WidgetBuilder builder) async {
  if (_depth < _depthMax) {
    return Navigator.push(
      context,
      MaterialPageRoute(builder: builder),
    );
  } else {
    return Navigator.pushReplacement(
      context,
      MaterialPageRoute(builder: builder),
    );
  }
}

var navigatorObserver = _NavigatorObserver();

const _depthMax = 15;
var _depth = 0;

class _NavigatorObserver extends NavigatorObserver {
  @override
  void didPop(Route route, Route? previousRoute) {
    _depth--;
    print("DEPTH : $_depth");
    super.didPop(route, previousRoute);
  }

  @override
  void didPush(Route route, Route? previousRoute) {
    _depth++;
    print("DEPTH : $_depth");
    super.didPush(route, previousRoute);
  }
}

Future circularPush(BuildContext context, Widget? widget) async {
  while (widget != null) {
    widget = await Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => widget!,
      ),
    );
  }
}
