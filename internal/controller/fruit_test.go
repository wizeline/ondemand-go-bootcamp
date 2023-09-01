package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/controller/mocks"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FruitTestSuite struct {
	suite.Suite
	svc *mocks.FruitSvc
}

func TestFruitTestSuite(t *testing.T) {
	suite.Run(t, new(FruitTestSuite))
}

func (s *FruitTestSuite) SetupSuite() {
	svc := mocks.NewFruitSvc()
	require.NotNil(s.T(), svc)
	s.svc = svc
}

func (s *FruitTestSuite) TestGetFruit() {
	type params struct {
		filter string
		value  string
	}
	tests := []struct {
		name    string
		params  params
		code    int
		svcResp entity.Fruits
		svcErr  error
		wantErr bool
	}{
		{
			name:    "Repository CSV Error",
			params:  params{filter: "id", value: "2"},
			code:    http.StatusInternalServerError,
			svcResp: nil,
			svcErr:  &repository.CsvErr{},
			wantErr: true,
		},
		{
			name:    "Empty",
			params:  params{},
			code:    http.StatusNotFound,
			svcResp: nil,
			svcErr:  &service.FilterErr{Err: service.ErrFilterNotSupported},
			wantErr: true,
		},
		{
			name:    "Arbitrary",
			params:  params{filter: "foo", value: "some-value"},
			code:    http.StatusUnprocessableEntity,
			svcResp: nil,
			svcErr:  &service.FilterErr{Err: service.ErrFilterNotImplemented},
			wantErr: true,
		},
		{
			name:    "Bad Value",
			params:  params{filter: "id", value: "asd"},
			code:    http.StatusUnprocessableEntity,
			svcResp: nil,
			svcErr:  &service.FilterErr{},
			wantErr: true,
		},
		{
			name:    "Not Found",
			params:  params{filter: "id", value: "123456"},
			code:    http.StatusOK,
			svcResp: entity.Fruits{},
			svcErr:  nil,
			wantErr: false,
		},
		{
			name:   "ID",
			params: params{filter: "id", value: "2"},
			code:   http.StatusOK,
			svcResp: entity.Fruits{
				{ID: 2, Name: "apple", Color: "green"},
			},
			svcErr:  nil,
			wantErr: false,
		},
		{
			name:   "Name",
			params: params{filter: "name", value: "apple"},
			code:   http.StatusOK,
			svcResp: entity.Fruits{
				{ID: 1, Name: "apple", Color: "red"},
				{ID: 2, Name: "apple", Color: "green"},
			},
			svcErr:  nil,
			wantErr: false,
		},
		{
			name:   "Color",
			params: params{filter: "color", value: "green"},
			code:   http.StatusOK,
			svcResp: entity.Fruits{
				{ID: 2, Name: "apple", Color: "green"},
				{ID: 3, Name: "pear", Color: "green"},
			},
			svcErr:  nil,
			wantErr: false,
		},
		{
			name:   "Country",
			params: params{filter: "country", value: "Mexico"},
			code:   http.StatusOK,
			svcResp: entity.Fruits{
				{ID: 1, Name: "apple", Color: "red", Country: "Mexico"},
				{ID: 2, Name: "apple", Color: "green", Country: "Mexico"},
			},
			svcErr:  nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Dependencies
			s.svc.ExpectedCalls = nil
			s.svc.On("Get", tt.params.filter, tt.params.value).
				Return(tt.svcResp, tt.svcErr)
			require.NotNil(t, s.svc)

			ctrl := NewFruit(s.svc)
			require.NotNil(t, ctrl)

			// Request
			path := fmt.Sprintf("/fruit/%v/%v",
				tt.params.filter,
				tt.params.value)
			req, err := http.NewRequest("GET", path, nil)
			require.Nil(t, err)

			// Server instance
			rr := httptest.NewRecorder()
			srv := s.newRouter(ctrl)
			srv.ServeHTTP(rr, req)

			// Tests
			assert.Equal(t, tt.code, rr.Code)
			if tt.wantErr {
				return
			}

			var resp entity.Fruits
			require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
			assert.Len(t, tt.svcResp, len(resp))
			assert.EqualValues(t, tt.svcResp, resp)
		})
	}
}

func (s *FruitTestSuite) TestGetFruits() {
	var testData = entity.Fruits{
		{ID: 1, Name: "apple", Color: "red", Country: "Mexico"},
		{ID: 2, Name: "apple", Color: "green", Country: "Mexico"},
		{ID: 3, Name: "pear", Color: "green", Country: "Brazil"},
		{ID: 4, Name: "banana", Color: "yellow", Country: "Brazil"},
		{ID: 5, Name: "orange", Color: "orange", Country: "USA"},
	}
	tests := []struct {
		name    string
		code    int
		svcResp entity.Fruits
		svcErr  error
		wantErr bool
	}{
		{
			name:    "Repository CSV Error",
			code:    http.StatusInternalServerError,
			svcResp: nil,
			svcErr:  &repository.CsvErr{},
			wantErr: true,
		},
		{
			name:    "Not Found",
			code:    http.StatusOK,
			svcResp: entity.Fruits{},
			svcErr:  nil,
			wantErr: false,
		},
		{
			name:    "All Records",
			code:    http.StatusOK,
			svcResp: testData,
			svcErr:  nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.svc.ExpectedCalls = nil
			s.svc.On("GetAll").Return(tt.svcResp, tt.svcErr)
			ctrl := NewFruit(s.svc)
			require.NotNil(t, ctrl)

			// Request
			req, err := http.NewRequest("GET", "/fruits", nil)
			require.Nil(t, err)

			// Server instance
			rr := httptest.NewRecorder()
			srv := s.newRouter(ctrl)
			srv.ServeHTTP(rr, req)

			// Tests
			assert.Equal(t, tt.code, rr.Code)
			if tt.wantErr {
				return
			}

			var resp entity.Fruits
			require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
			assert.Len(t, tt.svcResp, len(resp))
			assert.EqualValues(t, tt.svcResp, resp)
		})
	}
}

func (s *FruitTestSuite) newRouter(ctrl HTTP) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ctrl.SetRoutes(r)
	return r
}
