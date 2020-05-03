# Computer Vision Service

To generate code according to proto file, install software according to https://github.com/golang/protobuf and run
```bash
protoc -I ./ ./cvservice.proto --go_out=plugins=grpc:.
```