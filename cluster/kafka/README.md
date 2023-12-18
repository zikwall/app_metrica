### How work with Apache Kafka in Docker

```shell
$ docker exec -t 0fc454d70be2 /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server kafka_kafka-1_1:9092 --topic appmetrica_main_bus --describe
$ docker exec -t 0fc454d70be2 /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server kafka_kafka-1_1:9092 --topic appmetrica_main_bus --delete
$ docker exec -it 0fc454d70be2 /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server kafka_kafka-1_1:9092 --topic appmetrica_main_bus --create --replication-factor 1 --partitions 6
```

### OLD

- [x] Create new topic

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --create \
    --topic ClickhouseTopic \
    --partitions 3 \
    --replication-factor 1
```

- [x] Print out the topics

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --list
```

- [x] Describe

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --describe \
    --topic MyTopic1
```

- [x] Run Kafka console consumer (run in another console)

```shell script
docker exec -t clickhouse-kafka \
  kafka-console-consumer.sh \
    --bootstrap-server :9092 \
    --group my-group \
    --topic MyTopic1
```

- [x] Run Kafka console producer

```shell script
docker exec -it clickhouse-kafka \
  kafka-console-producer.sh \
    --broker-list :9092 \
    --topic MyTopic1
```

- [x] Get count messages of topic

```shell script
docker exec -it clickhouse-kafka \
  kafka-run-class.sh \
  kafka.tools.GetOffsetShell \
    --broker-list :9092 \
    --topic MyTopic1 \
    --time -1
```

- [x] Drop topic

```shell script
docker exec -it clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :2181 \
    --topic MyTopic1 \
    --delete
```

**after put messages in producer console && to see messages printed out in second terminal, where run Kafka CLI consumer**

- [x] Show full logs fro Kafka run: `$ docker logs -f clickhouse-kafka`
- [x] [Kafkacat](https://github.com/edenhill/kafkacat)
- [x] [Kafdrop](https://github.com/obsidiandynamics/kafdrop)