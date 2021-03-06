package main

import (
	"log"

	"github.com/jakobii/dcraft/internal/app"
	"github.com/jakobii/dcraft/internal/app/cfg"
)

func main() {
	opts, err := cfg.Get()
	if err != nil {
		log.Fatal(err)
	}
	app.Start(opts...)
}
