# go-gRPC

This is a simple project focused on learning and improving with gRPC. It consists of a basic entity, Category, which is stored in an SQLite file database. Additionally, by using gRPC, it is possible to send operations to the gRPC server for saving and listing.

## Genereate proto files

```bash
protoc --go_out=. --go-grpc_out=. proto/category.proto
```

## Evans Client Execution

```bash
evans -r repl
package pb
service CategoryService
call {function}
```

## Useful Links

1. [gRPC](https://grpc.io/)
1. [Proto Buf](https://protobuf.dev/)
1. [Evans Client](https://github.com/ktr0731/evans)
