package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheckHTTP checks the API health and performance.
type HealthCheckHTTP struct{}

var _ HTTP = HealthCheckHTTP{}

// NewHealthCheckHTTP returns a new HealthCheckHTTP implementation.
func NewHealthCheckHTTP() HealthCheckHTTP {
	return HealthCheckHTTP{}
}

func (h HealthCheckHTTP) SetRoutes(r *mux.Router) {
	r.HandleFunc("/healthz", h.heartbeat).Methods("GET")
}

// heartbeat is a handler function that checks the heart rate.
func (h HealthCheckHTTP) heartbeat(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Add live infrastructure ping, E.g. Ping to the fake http api

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(BasicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
