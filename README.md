# GPRC-Server-With-Go

## Compile the Protobuf file

```bash
protoc \
    --go_out=invoicer \
    --go_opt=paths=source_relative \
    --go-grpc_out=invoicer \
    --go-grpc_opt=paths=source_relative \
    bufs/v1/invoice.proto
```

1. **`--go_out=invoicer`**: Specifies the directory (`invoicer`) where the generated Protobuf message definition files will be saved.
2. **`--go_opt=paths=source_relative`**: Ensures output file paths are relative to the location of the input `.proto` file.
3. **`--go-grpc_out=invoicer`**: Specifies the directory (`invoicer`) where the generated gRPC service stubs will be saved.
4. **`--go-grpc_opt=paths=source_relative`**: Ensures gRPC service stub file paths are relative to the input `.proto` file's location.

