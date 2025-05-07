FROM golang:1.24.2-alpine AS build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /gin-demo-server -ldflags="-s -w" cmd/gin-demo-server/main.go

FROM scratch

COPY --from=build /gin-demo-server /gin-demo-server

EXPOSE 8080

ENTRYPOINT ["/gin-demo-server"]
