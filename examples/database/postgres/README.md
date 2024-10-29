# Usage of database/sql

Create a postgres docker container.

```console
$ docker run -dt \
--name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=my_db \
-p 5432:5432 \
postgres:14
9aa033b1af865c08df37b5e8bf9dc71c1cbf0477165deca05a6c880af2a9b940
$ docker ps --filter name=postgres
CONTAINER ID   IMAGE         COMMAND                  CREATED         STATUS         PORTS                    NAMES
9aa033b1af86   postgres:14   "docker-entrypoint.s…"   4 seconds ago   Up 3 seconds   0.0.0.0:5432->5432/tcp   postgres
```

Find and import the [Go-PostgreSQL Driver](https://github.com/lib/pq)

```bash
go get github.com/lib/pq
```

```console
$ go run ping-database/main.go
Database connected.
max open connections: 50
```

Create a `todo` table

```console
$ docker exec -it postgres \
psql -U postgres -d my_db \
-c "CREATE TABLE todo (id serial PRIMARY KEY, name text, is_completed bool);"
CREATE TABLE
$ docker exec -it postgres \
psql -U postgres -d my_db \
-c "\dt"
        List of relations
 Schema | Name | Type  |  Owner
--------+------+-------+----------
 public | todo | table | postgres
(1 row)
```

```console
$ go run create-todos/main.go
{Buy groceries false}
{Do laundry false}
{Finish project true}
Todo records inserted successfully.
```

```console
$ docker exec -it postgres \
psql -U postgres -d my_db \
-c "SELECT * from todo;"
 id |      name      | is_completed
----+----------------+--------------
  1 | Buy groceries  | f
  2 | Do laundry     | f
  3 | Finish project | t
(3 rows)
```

Tear down the postgres container

```console
$ docker rm -f postgres
postgres
```
