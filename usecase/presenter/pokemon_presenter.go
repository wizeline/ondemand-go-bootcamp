package presenter

import "github.com/GerardoHP/ondemand-go-bootcamp/domain/model"

type PokemonPresenter interface {
	ResponsePresenter(p []*model.Pokemon) []*model.Pokemon
}
