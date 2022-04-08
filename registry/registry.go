package registry

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/interface/controller"
)

type registry struct {
	fileName string
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(fn string) Registry {
	return &registry{fileName: fn}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
