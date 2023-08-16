package sharedhttp

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

const apiPrefixFmt = "/api/v%d"

var (
	_ Mux = &muxRouter{}
)

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

func NewMux(cfg configuration.SemanticVersion) Mux {
	r := mux.NewRouter().StrictSlash(true)
	return &muxRouter{
		cfg:    cfg,
		router: r,
	}
}

func (m *muxRouter) Add(name string, ctrl controller.HTTP) {
	if m.controllers == nil {
		m.controllers = make(map[string]controller.HTTP)
	}
	m.controllers[name] = ctrl
}

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

func (m *muxRouter) Router() *mux.Router {
	return m.router
}
