package Home

import (
	"github.com/EdwBaeza/inhouse/internal"
)

type Home struct {
	ID          *string           `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	RawAddress  string            `json:"raw_address"`
	Features    map[string]string `json:"features"`
	Price       float64           `json:"price"`
	Currency    *string           `json:"currency"`
}

type HomeRequest struct {
	Home Home `json:"home"`
}

func FromHomeToReq(home internal.Home) HomeRequest {
	return HomeRequest{
		Home: Home{
			ID:          home.GetId(),
			Title:       home.GetTitle(),
			Description: home.GetDescription(),
			RawAddress:  home.GetRawAddress(),
			Features:    home.GetFeatures(),
			Price:       home.GetPrice(),
			Currency:    home.GetCurrency(),
		},
	}
}
