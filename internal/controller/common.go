package controller

import "github.com/gorilla/mux"

// HTTP controller for HTTP protocol.
type HTTP interface {
	// SetRoutes allocates routes into the given router pointer.
	SetRoutes(router *mux.Router)
}

type BasicMessage struct {
	Message string `json:"message"`
}
