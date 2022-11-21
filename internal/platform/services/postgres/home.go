package postgres

import (
	mooc "github.com/EdwBaeza/inhouse/internal"
)

// Mocked postgres repository, TODO add conecction to postgres

type HomeRepository struct {
}

func NewHomeRepository() *HomeRepository {
	return &HomeRepository{}
}

func (h *HomeRepository) Find(id string) mooc.Home {
	home := mooc.Home{}
	home.SetId(id)
	home.SetDescription("mock")
	home.SetTitle("mock title")
	home.SetFeatures(map[string]string{"test": "testing...."})

	return home
}

func (h *HomeRepository) Save(home mooc.Home) error {

	return nil
}
