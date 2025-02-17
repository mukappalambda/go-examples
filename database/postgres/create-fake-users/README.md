# Create Fake Users

```bash
docker run -dt --name postgres -p 5432:5432 \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=demo postgres:16

go build
./create-fake-users -num-fake-users 10

go clean

docker rm -f postgres
```
