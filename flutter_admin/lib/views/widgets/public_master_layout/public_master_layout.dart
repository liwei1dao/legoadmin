import 'package:flutter/material.dart';

class PublicMasterLayout extends StatelessWidget {
  final Widget body;
  const PublicMasterLayout({
    super.key,
    required this.body,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [SizedBox(height: kToolbarHeight, child: Row())],
      ),
    );
  }
}
