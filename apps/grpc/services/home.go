package services

import (
	"context"
	"log"

	"github.com/edwbaeza/inhouse/apps/grpc/protos/homepb"
	"github.com/edwbaeza/inhouse/src/domain/entity"
)

type HomeService struct {
	repository entity.HomeRepository
	homepb.UnimplementedHomeServiceServer
}

func NewHomeService(repository entity.HomeRepository) *HomeService {
	return &HomeService{repository: repository}
}

func (s *HomeService) GetHome(ctx context.Context, request *homepb.GetHomeRequest) (*homepb.Home, error) {
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

func (s *HomeService) SaveHome(ctx context.Context, home *homepb.Home) (*homepb.HomeResponse, error) {
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
