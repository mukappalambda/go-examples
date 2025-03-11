# Set up a web server with HTTPS

Generate the server key and server certificate:

```bash
bash run.sh
```

Build:

```bash
make build
```

Start the server:

```bash
./bin/https-server
```

Make a client request:

```bash
curl --insecure https://localhost:4443
```

Clean the binary:

```bash
make clean
```
