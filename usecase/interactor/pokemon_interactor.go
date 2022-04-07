package interactor

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
	"github.com/GerardoHP/ondemand-go-bootcamp/usecase/presenter"
	"github.com/GerardoHP/ondemand-go-bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepositoty
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepositoty, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{PokemonRepository: r, PokemonPresenter: p}
}

func (pk *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := pk.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return pk.PokemonPresenter.ResponsePresenter(p), nil
}
