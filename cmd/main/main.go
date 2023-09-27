package main

import (
	"log"

	"github.com/popooq/rodnoolee_birthday/internal/bootstrap"
)

func main() {
	err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}
}
