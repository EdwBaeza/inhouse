package server

import (
	"context"
	"log"
	"time"

	"github.com/edwbaeza/inhouse/apps/grpc/protos/homepb"
	"github.com/edwbaeza/inhouse/src/domain/entity"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *HomeServer) ListHomes(empty *emptypb.Empty, stream homepb.HomeService_ListHomesServer) error {
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
