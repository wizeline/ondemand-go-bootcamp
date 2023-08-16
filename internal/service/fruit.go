package service

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

var _ Fruit = &fruit{}

type Fruit interface {
	Get(filter, value string) ([]entity.Fruit, error)
	GetAll() ([]entity.Fruit, error)
}

type FruitRepo interface {
	Read(filter, value string) ([]entity.Fruit, error)
	ReadAll() ([]entity.Fruit, error)
}

type fruit struct {
	repo FruitRepo
}

func NewFruit(repo FruitRepo) Fruit {
	logger.Log().Debug().
		Str("service", "Fruit").
		Msg("created service")
	return &fruit{
		repo: repo,
	}
}

func (f fruit) Get(filter, value string) ([]entity.Fruit, error) {
	return f.repo.Read(filter, value)
}

func (f fruit) GetAll() ([]entity.Fruit, error) {
	return f.repo.ReadAll()
}
