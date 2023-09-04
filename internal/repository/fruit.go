package repository

import "github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

// Fruit is the interface that describes the behavior of the Fruit repository
type Fruit interface {
	ReadAll() (entity.Fruits, error)
	Create(fruit entity.Fruit) error
	CreateAll(fruits entity.Fruits) error
}
