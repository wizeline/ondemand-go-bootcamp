package controller

import (
	"github.com/go-chi/chi/v5"
)

// HTTP controller for HTTP protocol.
type HTTP interface {
	SetRoutes(r chi.Router)
}

// BasicMessage is the representation of a basic http JSON response.
type BasicMessage struct {
	Message string `json:"message"`
}
