package server

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/edwbaeza/inhouse/apps/grpc/protos/homepb"
	"github.com/edwbaeza/inhouse/src/domain/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HomeServer struct {
	repository entity.HomeRepository
	homepb.UnimplementedHomeServiceServer
}

func NewHomeServer(repository entity.HomeRepository) *HomeServer {
	return &HomeServer{repository: repository}
}

func (s *HomeServer) GetHome(ctx context.Context, request *homepb.GetHomeRequest) (*homepb.Home, error) {
	home, err := s.repository.FindHomeById(request.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if home == nil {
		return &homepb.Home{}, nil
	}
	return &homepb.Home{
		Id:         home.Id,
		Name:       home.Name,
		RawAddress: home.RawAddress,
	}, nil
}

func (s *HomeServer) SaveHome(ctx context.Context, home *homepb.Home) (*homepb.HomeResponse, error) {
	err := s.repository.CreateHome(&entity.Home{
		Id:         home.GetId(),
		Name:       home.GetName(),
		RawAddress: home.GetRawAddress(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &homepb.HomeResponse{
		Id: home.GetId(),
	}, nil
}

func (s *HomeServer) ListHomes(empty *homepb.Empty, stream homepb.HomeService_ListHomesServer) error {
	homes, err := s.repository.ListHomes()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, home := range homes {
		err := stream.Send(&homepb.Home{
			Id:         home.Id,
			Name:       home.Name,
			RawAddress: home.RawAddress,
		})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(10 * time.Second)
	}

	return nil
}

func (s *HomeServer) SaveHomes(stream homepb.HomeService_SaveHomesServer) error {
	for {
		home, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Println(err)
			return status.Errorf(codes.Unknown, "Failed to receive a home: %v", err)
		}

		err = s.repository.CreateHome(&entity.Home{
			Id:         home.GetId(),
			Name:       home.GetName(),
			RawAddress: home.GetRawAddress(),
		})

		if err != nil {
			log.Println(err)
			return status.Errorf(codes.Unknown, "Failed to save a home: %v", err)
		}
		err = stream.Send(&homepb.HomeResponse{
			Id: home.GetId(),
		})

		if err != nil {
			log.Println(err)
			return status.Errorf(codes.Unknown, "Failed to send a home: %v", err)
		}
	}
}
