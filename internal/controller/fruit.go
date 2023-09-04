package controller

import (
	"net/http"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var (
	_ HTTP  = &fruit{}
	_ Fruit = &fruit{}
)

// Fruit set the routes and handler functions related to the Fruit controller
type Fruit interface {
	SetRoutes(r chi.Router)
	GetFruit(w http.ResponseWriter, r *http.Request)
	GetFruits(w http.ResponseWriter, r *http.Request)
}

// FruitSvc is the abstraction of the Fruit service dependency.
type FruitSvc interface {
	Get(filter, value string) (entity.Fruits, error)
	GetAll() (entity.Fruits, error)
}

type fruit struct {
	svc FruitSvc
}

// NewFruit returns a new Fruit implementation.
func NewFruit(svc FruitSvc) Fruit {
	return &fruit{
		svc: svc,
	}
}

// SetRoutes sets a fresh middleware stack for the Fruit controller's handle functions and mounts them to the provided sub router.
func (f fruit) SetRoutes(r chi.Router) {
	r.Get("/fruit/{filter}/{value}", f.GetFruit)
	r.Get("/fruits", f.GetFruits)
}

// GetFruit is a handler function that retrieve a list of filtered fruits in JSON format.
func (f fruit) GetFruit(w http.ResponseWriter, r *http.Request) {
	filter := chi.URLParam(r, "filter")
	value := chi.URLParam(r, "value")

	fruits, err := f.svc.Get(filter, value)
	if err != nil {
		errJSON(w, r, err)
		return
	}

	render.JSON(w, r, fruits)
}

// GetFruits is a handler function that retrieve all the fruits in JSON format.
func (f fruit) GetFruits(w http.ResponseWriter, r *http.Request) {
	fruits, err := f.svc.GetAll()
	if err != nil {
		errJSON(w, r, err)
		return
	}

	render.JSON(w, r, fruits)
}
