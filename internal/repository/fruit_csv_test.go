package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
)

const testDataDir = "../../test/data"

func TestGetIntg(t *testing.T) {
	type args struct {
		filter string
		value  string
	}
	tests := []struct {
		name    string
		cfg     configuration.CsvDB
		args    args
		exp     []entity.Fruit
		err     error
		wantErr bool
	}{
		//{
		//	name:    "Configuration Empty",
		//	cfg:     configuration.NewCsvDB("", ""),
		//	err:     Err(""),
		//	wantErr: true,
		//},
		//{
		//	name:     "Arbitrary data file",
		//	fileName: "foo.csv",
		//	//err:      &Err{},
		//	wantErr: true,
		//},
		//{
		//	name: "Invalid data file ",
		//},
		{
			name: "Filter by ID",
			cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
			args: args{filter: "id", value: "5"},
			exp: []entity.Fruit{
				{ID: 5, Name: "orange", Color: "orange"},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Filter by Name",
			cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
			args: args{filter: "name", value: "apple"},
			exp: []entity.Fruit{
				{ID: 1, Name: "apple", Color: "red"},
				{ID: 2, Name: "apple", Color: "green"},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Filter by Color",
			cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
			args: args{filter: "color", value: "green"},
			exp: []entity.Fruit{
				{ID: 2, Name: "apple", Color: "green"},
				{ID: 3, Name: "pear", Color: "green"},
				{ID: 8, Name: "lime", Color: "green"},
				{ID: 9, Name: "grape", Color: "green"},
			},
			err:     nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewFruitCsv(tt.cfg)
			assert.NotNil(t, repo)
			out, err := repo.Read(tt.args.filter, tt.args.value)
			//t.Logf("ERROR: %v", err)
			//t.Logf("OUT: %v", out)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, repo)
				//	assert.IsType(t, tt.err, err)
				return
			}
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}
