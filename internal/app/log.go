package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) log(v ...interface{}) {
	log.Println(v...)
}

func (s *Server) logRequest(r *http.Request) {
	l := newLogHTTP(r)
	b, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		s.log(err)
	}
	s.log(string(b))
}

type logHTTP struct {
	Address string     `json:"address"`
	Request logRequest `json:"request"`
}

type logRequest struct {
	ID      uuid.UUID         `json:"id"`
	Method  string            `json:"method"`
	URI     string            `json:"uri"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func newLogHTTP(r *http.Request) logHTTP {
	header := make(map[string]string)
	for k, v := range r.Header {
		header[k] = v[0]
	}
	var body string
	if r.Body != nil {
		body = readRequestBodyString(r)
	}
	return logHTTP{
		Address: r.RemoteAddr,
		Request: logRequest{
			ID:      uuid.New(),
			Method:  r.Method,
			URI:     r.RequestURI,
			Headers: header,
			Body:    body,
		},
	}
}
