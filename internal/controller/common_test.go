package controller

import (
	"errors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var testSvcError = errors.New("some service error")

func newTestRouter(ctrl HTTP) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ctrl.SetRoutes(r)
	return r
}
