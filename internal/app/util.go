package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func readRequestBody(r *http.Request) []byte {
	if r.Body == nil {
		return []byte{}
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}
	}
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(b))
	return b
}

func readRequestBodyString(r *http.Request) string {
	return string(readRequestBody(r))
}

// WriteJSON marshals the object provided as json and writes it to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		http.Error(w, fmt.Sprintf(`WriteJSON error: %v`, err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(j))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprint(w, string(j))
}
