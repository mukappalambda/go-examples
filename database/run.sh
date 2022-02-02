docker network create postgres-net

docker run -dt \
--name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_DB=demo \
--network=postgres-net \
-p 5432:5432 \
postgres:14

docker run -dt --name golang \
--network=postgres-net \
-v $PWD:/data \
-w /data \
golang:1.16.5