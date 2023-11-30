# app_metrica

## What is it?

`// TODO`

## Usage

`// TODO`

### UI

- **KafDrop**: http://<host>:9000/
- **UI for Apache Kafka**: http://<host>:8087/

### How work with Clickhouse

- [See manual](/cluster/clickhouse/README.md)

### How work with Kafka

- [See manual](/cluster/kafka/README.md)

### Tests (review)

- [x] First add next line `127.0.0.1    clickhouse-kafka` in `/etc/hosts`, for DEV, why? [1](https://ealebed.github.io/posts/2018/docker-%D1%81%D0%BE%D0%B2%D0%B5%D1%82-28-%D0%BA%D0%B0%D0%BA-%D0%B8%D1%81%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D1%82%D1%8C-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83-connection-reset-by-peer/), [2](https://github.com/grafana/metrictank/issues/1286), [3](https://github.com/wurstmeister/kafka-docker/issues/424)
- [x] Create topic `ClickhouseTopic` if already is not created
- [x] Run consumer in another terminal
- [x] Run app from example/backend `$ go run .`

<details>
  <summary>Output in Go terminal</summary>

  ```shell script
      Send message to broker: user 23, time 2020-08-04 13:58:14
      Send message to broker: user 16, time 2020-08-04 13:58:15
      Send message to broker: user 29, time 2020-08-04 13:58:16
      Send message to broker: user 11, time 2020-08-04 13:58:17
      Send message to broker: user 22, time 2020-08-04 13:58:18
      Send message to broker: user 25, time 2020-08-04 13:58:19
      Send message to broker: user 15, time 2020-08-04 13:58:20
      Send message to broker: user 20, time 2020-08-04 13:58:21
      Send message to broker: user 17, time 2020-08-04 13:58:22
      message at topic/partition/offset MyTopic/0/189:  = {"user_id":23,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:14"}
      message at topic/partition/offset MyTopic/0/190:  = {"user_id":16,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:15"}
      message at topic/partition/offset MyTopic/0/191:  = {"user_id":29,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:16"}
      message at topic/partition/offset MyTopic/0/192:  = {"user_id":11,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:17"}
      message at topic/partition/offset MyTopic/0/193:  = {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:18"}
      message at topic/partition/offset MyTopic/0/194:  = {"user_id":25,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:19"}
      message at topic/partition/offset MyTopic/0/195:  = {"user_id":15,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:20"}
      message at topic/partition/offset MyTopic/0/196:  = {"user_id":20,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:21"}
      message at topic/partition/offset MyTopic/0/197:  = {"user_id":17,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:22"}
      Send message to broker: user 19, time 2020-08-04 13:58:23
      Send message to broker: user 18, time 2020-08-04 13:58:24
      Send message to broker: user 28, time 2020-08-04 13:58:25
  ```
</details>


<details>
  <summary>Output in consumer terminal</summary>

  ```shell script
      {"user_id":19,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:23"}
      {"user_id":18,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:24"}
      {"user_id":28,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:25"}
      {"user_id":14,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:26"}
      {"user_id":13,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:27"}
      {"user_id":14,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:28"}
      {"user_id":17,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:29"}
      {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:30"}
      {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:31"}
      {"user_id":26,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:32"}
      {"user_id":29,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:33"}
      {"user_id":10,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:34"}
      {"user_id":16,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:35"}
      {"user_id":21,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:36"}
      {"user_id":26,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:37"}
  ```
</details>

- [x] Another terminal `make cluster-client ch=01` for connect `ch-01` server
- [x] `SELECT * from main;`

**Output**

```shell script
clickhouse-01 :) select * from main;

┌─user_id─┬─app─┬─host─┬─event─┬─ip─┬─guid─┬──────────created_at─┐
│      10 │     │      │       │    │      │ 2020-08-04 15:48:12 │
│      29 │     │      │       │    │      │ 2020-08-04 15:48:13 │
│      24 │     │      │       │    │      │ 2020-08-04 15:48:14 │
│      14 │     │      │       │    │      │ 2020-08-04 15:48:15 │
│      10 │     │      │       │    │      │ 2020-08-04 15:48:16 │
│      21 │     │      │       │    │      │ 2020-08-04 15:48:17 │
│      11 │     │      │       │    │      │ 2020-08-04 15:48:18 │
│      27 │     │      │       │    │      │ 2020-08-04 15:48:19 │
└─────────┴─────┴──────┴───────┴────┴──────┴─────────────────────┘
```

### Cluster

- [x] Create `main`, `queue` and `mainconsumer` tables each hosts `ch-`: 01, 02, 03, 04, 05
    - you can use `$ bin/create-replica.sh 03`
- [x] Create distributed table on last host `ch-06` from `example/database/distributed.sql`
    - or use `$ bin/create-distributed.sh 06`
- [x] connect to `ch-06`
    - `$ make cluster-client ch=06`
- [x] `SELECT COUNT() FROM main_distributed`;

**Output**

```shell script
clickhouse-06 :) select count() from main_distributed;

SELECT count()
FROM main_distributed

┌─count()─┐
│    1016 │
└─────────┘
```

- [x] Check connect one of server, example `ch-01`, `ch-02` and `SELECT count(*) from main;`

**Output**

```shell script
clickhouse-01 :) select count() from main;

SELECT count()
FROM main

┌─count()─┐
│     945 │
└─────────┘
```

**Output**

```shell script
clickhouse-02 :) select count() from main;

SELECT count()
FROM main

┌─count()─┐
│      71 │
└─────────┘
```

Happy! ^_^

### Following manuals

- [x] [Hard manual](https://github.com/zikwall/clickhouse-docs)