package mocks

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/stretchr/testify/mock"
)

// FruitSvc is a mock type for the FruitSvc type
type FruitSvc struct {
	mock.Mock
}

func (o *FruitSvc) Get(filter, value string) (entity.Fruits, error) {
	args := o.Called(filter, value)
	return args.Get(0).(entity.Fruits), args.Error(1)
}

func (o *FruitSvc) GetAll() (entity.Fruits, error) {
	args := o.Called()
	return args.Get(0).(entity.Fruits), args.Error(1)
}

// NewFruitSvc creates a new instance of FruitSvc of type Mock.
func NewFruitSvc() *FruitSvc {
	return &FruitSvc{}
}
