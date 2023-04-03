FROM golang:1.19.7 as builder

WORKDIR /app

COPY main.go go.mod go.sum ./

RUN go build -o arcadia

FROM docker:latest

COPY --from=builder /app/arcadia /usr/local/bin/arcadia

RUN ls -l /usr/local/bin/

CMD ["/usr/local/bin/arcadia"]
