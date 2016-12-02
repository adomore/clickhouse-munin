package main

type MetricConfig struct {
    id string
    label string
    metric_type string
    draw_type string
    color string
    clickhouseEvent string
}

type Widget struct {
    graph_title string
    graph_category string
    graph_info string
    graph_vlabel string
    graph_period string
    graph_args string
    is_percent bool
    metrics []MetricConfig
}

var AvailableWidgets = map[string]Widget{
    "queries": Widget{
        graph_title: "ClickHouse queries",
        graph_category: "clickhouse",
        graph_info: "Values received from ClickHouse system.events table",
        graph_vlabel: "queries / second",
        graph_period: "second",
        graph_args: "--lower-limit 0",
        metrics: []MetricConfig{
            MetricConfig{
                id: "select",
                label: "Selects",
                metric_type: "DERIVE",
                draw_type: "AREA",
                color: "COLOUR0",
                clickhouseEvent: "SelectQuery",
            },
            MetricConfig{
                id: "insert",
                label: "Inserts",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR1",
                clickhouseEvent: "InsertQuery",
            },
        },
    },
    "files": Widget{
        graph_title: "ClickHouse files",
        graph_category: "clickhouse",
        graph_info: "Values received from ClickHouse system.events table",
        graph_vlabel: "operations / second",
        graph_period: "second",
        graph_args: "--lower-limit 0",
        metrics: []MetricConfig{
            MetricConfig{
                id: "file_open",
                label: "Opens",
                metric_type: "DERIVE",
                draw_type: "AREA",
                color: "COLOUR0",
                clickhouseEvent: "FileOpen",
            },
            MetricConfig{
                id: "seek",
                label: "Seeks",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR1",
                clickhouseEvent: "Seek",
            },
        },
    },
    "cache": Widget{
        graph_title: "ClickHouse cache",
        graph_category: "clickhouse",
        graph_info: "Values received from ClickHouse system.events table",
        graph_vlabel: "per second",
        graph_period: "second",
        graph_args: "-u 100 -l 0 -r --base 1000",
        is_percent: true,
        metrics: []MetricConfig{
            MetricConfig{
                id: "hits",
                label: "hit ratio",
                metric_type: "GAUGE",
                draw_type: "AREA",
                color: "COLOUR0",
                clickhouseEvent: "MarkCacheHits",
            },
            MetricConfig{
                id: "misses",
                label: "miss ratio",
                metric_type: "GAUGE",
                draw_type: "STACK",
                color: "COLOUR1",
                clickhouseEvent: "MarkCacheMisses",
            },
        },
    },
    "zookeeper": Widget{
        graph_title: "ClickHouse Zookeeper transactions",
        graph_category: "clickhouse",
        graph_info: "Values received from ClickHouse system.events table",
        graph_vlabel: "operations / second",
        graph_period: "second",
        graph_args: "--lower-limit 0",
        metrics: []MetricConfig{
            MetricConfig{
                id: "ZooKeeperGetChildren",
                label: "ZooKeeperGetChildren",
                metric_type: "DERIVE",
                draw_type: "AREA",
                color: "COLOUR0",
                clickhouseEvent: "ZooKeeperGetChildren",
            },
            MetricConfig{
                id: "ZooKeeperCreate",
                label: "ZooKeeperCreate",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR1",
                clickhouseEvent: "ZooKeeperCreate",
            },
            MetricConfig{
                id: "ZooKeeperRemove",
                label: "ZooKeeperRemove",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR2",
                clickhouseEvent: "ZooKeeperRemove",
            },
            MetricConfig{
                id: "ZooKeeperExists",
                label: "ZooKeeperExists",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR3",
                clickhouseEvent: "ZooKeeperExists",
            },
            MetricConfig{
                id: "ZooKeeperGet",
                label: "ZooKeeperGet",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR4",
                clickhouseEvent: "ZooKeeperGet",
            },
            MetricConfig{
                id: "ZooKeeperSet",
                label: "ZooKeeperSet",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR5",
                clickhouseEvent: "ZooKeeperSet",
            },
            MetricConfig{
                id: "ZooKeeperMulti",
                label: "ZooKeeperMulti",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR6",
                clickhouseEvent: "ZooKeeperMulti",
            },
            MetricConfig{
                id: "ZooKeeperExceptions",
                label: "ZooKeeperExceptions",
                metric_type: "DERIVE",
                draw_type: "STACK",
                color: "COLOUR7",
                clickhouseEvent: "ZooKeeperExceptions",
            },
        },
    },
}
