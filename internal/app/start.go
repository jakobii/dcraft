package app

import (
	"log"
	"net/http"
	"strconv"
)

// Start is the main entry point to the app
func Start(opts ...ServerOption) {
	s := NewServer(opts...)
	log.Println("starting server on port", s.port)
	http.ListenAndServe(
		":"+strconv.FormatUint(s.port, 10),
		s.Handler(),
	)
}
