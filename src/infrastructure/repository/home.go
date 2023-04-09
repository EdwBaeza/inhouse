package repository

import (
	"github.com/edwbaeza/inhouse/src/domain/entity"
)

type HomeMemoryRepository struct {
	homes map[string]entity.Home
}

func NewHomeMemoryRepository() *HomeMemoryRepository {
	return &HomeMemoryRepository{
		homes: make(map[string]entity.Home),
	}
}

func (r *HomeMemoryRepository) FindHomeById(id string) (*entity.Home, error) {
	home, ok := r.homes[id]
	if !ok {
		return nil, nil
	}
	return &home, nil
}

func (r *HomeMemoryRepository) CreateHome(home *entity.Home) error {
	r.homes[home.Id] = *home
	return nil
}

func (r *HomeMemoryRepository) ListHomes() ([]entity.Home, error) {
	homes := make([]entity.Home, 0)
	for _, home := range r.homes {
		homes = append(homes, home)
	}
	return homes, nil
}
