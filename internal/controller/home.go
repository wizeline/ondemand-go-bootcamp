package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	_ HTTP     = &homeHTTP{}
	_ HomeHTTP = &homeHTTP{}
)

type HomeHTTP interface {
	SetRoutes(r *mux.Router)
	Home(w http.ResponseWriter, r *http.Request)
}

type homeHTTP struct{}

func NewHomeHTTP() HomeHTTP {
	return &homeHTTP{}
}

func (h homeHTTP) SetRoutes(r *mux.Router) {
	r.HandleFunc("/", h.Home).Methods("GET")
}

func (homeHTTP) Home(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to Capstone API \nGO Bootcamp - Marcos Jauregui")
	w.WriteHeader(http.StatusOK)
}
