package app

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

// Start is the main entry point to the app
func Start() {
	port := flag.Uint64("port", 80, "http port")
	flag.Parse()

	log.Println("starting server on port", *port)

	s := Server{}
	http.ListenAndServe(
		":"+strconv.FormatUint(*port, 10),
		s.Handler(),
	)
}
