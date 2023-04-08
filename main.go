package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/edwbaeza/inhouse/apps/grpc/protos/homepb"
	"github.com/edwbaeza/inhouse/apps/grpc/services"
	"github.com/edwbaeza/inhouse/src/infrastructure/repository"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, tcp := net.Listen("tcp", ":8080")

	if tcp != nil {
		log.Fatalf("Failed to listen: %s", tcp.Error())
	}

	defer listener.Close()
	repository := repository.NewHomeMemoryRepository()
	homeService := services.NewHomeService(repository)
	logrus.SetLevel(logrus.DebugLevel)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				logrus.WithFields(logrus.Fields{
					"method": info.FullMethod,
					"req":    req,
				}).Debug(time.Now(), ": ")
				return handler(ctx, req)
			},
		),
	)
	homepb.RegisterHomeServiceServer(grpcServer, homeService)
	reflection.Register(grpcServer)

	error := grpcServer.Serve(listener)

	if error != nil {
		log.Fatalf("Failed to serve: %s", error.Error())
	}
}
