package repository

import "github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

type Fruit interface {
	Read(filter, value string) ([]entity.Fruit, error)
	ReadAll() ([]entity.Fruit, error)
	Create(fruit entity.Fruit) error
	CreateAll(fruits []entity.Fruit) error
}
