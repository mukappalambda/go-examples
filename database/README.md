# Usage of database/sql

Create a postgres docker container.

```bash
docker run -dt \
--name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=my_db \
-p 5432:5432 \
postgres:14

docker ps --filter name=postgres
```

Find and import the [Go-PostgreSQL Driver](https://github.com/lib/pq)

```bash
go get github.com/lib/pq
```

```bash
go run ping_database.go
#Database connected.
#max open connections: 50
```

Create a `todo` table

```bash
docker exec -it postgres \
psql -U postgres -d my_db \
-c "CREATE TABLE todo (id serial PRIMARY KEY, name text, is_completed bool);"
#CREATE TABLE

docker exec -it postgres \
psql -U postgres -d my_db \
-c "\dt"
#        List of relations
# Schema | Name | Type  |  Owner
#--------+------+-------+----------
# public | todo | table | postgres
#(1 row)
```

```bash
go run create_todos.go
#{Buy groceries false}
#{Do laundry false}
#{Finish project true}
#Todo records inserted successfully.
```

```bash
docker exec -it postgres \
psql -U postgres -d my_db \
-c "SELECT * from todo;"
# id |      name      | is_completed
#----+----------------+--------------
#  1 | Buy groceries  | f
#  2 | Do laundry     | f
#  3 | Finish project | t
#(3 rows)
```

Tear down the postgres container

```bash
docker rm -f postgres
```
