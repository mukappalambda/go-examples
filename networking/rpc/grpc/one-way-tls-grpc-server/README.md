# One-Way TLS gRPC Server

```bash
mise search step
mise ls-remote step
mise use step -g

step version
```

Generating the Root CA certificate and key:

```bash
step certificate create root-ca root_ca_crt.pem root_ca_key.pem --profile root-ca --insecure --no-password
```

Expected output:

```bash
Your certificate has been saved in root_ca_crt.pem.
Your private key has been saved in root_ca_key.pem.
```

Generating the server certificate and key:

```bash
step certificate create my-server server_crt.pem server_key.pem --ca root_ca_crt.pem --ca-key root_ca_key.pem --san 127.0.0.1 --san localhost --san myserver.com --insecure --no-password
```

Expected output:

```bash
Your certificate has been saved in server_crt.pem.
Your private key has been saved in server_key.pem.
```

Starting the server:

```bash
go run server/main.go
```

Connecting to the server programmatically:

```bash
go run client/main.go -m 'alpha beta gamma'
```
