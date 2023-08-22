package repository

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
)

var (
	testFruitValidRecords = []byte(`
1,apple,,red,,,,
2,apple,,green,,,,
3,pear,,green,,,,
4,banana,,yellow,,,,
5,orange,,orange,,,,`)
	testFruitWrongNumFields = []byte(`
1,apple,,red
2,apple,,green
3,pear,,green,,,,
4,banana,,yellow
5,orange,,orange`)
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
	// Initialize test data directory
	if _, err := os.Stat(testDataDir); errors.Is(err, os.ErrNotExist) {
		require.NoError(s.T(), os.Mkdir(testDataDir, os.ModePerm),
			"creation test data directory is mandatory")
	} else {
		require.Nil(s.T(), err)
	}

	cfg := configuration.NewCsvDB("fruits_test.csv", testDataDir)
	require.NotNil(s.T(), cfg)
	s.cfg = cfg

	// Initialize test data file
	require.NoError(s.T(), os.WriteFile(cfg.FilePath(), nil, 0600))

	repo, err := NewFruitCsv(cfg)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), repo)
	s.repo = repo
}

func (s *FruitCsvTestSuite) TearDownSuite() {
	assert.NoError(s.T(), os.RemoveAll(testDataDir),
		"remove test data directory is mandatory")
	s.T().Logf("removed test data directory: %v", testDataDir)
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
			name: "Wrong num fields",
			data: testFruitWrongNumFields,
			exp: entity.Fruits{
				{ID: 1, Name: "apple", Color: "red"},
				{ID: 2, Name: "apple", Color: "green"},
				{ID: 4, Name: "banana", Color: "yellow"},
				{ID: 5, Name: "orange", Color: "orange"},
			},
			err: nil,
		},
		{
			name: "Valid",
			data: testFruitValidRecords,
			exp: entity.Fruits{
				{ID: 1, Name: "apple", Color: "red"},
				{ID: 2, Name: "apple", Color: "green"},
				{ID: 3, Name: "pear", Color: "green"},
				{ID: 4, Name: "banana", Color: "yellow"},
				{ID: 5, Name: "orange", Color: "orange"},
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
			require.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}

}
