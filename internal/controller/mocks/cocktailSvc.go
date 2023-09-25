package mocks

import (
	ct "github.com/marcos-wz/capstone-go-bootcamp/internal/customtype"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/stretchr/testify/mock"
)

// CocktailSvc is a mock type for the CocktailSvc dependency
type CocktailSvc struct {
	mock.Mock
}

func (o *CocktailSvc) GetFiltered(filter, value string) ([]entity.Cocktail, error) {
	args := o.Called(filter, value)
	return args.Get(0).([]entity.Cocktail), args.Error(1)
}

func (o *CocktailSvc) GetAll() ([]entity.Cocktail, error) {
	args := o.Called()
	return args.Get(0).([]entity.Cocktail), args.Error(1)
}

func (o *CocktailSvc) UpdateDB() (ct.DBOpsSummary, error) {
	args := o.Called()
	return args.Get(0).(ct.DBOpsSummary), args.Error(1)
}

// NewCocktailSvc creates a new instance of the CocktailSvc of type Mock.
func NewCocktailSvc() *CocktailSvc {
	return &CocktailSvc{}
}
