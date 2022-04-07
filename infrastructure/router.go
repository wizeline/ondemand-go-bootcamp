package infrastructure

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/interfaces"
	"github.com/go-chi/chi"
)

func Dispatch(logger interfaces.Logger) {
	pokemonController := interfaces.PokemonController{Logger: logger}
	r := chi.NewRouter()
	r.Get("/pokemons", pokemonController.Index)
}
