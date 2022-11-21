package internal

import "gopkg.in/guregu/null.v4"

type Home struct {
	id          null.String
	title       string
	description string
	rawAddress  string
	features    map[string]string
	price       float64
	currency    null.String
}

type HomeRepository interface {
	Save(home Home) error
	Find(id string) Home
}

func (home *Home) GetId() *string {
	return home.id.Ptr()
}

func (home *Home) SetId(id string) {
	home.id.SetValid(id)
}

func (home *Home) GetTitle() string {

	return home.title
}

func (home *Home) SetTitle(title string) {
	home.title = title

}

func (home *Home) GetDescription() string {

	return home.description
}

func (home *Home) SetDescription(description string) {
	home.description = description
}

func (home *Home) GetRawAddress() string {

	return home.rawAddress
}

func (home *Home) SetRawAddress(rawAddress string) {
	home.rawAddress = rawAddress
}

func (home *Home) GetFeatures() map[string]string {

	return home.features
}

func (home *Home) SetFeatures(features map[string]string) {
	home.features = features
}

func (home *Home) GetPrice() float64 {

	return home.price
}

func (home *Home) SetPrice(price float64) {
	home.price = price
}

func (home *Home) GetCurrency() *string {
	return home.currency.Ptr()
}

func (home *Home) SetCurrency(currency string) {
	home.currency.SetValid(currency)
}
