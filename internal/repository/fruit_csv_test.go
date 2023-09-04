package repository

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FruitCsvTestSuite struct {
	suite.Suite
	cfg  configuration.CsvDB
	repo Fruit
}

func TestFruitCsvTestSuite(t *testing.T) {
	suite.Run(t, new(FruitCsvTestSuite))
}

func (s *FruitCsvTestSuite) SetupSuite() {
	// Initialize test validData directory
	if _, err := os.Stat(testDataDir); errors.Is(err, os.ErrNotExist) {
		require.NoError(s.T(), os.Mkdir(testDataDir, os.ModePerm),
			"creation test validData directory is mandatory")
	} else {
		require.Nil(s.T(), err)
	}

	cfg := configuration.NewCsvDB("fruits_test.csv", testDataDir)
	require.NotNil(s.T(), cfg)
	s.cfg = cfg

	// Initialize test validData file
	require.NoError(s.T(), os.WriteFile(cfg.FilePath(), nil, 0600))

	repo, err := NewFruitCsv(cfg)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), repo)
	s.repo = repo
}

func (s *FruitCsvTestSuite) TearDownSuite() {
	assert.NoError(s.T(), os.RemoveAll(testDataDir),
		"remove test validData directory is mandatory")
	s.T().Logf("removed test validData directory: %v", testDataDir)
}

func (s *FruitCsvTestSuite) TestNewFruitCsv() {
	tests := []struct {
		name     string
		fileName string
		filePerm os.FileMode
		exp      Fruit
		err      error
		wantFile bool
	}{
		{
			name:     "File name empty",
			fileName: "",
			exp:      nil,
			err:      &CsvErr{ErrFilePathIsDir},
			wantFile: false,
		},
		{
			name:     "Arbitrary file",
			fileName: "foo.csv",
			exp:      nil,
			err:      &CsvErr{&fs.PathError{}},
			wantFile: false,
		},
		{
			name:     "File permissions",
			fileName: "invalid.csv",
			filePerm: 0666,
			exp:      nil,
			err:      &CsvErr{ErrFileModeInvalid},
			wantFile: true,
		},
		{
			name:     "Valid",
			fileName: "fruits_valid.csv",
			filePerm: 0600,
			exp:      &fruitCsv{},
			err:      nil,
			wantFile: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			cfg := configuration.NewCsvDB(tt.fileName, s.cfg.DataDir())
			require.NotNil(t, cfg)
			if tt.wantFile {
				require.NoError(t, os.WriteFile(cfg.FilePath(), nil, tt.filePerm))
			}

			out, err := NewFruitCsv(cfg)
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
			assert.IsType(t, tt.exp, out)
		})
	}
}

func (s *FruitCsvTestSuite) TestReadAll() {
	tests := []struct {
		name string
		data []byte
		exp  entity.Fruits
		err  error
	}{
		{
			name: "File empty",
			data: nil,
			exp:  entity.Fruits{},
			err:  nil,
		},
		{
			name: "Valid",
			data: []byte(`1,apple,,red,Mexico,1693843857,1693843857,1693843857
2,apple,,green,Brazil,1693843857,1693843857,1693843857
3,pear,,green,USA,1693843857,1693843857,1693843857
4,banana,,yellow,Brazil,1693843857,1693843857,1693843857
5,orange,,orange,USA,1693843857,1693843857,1693843857`),
			exp: entity.Fruits{
				{ID: 1, Name: "apple", Description: "", Color: "red", Country: "Mexico", ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), CreatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), UpdatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC)},
				{ID: 2, Name: "apple", Description: "", Color: "green", Country: "Brazil", ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), CreatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), UpdatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC)},
				{ID: 3, Name: "pear", Description: "", Color: "green", Country: "USA", ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), CreatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), UpdatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC)},
				{ID: 4, Name: "banana", Description: "", Color: "yellow", Country: "Brazil", ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), CreatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), UpdatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC)},
				{ID: 5, Name: "orange", Description: "", Color: "orange", Country: "USA", ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), CreatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC), UpdatedAt: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC)},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			require.NoError(t, os.WriteFile(s.cfg.FilePath(), tt.data, 0600))

			out, err := s.repo.ReadAll()
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
			//require.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}

}

func TestParseFruit(t *testing.T) {
	repo := fruitCsv{
		numFields: reflect.TypeOf(entity.Fruit{}).NumField(),
	}
	tests := []struct {
		name   string
		record []string
		exp    entity.Fruit
		err    error
	}{
		{
			name:   "Nil",
			record: nil,
			exp:    entity.Fruit{},
			err:    ErrCSVRecordEmpty,
		},
		{
			name:   "Empty",
			record: []string{},
			exp:    entity.Fruit{},
			err:    ErrCSVRecordEmpty,
		},
		{
			name:   "ID Empty",
			record: []string{"", "foo", "some description", "red", "Mexico", "1693843857", "1693843857", "1693843857"},
			exp:    entity.Fruit{},
			err:    &strconv.NumError{},
		},
		{
			name:   "Bad ID",
			record: []string{"bad-id", "foo", "some description", "red", "Mexico", "1693843857", "1693843857", "1693843857"},
			exp:    entity.Fruit{},
			err:    &strconv.NumError{},
		},
		{
			name:   "Bad Expiration Date",
			record: []string{"1", "foo", "some description", "red", "Mexico", "2023-09-04", "1693843857", "1693843857"},
			exp:    entity.Fruit{},
			err:    ErrInvalidTimeString,
		},
		{
			name:   "Bad Created At",
			record: []string{"1", "foo", "some description", "red", "Mexico", "1693843857", "2023-09-04", "1693843857"},
			exp:    entity.Fruit{},
			err:    ErrInvalidTimeString,
		},
		{
			name:   "Bad Updated At",
			record: []string{"1", "foo", "some description", "red", "Mexico", "1693843857", "1693843857", "2023-09-04"},
			exp:    entity.Fruit{},
			err:    ErrInvalidTimeString,
		},
		{
			name:   "Valid",
			record: []string{"1", "foo", "some description", "red", "Mexico", "1693843857", "1693843857", "1693843857"},
			exp: entity.Fruit{
				ID: 1, Name: "foo", Description: "some description", Color: "red", Country: "Mexico",
				ExpirationDate: time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC),
				CreatedAt:      time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC),
				UpdatedAt:      time.Date(2023, time.September, 4, 16, 10, 57, 0, time.UTC),
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := repo.parseFruit(tt.record)
			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, entity.Fruit{}, out)
				assert.IsType(t, tt.err, err)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, out)
			assert.Equal(t, tt.exp, out)
		})
	}
}
