package interfaces

import "net/http"

type PokemonController struct {
	Logger Logger
}

func NewPokemonController(logger Logger) *PokemonController {
	return &PokemonController{Logger: logger}
}

func (uc *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	uc.Logger.LogAccess("access!!!")
}
