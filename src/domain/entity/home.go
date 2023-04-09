package entity

type Home struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	RawAddress string `json:"rawAddress,omitempty"`
}

type HomeRepository interface {
	FindHomeById(id string) (*Home, error)
	CreateHome(home *Home) error
	ListHomes() ([]Home, error)
}
