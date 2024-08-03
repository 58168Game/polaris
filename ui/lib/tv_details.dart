import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:ui/providers/series_details.dart';
import 'package:ui/widgets/detail_card.dart';
import 'package:ui/widgets/utils.dart';
import 'package:ui/widgets/progress_indicator.dart';
import 'package:ui/widgets/widgets.dart';

class TvDetailsPage extends ConsumerStatefulWidget {
  static const route = "/series/:id";

  static String toRoute(int id) {
    return "/series/$id";
  }

  final String seriesId;

  const TvDetailsPage({super.key, required this.seriesId});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() {
    return _TvDetailsPageState();
  }
}

class _TvDetailsPageState extends ConsumerState<TvDetailsPage> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    var seriesDetails = ref.watch(mediaDetailsProvider(widget.seriesId));
    return seriesDetails.when(
        data: (details) {
          Map<int, List<DataRow>> m = {};
          for (final ep in details.episodes!) {
            var row = DataRow(cells: [
              DataCell(Text("${ep.episodeNumber}")),
              DataCell(Text("${ep.title}")),
              DataCell(Opacity(
                opacity: 0.5,
                child: Text(ep.airDate ?? "-"),
              )),
              DataCell(
                Opacity(
                    opacity: 0.7,
                    child: ep.status == "downloading"
                        ? const Tooltip(
                            message: "下载中",
                            child: Icon(Icons.downloading),
                          )
                        : (ep.status == "downloaded"
                            ? const Tooltip(
                                message: "已下载",
                                child: Icon(Icons.download_done),
                              )
                            : (ep.monitored == true
                                ? Tooltip(
                                    message: "监控中",
                                    child: IconButton(
                                        onPressed: () {
                                          ref
                                              .read(mediaDetailsProvider(
                                                      widget.seriesId)
                                                  .notifier)
                                              .changeMonitoringStatus(
                                                  ep.id!, false);
                                        },
                                        icon: const Icon(Icons.alarm)),
                                  )
                                : Opacity(
                                    opacity: 0.7,
                                    child: Tooltip(
                                      message: "未监控",
                                      child: IconButton(
                                          onPressed: () {
                                            ref
                                                .read(mediaDetailsProvider(
                                                        widget.seriesId)
                                                    .notifier)
                                                .changeMonitoringStatus(
                                                    ep.id!, true);
                                          },
                                          icon: const Icon(Icons.alarm_off)),
                                    ),
                                  )))),
              ),
              DataCell(Row(
                children: [
                  Tooltip(
                    message: "搜索下载对应剧集",
                    child: IconButton(
                        onPressed: () {
                          var f = ref
                              .read(mediaDetailsProvider(widget.seriesId)
                                  .notifier)
                              .searchAndDownload(widget.seriesId,
                                  ep.seasonNumber!, ep.episodeNumber!)
                              .then((v) => showSnakeBar("开始下载: $v"));
                          showLoadingWithFuture(f);
                        },
                        icon: const Icon(Icons.download)),
                  ),
                  const SizedBox(
                    width: 10,
                  ),
                  Tooltip(message: "查看可用资源",child: IconButton(
                      onPressed: () => showAvailableTorrents(widget.seriesId,
                          ep.seasonNumber ?? 0, ep.episodeNumber ?? 0),
                      icon: const Icon(Icons.manage_search)),)
                ],
              ))
            ]);

            if (m[ep.seasonNumber] == null) {
              m[ep.seasonNumber!] = List.empty(growable: true);
            }
            m[ep.seasonNumber!]!.add(row);
          }
          List<ExpansionTile> list = List.empty(growable: true);
          for (final k in m.keys.toList().reversed) {
            var seasonList = ExpansionTile(
              tilePadding: const EdgeInsets.fromLTRB(10, 0, 10, 0),
              //childrenPadding: const EdgeInsets.fromLTRB(50, 0, 50, 0),
              initiallyExpanded: false,
              title: k == 0 ? const Text("特别篇") : Text("第 $k 季"),
              expandedCrossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                DataTable(columns: [
                  const DataColumn(label: Text("#")),
                  const DataColumn(
                    label: Text("标题"),
                  ),
                  const DataColumn(label: Text("播出时间")),
                  const DataColumn(label: Text("状态")),
                  DataColumn(
                      label: Row(
                    children: [
                      Tooltip(
                        message: "搜索下载全部剧集",
                        child: IconButton(
                            onPressed: () {
                              final f = ref
                                  .read(mediaDetailsProvider(widget.seriesId)
                                      .notifier)
                                  .searchAndDownload(widget.seriesId, k, 0)
                                  .then((v) => showSnakeBar("开始下载: $v"));
                              showLoadingWithFuture(f);
                            },
                            icon: const Icon(Icons.download)),
                      ),
                      const SizedBox(
                        width: 10,
                      ),
                      Tooltip(message: "查看可用资源",child: IconButton(
                          onPressed: () =>
                              showAvailableTorrents(widget.seriesId, k, 0),
                          icon: const Icon(Icons.manage_search)),)
                    ],
                  ))
                ], rows: m[k]!),
              ],
            );
            list.add(seasonList);
          }
          return ListView(
            children: [
              DetailCard(details: details),
              Column(
                children: list,
              ),
            ],
          );
        },
        error: (err, trace) {
          return Text("$err");
        },
        loading: () => const MyProgressIndicator());
  }

  Future<void> showAvailableTorrents(String id, int season, int episode) {
    return showDialog<void>(
      context: context,
      barrierDismissible: true,
      builder: (BuildContext context) {
        return Consumer(builder: (context, ref, _) {
          final torrents = ref.watch(mediaTorrentsDataProvider(
              (mediaId: id, seasonNumber: season, episodeNumber: episode)));

          return AlertDialog(
              //title: Text("资源"),
              content: SelectionArea(
            child: SizedBox(
              width: MediaQuery.of(context).size.width * 0.7,
              height: MediaQuery.of(context).size.height * 0.6,
              child: torrents.when(
                  data: (v) {
                    bool hasPrivate = false;
                    for (final item in v) {
                      if (item.isPrivate == true) {
                        hasPrivate = true;
                      }
                    }
                    final columns = [
                      const DataColumn(label: Text("名称")),
                      const DataColumn(label: Text("大小")),
                      const DataColumn(label: Text("S/P")),
                      const DataColumn(label: Text("来源")),
                    ];
                    if (hasPrivate) {
                      columns.add(const DataColumn(label: Text("消耗")));
                    }
                    columns.add(const DataColumn(label: Text("下载")));

                    return SingleChildScrollView(
                        child: DataTable(
                            dataTextStyle: const TextStyle(fontSize: 12),
                            columns: columns,
                            rows: List.generate(v.length, (i) {
                              final torrent = v[i];
                              final rows = [
                                DataCell(Text("${torrent.name}")),
                                DataCell(Text(
                                    "${torrent.size?.readableFileSize()}")),
                                DataCell(Text(
                                    "${torrent.seeders}/${torrent.peers}")),
                                DataCell(Text(torrent.source ?? "-")),
                              ];
                              if (hasPrivate) {
                                rows.add(DataCell(Text(torrent.isPrivate == true
                                    ? "${torrent.downloadFactor}dl/${torrent.uploadFactor}up"
                                    : "-")));
                              }

                              rows.add(DataCell(IconButton(
                                icon: const Icon(Icons.download),
                                onPressed: () async {
                                  var f = ref
                                      .read(mediaTorrentsDataProvider((
                                        mediaId: id,
                                        seasonNumber: season,
                                        episodeNumber: episode
                                      )).notifier)
                                      .download(torrent)
                                      .then((v) =>
                                          showSnakeBar("开始下载：${torrent.name}"));
                                  showLoadingWithFuture(f);
                                },
                              )));
                              return DataRow(cells: rows);
                            })));
                  },
                  error: (err, trace) {
                    return "$err".contains("no resource found")
                        ? const Center(
                            child: Text("没有资源"),
                          )
                        : Text("$err");
                  },
                  loading: () => const MyProgressIndicator()),
            ),
          ));
        });
      },
    );
  }
}
