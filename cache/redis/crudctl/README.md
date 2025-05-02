# crudctl with Redis

Build the binary:

```bash
make build
```

Start up the Redis instance:

```bash
docker run -dt --name redis -p 6379:6379 redis:8.0-rc1-alpine
```

Usage of `crudctl`:

```bash
./bin/crudctl -h
```

Create:

```bash
./bin/crudctl c -k my.key -v my.value
```

Read:

```bash
./bin/crudctl r -k my.key
```

Update:

```bash
./bin/crudctl u -k my.key -v your.value
```

Delete:

```bash
./bin/crudctl d -k my.key
```

Tear down the Redis container:

```bash
docker rm -f redis
```
