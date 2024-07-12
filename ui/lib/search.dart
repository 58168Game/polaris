import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:ui/providers/APIs.dart';
import 'package:ui/providers/settings.dart';
import 'package:ui/providers/welcome_data.dart';

class SearchPage extends ConsumerStatefulWidget {
  const SearchPage({super.key});

  static const route = "/search";

  @override
  ConsumerState<ConsumerStatefulWidget> createState() {
    return _SearchPageState();
  }
}

class _SearchPageState extends ConsumerState<SearchPage> {
  List<dynamic> list = List.empty();

  @override
  Widget build(BuildContext context) {
    var searchList = ref.watch(searchPageDataProvider);

    List<Widget> res = searchList.when(
        data: (data) {
          var cards = List<Widget>.empty(growable: true);
          for (final item in data) {
            cards.add(Card(
                margin: const EdgeInsets.all(4),
                clipBehavior: Clip.hardEdge,
                child: InkWell(
                  //splashColor: Colors.blue.withAlpha(30),
                  onTap: () {
                    //showDialog(context: context, builder: builder)
                    _showSubmitDialog(context, item);
                  },
                  child: Row(
                    children: <Widget>[
                      Flexible(
                        child: SizedBox(
                          width: 150,
                          height: 200,
                          child: Image.network(
                            APIs.tmdbImgBaseUrl + item.posterPath!,
                            fit: BoxFit.contain,
                          ),
                        ),
                      ),
                      Flexible(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              "${item.name} (${item.firstAirDate?.split("-")[0]})",
                              style: const TextStyle(
                                  fontSize: 14, fontWeight: FontWeight.bold),
                            ),
                            const Text(""),
                            Text(item.overview!)
                          ],
                        ),
                      )
                    ],
                  ),
                )));
          }
          return cards;
        },
        error: (err, trace) => [Text("$err")],
        loading: () => [const CircularProgressIndicator()]);
    return Column(
      children: [
        TextField(
          autofocus: true,
          onSubmitted: (value) {
            ref.read(searchPageDataProvider.notifier).queryResults(value);
          },
          decoration: const InputDecoration(
              labelText: "搜索",
              hintText: "搜索剧集名称",
              prefixIcon: Icon(Icons.search)),
        ),
        Expanded(
            child: ListView(
          children: res,
        ))
      ],
    );
  }

  Future<void> _showSubmitDialog(BuildContext context, SearchResult item) {
    TextEditingController resolutionController =
        TextEditingController(text: "1080p");
    TextEditingController storageController = TextEditingController();
    var storage = ref.watch(storageSettingProvider);
    return showDialog<void>(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text('添加剧集: ${item.name}'),
            content: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                DropdownMenu(
                  label: const Text("清晰度"),
                  controller: resolutionController,
                  dropdownMenuEntries: const [
                    DropdownMenuEntry(value: "720p", label: "720p"),
                    DropdownMenuEntry(value: "1080p", label: "1080p"),
                    DropdownMenuEntry(value: "4k", label: "4k"),
                  ],
                ),
                storage.when(
                    data: (v) {
                      return DropdownMenu(
                        label: const Text("存储位置"),
                        controller: storageController,
                        dropdownMenuEntries: v
                            .map((s) =>
                                DropdownMenuEntry(label: s.name!, value: s.id))
                            .toList(),
                      );
                    },
                    error: (err, trace) => Text("$err"),
                    loading: () => const CircularProgressIndicator()),
              ],
            ),
            actions: <Widget>[
              TextButton(
                style: TextButton.styleFrom(
                  textStyle: Theme.of(context).textTheme.labelLarge,
                ),
                child: const Text('取消'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
              TextButton(
                style: TextButton.styleFrom(
                  textStyle: Theme.of(context).textTheme.labelLarge,
                ),
                child: const Text('确定'),
                onPressed: () {
                  ref
                      .read(searchPageDataProvider.notifier)
                      .submit2Watchlist(item.id!);
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        });
  }
}

class SearchBarApp extends StatefulWidget {
  const SearchBarApp({
    super.key,
    required this.onChanged,
  });

  final ValueChanged<String> onChanged;
  @override
  State<SearchBarApp> createState() => _SearchBarAppState();
}

class _SearchBarAppState extends State<SearchBarApp> {
  @override
  Widget build(BuildContext context) {
    return SearchAnchor(
        builder: (BuildContext context, SearchController controller) {
      return SearchBar(
        controller: controller,
        padding: const WidgetStatePropertyAll<EdgeInsets>(
            EdgeInsets.symmetric(horizontal: 16.0)),
        onSubmitted: (value) => {widget.onChanged(controller.text)},
        leading: const Icon(Icons.search),
      );
    }, suggestionsBuilder: (BuildContext context, SearchController controller) {
      return List<ListTile>.generate(0, (int index) {
        final String item = 'item $index';
        return ListTile(
          title: Text(item),
          onTap: () {
            setState(() {
              controller.closeView(item);
            });
          },
        );
      });
    });
  }
}
