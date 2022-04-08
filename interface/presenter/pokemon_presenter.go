package presenter

import (
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
)

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePresenter(pk []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePresenter(pk []*model.Pokemon) []*model.Pokemon {
	for _, p := range pk {
		p.Name = strings.ToUpper(p.Name)
	}

	return pk
}
