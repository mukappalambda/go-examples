FROM golang:1.24.2-alpine AS build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /gin-demo -ldflags="-s -w" ./main.go

FROM scratch

COPY --from=build /gin-demo /gin-demo

EXPOSE 8080

ENTRYPOINT ["/gin-demo"]
