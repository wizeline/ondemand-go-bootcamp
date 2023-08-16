package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthCheckHTTP struct{}

var _ HTTP = HealthCheckHTTP{}

func NewHealthCheckHTTP() HealthCheckHTTP {
	return HealthCheckHTTP{}
}

func (h HealthCheckHTTP) SetRoutes(r *mux.Router) {
	r.HandleFunc("/healthz", h.heartbeat).Methods("GET")
}

func (h HealthCheckHTTP) heartbeat(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Add live infrastructure ping, E.g. Ping to the fake http api

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(BasicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
