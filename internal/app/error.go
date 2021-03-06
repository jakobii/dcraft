package app

import (
	"net/http"
)

// Err handles errors
func (s *Server) Err(w http.ResponseWriter, r *http.Request, status int, err error) {
	s.logRequest(r)
	WriteJSON(w, status, NewErrorResponse(r, err))
}

// ErrInternal handles internal server error
func (s *Server) ErrInternal(w http.ResponseWriter, r *http.Request, err error) {
	s.Err(w, r, http.StatusInternalServerError, err)
}

// ErrorResponse is the standard http error response
type ErrorResponse struct {
	Error   string     `json:"error"`
	Request logRequest `json:"request"`
}

// NewErrorResponse ..
func NewErrorResponse(r *http.Request, err error) ErrorResponse {
	return ErrorResponse{
		Error:   err.Error(),
		Request: newLogRequest(r),
	}
}
