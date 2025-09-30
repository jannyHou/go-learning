# Overview

- Generate grpc proto files

```bash
protoc --go_out=./order_proto --go_opt=paths=source_relative \
--go-grpc_out=./order_proto --go-grpc_opt=paths=source_relative \
order.proto
```

- Run server and client

```bash
go run server/main.go
# Another terminal
go run client/main.go
```