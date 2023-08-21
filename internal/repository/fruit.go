package repository

import "github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

type Fruit interface {
	ReadAll() (entity.Fruits, error)
	Create(fruit entity.Fruit) error
	CreateAll(fruits entity.Fruits) error
}
