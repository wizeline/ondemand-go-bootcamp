package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var _ HTTP = HealthCheck{}

// HealthCheck checks the API health and performance.
type HealthCheck struct{}

// NewHealthCheck returns a new HealthCheck implementation.
func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

// SetRoutes sets a fresh middleware stack for the HealthCheck controller's handle functions and mounts them to the provided sub router.
func (h HealthCheck) SetRoutes(r chi.Router) {
	r.Get("/healthz", h.heartbeat)
}

// heartbeat is a handler function that checks the heart rate.
func (h HealthCheck) heartbeat(w http.ResponseWriter, r *http.Request) {

	// TODO:
	// - Add live external API check(ping)
	// - Add database check.

	render.JSON(w, r, BasicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
