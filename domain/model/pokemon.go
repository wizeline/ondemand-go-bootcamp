package model

import (
	"errors"
	"fmt"
	"strings"
)

// const(
// 	const c string="pokemons.csv"
// )

const (
	fileName string = "pokemons.csv"
)

type Pokemon struct {
	ID   string
	Name string
}

func (p Pokemon) String() string {
	pstr := fmt.Sprintf("%v, %v", p.ID, p.Name)
	return pstr
}

func FileName() string { return fileName }

func ToPokemon(s string) (*Pokemon, error) {
	str := strings.Split(s, ",")
	if len(str) != 2 {
		return &Pokemon{}, errors.New("It's not a Pokemon")
	}

	return &Pokemon{ID: str[0], Name: str[1]}, nil
}
