package repository

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDataDir = "testdata"

func TestValidateDataFile(t *testing.T) {
	if _, err := os.Stat(testDataDir); errors.Is(err, os.ErrNotExist) {
		require.NoError(t, os.Mkdir(testDataDir, os.ModePerm),
			"create test data directory is mandatory")
	}

	tests := []struct {
		name     string
		filePath string
		filePerm os.FileMode
		err      error
		wantFile bool
	}{
		{
			name:     "File path empty",
			filePath: "",
			err:      ErrFilePathEmpty,
		},
		{
			name:     "Directory",
			filePath: testDataDir,
			err:      ErrFilePathIsDir,
		},
		{
			name:     "Arbitrary",
			filePath: filepath.Join(testDataDir, "foo.txt"),
			err:      &fs.PathError{},
		},
		{
			name:     "Invalid Mode",
			filePath: filepath.Join(testDataDir, "invalid-mode.txt"),
			filePerm: 0666,
			err:      ErrFileModeInvalid,
			wantFile: true,
		},
		{
			name:     "Valid",
			filePath: filepath.Join(testDataDir, "valid.txt"),
			filePerm: 0600,
			err:      nil,
			wantFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantFile {
				require.NoError(t, os.WriteFile(tt.filePath, nil, tt.filePerm))
				t.Logf("created file(%v): %v", tt.filePerm, tt.filePath)
			}

			err := validateDataFile(tt.filePath)
			if tt.err != nil {
				assert.NotNil(t, err)
				assert.IsType(t, tt.err, err)
				return
			}
			assert.Nil(t, err)
		})
	}

	// Clean up test data directory
	assert.NoError(t, os.RemoveAll(testDataDir),
		"remove test data directory is mandatory")
}
