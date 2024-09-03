# MQTT Go client example

First of all, make sure to have [mosquitto](https://github.com/eclipse/mosquitto) installed on the machine.

```bash
$ mosquitto -h
mosquitto version 2.0.18

mosquitto is an MQTT v5.0/v3.1.1/v3.1 broker.

Usage: mosquitto [-c config_file] [-d] [-h] [-p port]

 -c : specify the broker config file.
 -d : put the broker into the background after starting.
 -h : display this help.
 -p : start the broker listening on the specified port.
      Not recommended in conjunction with the -c option.
 -v : verbose mode - enable all logging types. This overrides
      any logging options given in the config file.

See https://mosquitto.org/ for more information.

```

To run this example, start the MQTT broker:

```bash
mosquitto
```

Usage of `main.go`:

```bash
  -clientId string
        client id (default "my-client")
  -host string
        host (default "localhost")
  -payload string
        payload (default "hello world")
  -port string
        port (default "1883")
  -scheme string
        scheme (default "tcp")
  -topic string
        topic name (default "my-topic")
```

Run `go run main.go`:

```bash
$ go run main.go
2024-05-07 08:49:23.243248 +0800 CST - Client connected
2024-05-07 08:49:23.243498 +0800 CST - Topic: my-topic; payload: hello world%
```
