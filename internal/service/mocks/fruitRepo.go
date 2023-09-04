package mocks

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/stretchr/testify/mock"
)

// FruitRepo is a mock type for the FruitRepo type
type FruitRepo struct {
	mock.Mock
}

func (o *FruitRepo) ReadAll() (entity.Fruits, error) {
	args := o.Called()
	return args.Get(0).(entity.Fruits), args.Error(1)
}

// NewFruitRepo creates a new instance of the FruitRepo of type Mock.
func NewFruitRepo() *FruitRepo {
	return &FruitRepo{}
}
