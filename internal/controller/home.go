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

// HomeHTTP set the routes and handler functions related to Home
type HomeHTTP interface {
	SetRoutes(r *mux.Router)
	Home(w http.ResponseWriter, r *http.Request)
}

type homeHTTP struct{}

// NewHomeHTTP returns a new HomeHTTP implementation.
func NewHomeHTTP() HomeHTTP {
	return &homeHTTP{}
}

func (h homeHTTP) SetRoutes(r *mux.Router) {
	r.HandleFunc("/", h.Home).Methods("GET")
}

// Home is a handler function that responses the home page in HTML format.
func (homeHTTP) Home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprintf(w, `
		<H1>Welcome to the Capstone API</H1>
		<H2>GO Bootcamp - Academy </H2>
		<B>Gopher:</B> Marcos Jauregui <BR>
		<B>Email:</B> marcos.jauregui@wizeline.com
	`)
	w.WriteHeader(http.StatusOK)
}
