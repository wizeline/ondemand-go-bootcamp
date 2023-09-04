package repository

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDataDir = "testdata"

func TestValidateDataFile(t *testing.T) {
	if _, err := os.Stat(testDataDir); errors.Is(err, os.ErrNotExist) {
		require.NoError(t, os.Mkdir(testDataDir, os.ModePerm),
			"create test validData directory is mandatory")
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

	// Clean up test validData directory
	assert.NoError(t, os.RemoveAll(testDataDir),
		"remove test validData directory is mandatory")
}

func TestParseUnixEpoch(t *testing.T) {
	tests := []struct {
		name string
		t    string
		exp  time.Time
		err  error
	}{
		{
			name: "Empty",
			t:    "",
			exp:  time.Time{},
			err:  ErrTimeStringEmpty,
		},
		{
			name: "Arbitrary value",
			t:    "abc",
			exp:  time.Time{},
			err:  ErrInvalidTimeString,
		},
		{
			name: "Valid no fraction",
			t:    "1650536555",
			exp:  time.Date(2022, time.April, 21, 10, 22, 35, 0, time.UTC),
			err:  nil,
		},
		{
			name: "Valid with msec",
			t:    "1650536555.203",
			exp:  time.Date(2022, time.April, 21, 10, 22, 35, int(203*time.Millisecond), time.UTC),
			err:  nil,
		},
		{
			name: "Valid with usec leading zeroes",
			t:    "1650536555.000234",
			exp:  time.Date(2022, time.April, 21, 10, 22, 35, int(234*time.Microsecond), time.UTC),
			err:  nil,
		},
		{
			name: "Valid with nsec leading zeroes",
			t:    "1650536555.000000345",
			exp:  time.Date(2022, time.April, 21, 10, 22, 35, 345, time.UTC),
			err:  nil,
		},
		{
			name: "Valid with many zeroes",
			t:    "1650536555.010020034",
			exp:  time.Date(2022, time.April, 21, 10, 22, 35, int(10020034*time.Nanosecond), time.UTC),
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp, err := parseUnixEpoch(tt.t)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.exp, exp)
		})
	}
}
