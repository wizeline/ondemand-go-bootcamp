package interfaces

type PokemonController struct {
	Logger Logger
}

func NewPokemonController(logger Logger) *PokemonController {
	return &PokemonController{Logger: logger}
}
