package grpc

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/edwbaeza/inhouse/apps/grpc/protos/homepb"
	serversGrpc "github.com/edwbaeza/inhouse/apps/grpc/server"
	"github.com/edwbaeza/inhouse/src/infrastructure/repository"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestGetHome(t *testing.T) {
	client, closer := serverTest(context.Background())
	defer closer()

	// TODO create home
	_, err := client.GetHome(context.Background(), &homepb.GetHomeRequest{
		Id: "1",
	})

	if err != nil {
		t.Fatalf("Failed to get home: %s", err.Error())
	}
}

func serverTest(ctx context.Context) (homepb.HomeServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	repository := repository.NewHomeMemoryRepository()
	homepb.RegisterHomeServiceServer(baseServer, serversGrpc.NewHomeServer(repository))
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := homepb.NewHomeServiceClient(conn)

	return client, closer
}
