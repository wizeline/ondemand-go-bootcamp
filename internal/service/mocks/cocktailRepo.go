package mocks

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/stretchr/testify/mock"
)

// CocktailRepo is a mock type for the CocktailRepo dependency
type CocktailRepo struct {
	mock.Mock
}

func (o *CocktailRepo) ReadAll() ([]entity.Cocktail, error) {
	args := o.Called()
	return args.Get(0).([]entity.Cocktail), args.Error(1)
}

func (o *CocktailRepo) ReplaceData(recs []entity.Cocktail) error {
	args := o.Called(recs)
	return args.Error(0)
}

func (o *CocktailRepo) FetchData() ([]entity.Cocktail, error) {
	args := o.Called()
	return args.Get(0).([]entity.Cocktail), args.Error(1)
}

// NewCocktailRepo creates a new instance of the CocktailRepo of type Mock.
func NewCocktailRepo() *CocktailRepo {
	return &CocktailRepo{}
}
