# MQTT Go client example

This example demonstrates how an application establishes a connection to a mosquitto broker.

Spin up the stack:

```bash
docker compose up
```

Expected output:

```bash
[+] up 1/1
 ✔ Image mqtt-app Built                                                                                                                     0.5s
Attaching to mosquitto, app-1
mosquitto  | 1774931926: mosquitto version 2.0.22 starting
mosquitto  | 1774931926: Config loaded from /mosquitto/config/mosquitto.conf.
mosquitto  | 1774931926: Opening ipv4 listen socket on port 1883.
mosquitto  | 1774931926: mosquitto version 2.0.22 running
mosquitto  | 1774931926: New connection from 172.19.0.3:50658 on port 1883.
mosquitto  | 1774931926: New client connected from 172.19.0.3:50658 as my-client (p2, c1, k30).
app-1      | 2026-03-31 04:38:46.406746 +0000 UTC - Client connected
app-1      | [topic]: my-topic
app-1      | [topic]: hello world
```

Tear down the stack:

```bash
docker compose down -v -t 1
```
