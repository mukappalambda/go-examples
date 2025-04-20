FROM golang:1.24.2 AS build

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -o /build/my-app -ldflags="-s -w" .

FROM scratch

COPY --from=build /build/my-app /my-app

ENTRYPOINT ["/my-app"]

