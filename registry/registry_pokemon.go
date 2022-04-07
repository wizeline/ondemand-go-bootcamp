package registry

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/interface/controller"
	"github.com/GerardoHP/ondemand-go-bootcamp/usecase/interactor"

	ip "github.com/GerardoHP/ondemand-go-bootcamp/interface/presenter"
	ir "github.com/GerardoHP/ondemand-go-bootcamp/interface/repository"

	pp "github.com/GerardoHP/ondemand-go-bootcamp/usecase/presenter"
	pr "github.com/GerardoHP/ondemand-go-bootcamp/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() pr.PokemonRepositoty {
	return ir.NewPokemonRepository(r.fileName)
}

func (r *registry) NewPokemonPresenter() pp.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
