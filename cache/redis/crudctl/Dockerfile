FROM golang:1.24.2 AS build

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/crudctl -ldflags="-s -w" .

FROM golang:1.24.2-alpine

COPY --from=build /bin/crudctl /bin/crudctl

ENTRYPOINT ["/bin/sh"]
