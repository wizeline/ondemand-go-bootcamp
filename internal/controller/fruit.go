package controller

import (
	"encoding/json"
	"net/http"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/gorilla/mux"
)

var (
	_ HTTP      = &fruitHTTP{}
	_ FruitHTTP = &fruitHTTP{}
)

// FruitHTTP set the routes and handler functions related to Fruit
type FruitHTTP interface {
	SetRoutes(r *mux.Router)
	GetFruit(w http.ResponseWriter, r *http.Request)
	GetFruits(w http.ResponseWriter, r *http.Request)
}

// FruitSvc is the abstraction of the fruit service dependency.
type FruitSvc interface {
	Get(filter, value string) (entity.Fruits, error)
	GetAll() (entity.Fruits, error)
}

type fruitHTTP struct {
	svc FruitSvc
}

// NewFruitHTTP returns a new FruitHTTP implementation.
func NewFruitHTTP(svc FruitSvc) FruitHTTP {
	return &fruitHTTP{
		svc: svc,
	}
}

func (f fruitHTTP) SetRoutes(r *mux.Router) {
	r.HandleFunc("/fruit/{filter}/{value}", f.GetFruit).Methods("GET")
	r.HandleFunc("/fruits", f.GetFruits).Methods("GET")
}

// GetFruit is a handler function that retrieve a list of filtered fruits in JSON format.
func (f fruitHTTP) GetFruit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	fruit, err := f.svc.Get(vars["filter"], vars["value"])
	if err != nil {
		errHTTPResponse(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(fruit); err != nil {
		errHTTPResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetFruits is a handler function that retrieve all the fruits in JSON format.
func (f fruitHTTP) GetFruits(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fruits, err := f.svc.GetAll()
	if err != nil {
		errHTTPResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(fruits)
}
