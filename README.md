# inhouse


### Build proto

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative apps/grpc/protos/homepb/home.proto
```

### Dev mode by docker

#### Build
```bash
docker build -t inhouse:dev -f Dockerfile.dev .
```

#### Run server
```bash
docker run --rm -it -p 8081:8080 -e PORT=8080 inhouse:dev
```

#### Run tests

```bash
docker run -e PORT=8080 inhouse:dev go test -v ./...
```