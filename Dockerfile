FROM golang:1.20 AS builder

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux

WORKDIR /src
COPY . .

RUN go mod download

RUN go build -o /bin/service ./main.go
RUN upx -q -9 /bin/service

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/service /src/service
COPY --from=builder /src /src


WORKDIR /src

ENV PORT 8080

ENTRYPOINT ["/src/service"]