package repository

import "github.com/GerardoHP/ondemand-go-bootcamp/domain/model"

type PokemonRepositoty interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
