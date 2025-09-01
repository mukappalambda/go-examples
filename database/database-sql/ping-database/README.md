# Ping database

```console
$ go build
$ docker run -dt --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=demo postgres:17
b60bc7cbe8a81e93d88bbfe6bd56bb304fdcf9fb5d2b500f50347f2475f492fc
$ docker ps --filter name=postgres
CONTAINER ID   IMAGE         COMMAND                  CREATED         STATUS         PORTS                                         NAMES
b60bc7cbe8a8   postgres:17   "docker-entrypoint.s…"   2 seconds ago   Up 2 seconds   0.0.0.0:5432->5432/tcp, [::]:5432->5432/tcp   postgres
$ DATABASE_DSN="postgres://postgres:password@localhost:5432/demo?sslmode=disable" ./ping-database
ping database successfully.
$ docker rm -f $(docker ps --filter name=postgres -aq)
b60bc7cbe8a8
$ go clean
```
