package sharedhttp

import (
	"fmt"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"

	"github.com/gorilla/mux"
)

const apiPrefixFmt = "/api/v%d"

var (
	_ Mux = &muxRouter{}
)

// Mux is the interface that set up the mux.Router instance.
type Mux interface {
	Add(name string, ctrls controller.HTTP)
	RegisterRoutes()
	Router() *mux.Router
}

type muxRouter struct {
	cfg         configuration.SemanticVersion
	controllers map[string]controller.HTTP
	router      *mux.Router
}

// NewMux returns a new Mux implementation.
func NewMux(cfg configuration.SemanticVersion) Mux {
	r := mux.NewRouter().StrictSlash(true)
	return &muxRouter{
		cfg:    cfg,
		router: r,
	}
}

// Add appends the controller.HTTP given to the controller definition list.
func (m *muxRouter) Add(name string, ctrl controller.HTTP) {
	if m.controllers == nil {
		m.controllers = make(map[string]controller.HTTP)
	}
	m.controllers[name] = ctrl
}

// RegisterRoutes register the routes defined on the controller definition list.
// The application's version is prefixed to each route of the controller.
func (m *muxRouter) RegisterRoutes() {
	prefix := fmt.Sprintf(apiPrefixFmt, m.cfg.MajorVersion())
	s := m.router.PathPrefix(prefix).Subrouter()

	for name, ctrl := range m.controllers {
		if ctrl != nil {
			logger.Log().Debug().
				Str("controller", name).
				Msg("registered http controller")
			ctrl.SetRoutes(s)
		}
	}
}

// Router returns the pointer of the mux.Router instance.
func (m *muxRouter) Router() *mux.Router {
	return m.router
}
