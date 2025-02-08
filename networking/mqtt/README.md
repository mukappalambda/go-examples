# MQTT Go client example

First of all, make sure to have [mosquitto](https://github.com/eclipse/mosquitto) installed on the machine. As of writing this note, the latest version of `mosquitto` is [`v2.0.20`](https://github.com/eclipse-mosquitto/mosquitto/releases/tag/v2.0.20).

Install `mosquitto`:

```bash
wget https://mosquitto.org/files/source/mosquitto-2.0.20.tar.gz
tar zxvf mosquitto-2.0.20.tar.gz
cd mosquitto-2.0.20
mkdir build
cd build
cmake ..
make -j$(nproc)
sudo make install
```

```console
$ mosquitto -h
mosquitto version 2.0.20

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

To run this example, start the MQTT broker by running `mosquitto` in the terminal:

```console
$ mosquitto
1730165930: mosquitto version 2.0.20 starting
1730165930: Using default config.
1730165930: Starting in local only mode. Connections will only be possible from clients running on this machine.
1730165930: Create a configuration file which defines a listener to allow remote access.
1730165930: For more details see https://mosquitto.org/documentation/authentication-methods/
1730165930: Opening ipv4 listen socket on port 1883.
1730165930: Opening ipv6 listen socket on port 1883.
1730165930: mosquitto version 2.0.20 running
```

Build this example:

```bash
go build
```

Usage:

```bash
./mqtt-workout -h
```

```bash
Usage of ./mqtt-workout:
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

Run `./mqtt-workout`:

```bash
$ ./mqtt-workout
2024-05-07 08:49:23.243248 +0800 CST - Client connected
2024-05-07 08:49:23.243498 +0800 CST - Topic: my-topic; payload: hello world%
```
