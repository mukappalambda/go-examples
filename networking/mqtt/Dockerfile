FROM golang:1.24.4 AS build

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -o /build/app -ldflags="-s -w" .

FROM scratch

COPY --from=build /build/app /app

ENTRYPOINT ["/app"]

CMD ["-host", "mosquitto"]
