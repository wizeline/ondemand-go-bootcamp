package sharedhttp

import (
	"fmt"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const apiPrefixFmt = "/api/v%d"

var _ Chi = &chiMux{}

// Chi is the interface that configure a chi.Mux instance.
type Chi interface {
	Add(name string, ctrl controller.HTTP)
	RegisterRoutes()
	Router() *chi.Mux
}

type chiMux struct {
	router      *chi.Mux
	cfg         configuration.SemanticVersion
	controllers map[string]controller.HTTP
}

// NewChi returns a Chi implementation.
// It allocates a pre-configured chi.Mux instance.
func NewChi(cfg configuration.SemanticVersion) Chi {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return &chiMux{
		cfg:    cfg,
		router: r,
	}
}

// Add appends the controller.HTTP given to the controller's definition list.
func (m *chiMux) Add(name string, ctrl controller.HTTP) {
	if m.controllers == nil {
		m.controllers = make(map[string]controller.HTTP)
	}
	m.controllers[name] = ctrl
}

// RegisterRoutes register the routes defined on the controller's'definition list.
// Each route of the controller.HTTP is prefixed with the API major version.
func (m *chiMux) RegisterRoutes() {
	if len(m.controllers) == 0 {
		return
	}

	m.router.Route(fmt.Sprintf(apiPrefixFmt, m.cfg.MajorVersion()), func(r chi.Router) {
		for name, ctrl := range m.controllers {
			if ctrl != nil {
				ctrl.SetRoutes(r)
				logger.Log().Debug().
					Str("controller", name).
					Msg("registered http controller")
			}
		}
	})
}

// Router returns the configured chi.Mux instance.
func (m *chiMux) Router() *chi.Mux {
	return m.router
}
