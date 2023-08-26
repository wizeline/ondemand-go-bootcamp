package service

import (
	"errors"
	"strconv"
	"testing"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FruitTestSuite struct {
	suite.Suite
	repo *mocks.FruitRepo
	data entity.Fruits
}

func TestFruitTestSuite(t *testing.T) {
	suite.Run(t, new(FruitTestSuite))
}

func (s *FruitTestSuite) SetupSuite() {
	repo := mocks.NewFruitRepo()
	require.NotNil(s.T(), repo)
	s.repo = repo

	s.data = entity.Fruits{
		{ID: 1, Name: "apple", Color: "red"},
		{ID: 2, Name: "apple", Color: "green"},
		{ID: 3, Name: "pear", Color: "green"},
		{ID: 4, Name: "banana", Color: "yellow"},
		{ID: 5, Name: "orange", Color: "orange"},
	}
}

func (s *FruitTestSuite) TestGet() {
	type args struct {
		filter string
		value  string
	}
	tests := []struct {
		name     string
		args     args
		exp      entity.Fruits
		err      error
		repoResp entity.Fruits
		repoErr  error
	}{
		{
			name:     "Repository CSV Error",
			args:     args{filter: idFilter.String(), value: "4"},
			exp:      nil,
			err:      &repository.CsvErr{},
			repoResp: nil,
			repoErr:  &repository.CsvErr{},
		},
		{
			name:     "Filter value empty",
			args:     args{filter: idFilter.String(), value: ""},
			exp:      nil,
			err:      &FilterErr{ErrFilterValueEmpty},
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name:     "Arbitrary Filter",
			args:     args{filter: "foo", value: "foo-value"},
			exp:      nil,
			err:      &FilterErr{ErrFilterNotImplemented},
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name:     "Not Found",
			args:     args{filter: idFilter.String(), value: "123456"},
			exp:      entity.Fruits{},
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name:     "Bad ID",
			args:     args{filter: idFilter.String(), value: "foo-id"},
			exp:      entity.Fruits{},
			err:      &FilterErr{&strconv.NumError{}},
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name: "Filter ID",
			args: args{filter: idFilter.String(), value: "4"},
			exp: entity.Fruits{
				{ID: 4, Name: "banana", Color: "yellow"},
			},
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name: "Filter Name",
			args: args{filter: nameFilter.String(), value: "apple"},
			exp: entity.Fruits{
				{ID: 1, Name: "apple", Color: "red"},
				{ID: 2, Name: "apple", Color: "green"},
			},
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name: "Filter Color",
			args: args{filter: colorFilter.String(), value: "green"},
			exp: entity.Fruits{
				{ID: 2, Name: "apple", Color: "green"},
				{ID: 3, Name: "pear", Color: "green"},
			},
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
		{
			name:     "Filter Country",
			args:     args{filter: countryFilter.String(), value: "Mexico"},
			exp:      entity.Fruits{},
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.repo.ExpectedCalls = nil
			s.repo.On("ReadAll").Return(tt.repoResp, tt.repoErr)
			svc := NewFruit(s.repo)

			out, err := svc.Get(tt.args.filter, tt.args.value)
			if tt.err != nil {
				require.NotNil(t, err)
				assert.Nil(t, out)
				assert.IsType(t, tt.err, err)
				if errWrp := errors.Unwrap(tt.err); errWrp != nil {
					assert.IsType(t, errWrp, errors.Unwrap(err))
				}
				return
			}
			require.NotNil(t, out)
			require.Nil(t, err)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}

func (s *FruitTestSuite) TestGetAll() {
	tests := []struct {
		name     string
		exp      entity.Fruits
		err      error
		repoResp entity.Fruits
		repoErr  error
	}{
		{
			name:     "Repository CSV Error",
			exp:      nil,
			err:      &repository.CsvErr{},
			repoResp: nil,
			repoErr:  &repository.CsvErr{},
		},
		{
			name:     "Not Found",
			exp:      entity.Fruits{},
			err:      nil,
			repoResp: entity.Fruits{},
			repoErr:  nil,
		},
		{
			name:     "All records",
			exp:      s.data,
			err:      nil,
			repoResp: s.data,
			repoErr:  nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.repo.ExpectedCalls = nil
			s.repo.On("ReadAll").Return(tt.repoResp, tt.repoErr)
			svc := NewFruit(s.repo)

			out, err := svc.GetAll()
			if tt.err != nil {
				//require.NotNil(t, err)
				assert.Nil(t, out)
				assert.IsType(t, tt.err, err)
				if errWrp := errors.Unwrap(tt.err); errWrp != nil {
					assert.IsType(t, errWrp, errors.Unwrap(err))
				}
				return
			}
			require.NotNil(t, out)
			require.Nil(t, err)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}
