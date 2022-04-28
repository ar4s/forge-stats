# forge-stats

forge-stats is a simple and self-contained binary.
The goal of this project is to save the statistics of your mod or modpack from CurseForge to a time series database (InfluxDB).

## Getting started
1. InfluxDB - [try a free plan](https://cloud2.influxdata.com/signup) or [install](https://docs.influxdata.com/influxdb/v2.2/get-started/).
1. Download latest version of [forge-stats](https://github.com/ar4s/forge-stats/releases).
1. Run periodically `forge-stats` by `cron` or `systemd`.
1. Create dasboard or use [template](./contrib/influx-board-template.json)

## Configuration
You can configure the app by passing argument to command line or by set `[$ENVIRONMENT_VARIABLE]`.
```
   --influx-url value      [$FORGE_STATS_INFLUX_URL]
   --influx-org value     (default: "default") [$FORGE_STATS_INFLUX_ORG]
   --influx-bucket value  (default: "default") [$FORGE_STATS_INFLUX_BUCKET]
   --status-url value     (default: "https://addons-ecs.forgesvc.net/api/v2/addon") [$FORGE_STATS_STATUS_URL]
   --mod-id value         (default: 0) [$FORGE_STATS_MOD_ID]
   --token value           [$FORGE_STATS_INFLUX_TOKEN]
   --stat-name value      (default: "default") [$FORGE_STAT_NAME]
   --help, -h             show help (default: false)
   ```
