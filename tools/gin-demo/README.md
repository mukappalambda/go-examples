# Gin Demo

To start up the server locally:

```bash
make run
```

To build the docker image:

```bash
make build-image
```

To start up the docker container:

```bash
make up-container
```

To tear down the docker container:

```bash
make down-container
```

To add all dependencies:

```bash
go mod tidy
```

To update all dependencies:

```bash
go get -u ./...
go mod tidy
```

To test the endpoints:

```bash
# get all books
curl --include \
http://localhost:8080/books/

# add a new book
curl --include \
-X POST \
http://localhost:8080/books/ \
--data '{"id": 3, "author": "mark", "title": "mark'\'' book"}'
```
