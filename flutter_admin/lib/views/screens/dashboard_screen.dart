import 'package:flutter/material.dart';
import 'package:flutter_admin/constants/dimens.dart';
import 'package:flutter_admin/generated/l10n.dart';
import 'package:flutter_admin/theme/theme_extensions/app_color_scheme.dart';
import 'package:flutter_admin/theme/theme_extensions/app_data_table_theme.dart';
import 'package:flutter_admin/views/widgets/portal_master_layout/portal_master_layout.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({super.key});

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  @override
  Widget build(BuildContext context) {
    final lang = Lang.of(context);
    final themeData = Theme.of(context);
    final appColorScheme = Theme.of(context).extension<AppColorScheme>()!;
    final appDataTableTheme = Theme.of(context).extension<AppDataTableTheme>()!;
    final size = MediaQuery.of(context).size;
    final summaryCardCrossAxisCount = (size.width >= kScreenWidthLg ? 4 : 2);
    return PortalMasterLayout(
      body: ListView(
        padding: const EdgeInsets.all(kDefaultPadding),
        children: [
          Text(
            lang.dashboard,
            style: themeData.textTheme.headlineMedium,
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: kDefaultPadding),
            child: LayoutBuilder(builder: (context, constraints) {
              final summaryCardWidth = ((constraints.maxWidth -
                      (kDefaultPadding * (summaryCardCrossAxisCount - 1))) /
                  summaryCardCrossAxisCount);
              return Wrap(
                direction: Axis.horizontal,
                spacing: kDefaultPadding,
                runSpacing: kDefaultPadding,
                children: [
                  SummaryCard(
                    title: lang.newOrders(2),
                    value: '150',
                    icon: Icons.shopping_cart_rounded,
                    backgroundColor: appColorScheme.info,
                    textColor: themeData.colorScheme.onPrimary,
                    iconColor: Colors.black12,
                    width: summaryCardWidth,
                  ),
                ],
              );
            }),
          )
        ],
      ),
    );
  }
}

class SummaryCard extends StatelessWidget {
  final String title;
  final String value;
  final IconData icon;
  final Color backgroundColor;
  final Color textColor;
  final Color iconColor;
  final double width;

  const SummaryCard({
    super.key,
    required this.title,
    required this.value,
    required this.icon,
    required this.backgroundColor,
    required this.textColor,
    required this.iconColor,
    required this.width,
  });

  @override
  Widget build(BuildContext context) {
    final textTheme = Theme.of(context).textTheme;

    return SizedBox(
      height: 120.0,
      width: width,
      child: Card(
        clipBehavior: Clip.antiAlias,
        color: backgroundColor,
        child: Stack(
          children: [
            Positioned(
              top: kDefaultPadding * 0.5,
              right: kDefaultPadding * 0.5,
              child: Icon(
                icon,
                size: 80.0,
                color: iconColor,
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(kDefaultPadding),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Padding(
                    padding:
                        const EdgeInsets.only(bottom: kDefaultPadding * 0.5),
                    child: Text(
                      value,
                      style: textTheme.headlineMedium!.copyWith(
                        color: textColor,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),
                  Text(
                    title,
                    style: textTheme.labelLarge!.copyWith(
                      color: textColor,
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
