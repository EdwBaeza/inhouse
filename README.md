# inhouse


### Build proto

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative apps/grpc/protos/homepb/home.proto
```