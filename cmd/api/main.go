package main

import (
	"log"

	"github.com/EdwBaeza/inhouse/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
