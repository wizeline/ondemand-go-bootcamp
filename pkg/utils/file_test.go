package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDataDir = "../../test/data"

func TestValidateDataFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		err      error
		wantErr  bool
	}{
		{
			name:     "Empty",
			filePath: "",
			err:      ErrDataFilePathEmpty,
			wantErr:  true,
		},
		{
			name:     "Directory",
			filePath: testDataDir,
			err:      ErrDataFileIsDir,
			wantErr:  true,
		},
		{
			name:     "Invalid Mode",
			filePath: testDataDir + "/fruits_invalid-mode.csv",
			err:      ErrDataFileInvalidMode,
			wantErr:  true,
		},
		{
			name:     "Valid",
			filePath: testDataDir + "/fruits_valid.csv",
			err:      nil,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDataFile(tt.filePath)
			t.Logf("ERROR: %s", err)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.IsType(t, tt.err, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
