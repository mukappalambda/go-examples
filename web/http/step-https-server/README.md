# HTTPS Server with step

```bash
mise search step
mise ls-remote step
mise use step -g

step version
```

Generating the Root CA certificate and key:

```bash
step certificate create root-ca root_ca.crt root_ca.key --profile root-ca --insecure --no-password
```

Expected output:

```bash
Your certificate has been saved in root_ca.crt.
Your private key has been saved in root_ca.key.
```

Generating the server certificate and key:

```bash
step certificate create my-server server.crt server.key --ca root_ca.crt --ca-key root_ca.key --san 127.0.0.1 --san localhost --san myserver.com --insecure --no-password
```

Expected output:

```bash
Your certificate has been saved in server.crt.
Your private key has been saved in server.key.
```

Starting the server:

```bash
go run server/main.go
```

Connecting to the server via curl:

```bash
curl --cacert ./root_ca.crt https://127.0.0.1:9443/data

curl --cacert ./root_ca.crt https://localhost:9443/data
```

Connecting to the server programatically:

```bash
go run client/main.go
```
