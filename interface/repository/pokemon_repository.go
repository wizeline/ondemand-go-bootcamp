package repository

import (
	"bufio"
	"log"
	"os"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
)

type pokemonRepository struct {
	pokemonFile string
}

type UserRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository(fn string) UserRepository {
	return &pokemonRepository{pokemonFile: fn}
}

func (pk *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	file, err := os.Open(pk.pokemonFile)
	if err != nil {
		log.Fatal("Failed to open", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		pk, errPk := model.ToPokemon(scanner.Text())
		if errPk != nil {
			log.Fatal(err)
			continue
		}

		p = append(p, pk)
	}

	return p, nil
}
