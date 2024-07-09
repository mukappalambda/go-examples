# Gin Demo

To start up the server:

```bash
make run
```

To add all dependencies:

```bash
go get .
```

To update all dependencies:

```bash
go get -u
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
