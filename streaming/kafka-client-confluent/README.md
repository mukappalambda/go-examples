# kafka client confluent

```bash
docker run -dt -p 19092:9092 --name broker apache/kafka:4.1.0

go build
./kafka-client-confluent -brokers localhost:19092

docker rm -f $(docker ps -aq)
```
