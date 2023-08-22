package service

import (
	"fmt"
	"strconv"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

const (
	idFilter      FilterFruit = "id"
	nameFilter    FilterFruit = "name"
	colorFilter   FilterFruit = "color"
	countryFilter FilterFruit = "country"
)

var _ fmt.Stringer = FilterFruit("")

var filterFruitMap = map[string]FilterFruit{
	"id":      idFilter,
	"name":    nameFilter,
	"color":   colorFilter,
	"country": countryFilter,
}

var _ Fruit = &fruit{}

type Fruit interface {
	Get(filter, value string) (entity.Fruits, error)
	GetAll() (entity.Fruits, error)
}

type FruitRepo interface {
	ReadAll() (entity.Fruits, error)
}

type fruit struct {
	repo FruitRepo
}

type FilterFruit string

func (f FilterFruit) String() string {
	return string(f)
}

func NewFruit(repo FruitRepo) Fruit {
	logger.Log().Debug().
		Str("service", "Fruit").
		Msg("created service")
	return &fruit{
		repo: repo,
	}
}

func (f fruit) Get(filter, value string) (entity.Fruits, error) {
	if value == "" {
		return nil, &FilterErr{ErrFilterValueEmpty}
	}
	filterType, found := filterFruitMap[filter]
	if !found {
		return nil, &FilterErr{ErrFilterNotSupported}
	}

	recs, err := f.repo.ReadAll()
	if err != nil {
		return nil, err
	}

	switch filterType {
	case idFilter:
		id, e := strconv.Atoi(value)
		if e != nil {
			return nil, &FilterErr{e}
		}
		return f.getById(id, recs), nil
	case nameFilter:
		return f.getByName(value, recs), nil
	case colorFilter:
		return f.getByColor(value, recs), nil
	case countryFilter:
		return f.getByCountry(value, recs), nil
	default:
		logger.Log().Error().Err(ErrFilterNotImplemented).
			Str("filter_type", filterType.String()).
			Str("value", value).
			Msgf("getting fruit records failed")
		return nil, &FilterErr{ErrFilterNotImplemented}
	}
}

func (f fruit) GetAll() (entity.Fruits, error) {
	return f.repo.ReadAll()
}

func (f fruit) getById(id int, recs entity.Fruits) entity.Fruits {
	fruits := entity.Fruits{}
	for _, rec := range recs {
		if id == rec.ID {
			fruits = append(fruits, rec)
		}
	}
	return fruits
}

func (f fruit) getByName(name string, recs entity.Fruits) entity.Fruits {
	fruits := entity.Fruits{}
	for _, rec := range recs {
		if name == rec.Name {
			fruits = append(fruits, rec)
		}
	}
	return fruits
}

func (f fruit) getByColor(color string, recs entity.Fruits) entity.Fruits {
	fruits := entity.Fruits{}
	for _, rec := range recs {
		if color == rec.Color {
			fruits = append(fruits, rec)
		}
	}
	return fruits
}

func (f fruit) getByCountry(country string, recs entity.Fruits) entity.Fruits {
	fruits := entity.Fruits{}
	for _, rec := range recs {
		if country == rec.Country {
			fruits = append(fruits, rec)
		}
	}
	return fruits
}
