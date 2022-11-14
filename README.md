## README Inhouse

#### Build

```console
docker-compose -f local.yml build
```

#### Run Server

```console
docker-compose -f local.yml up
```

#### Run Tests

```console
docker-compose -f local.yml run --rm app go test --cover -v ./...
```
