package main

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
}
