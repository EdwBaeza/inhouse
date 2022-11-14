package internal

type Home struct {
	id          string
	title       string
	description string
	raw_address string
	features    map[string]string
	price       float64
	currency    string
}

type HomeRepository interface {
}
