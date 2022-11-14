FROM golang:1.19.0 as base

FROM base as dev

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/cosmtrek/air@latest

COPY . .
# RUN go build -v -o /usr/local/bin/app ./cmd/api/

EXPOSE 3000
CMD air --build.cmd "go build -o /usr/local/bin/app ./cmd/api/main.go" --build.bin "/usr/local/bin/app"
