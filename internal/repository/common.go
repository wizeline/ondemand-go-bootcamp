package repository

import (
	"os"
)

const dataFileMode = "-rw-------"

func validateDataFile(filePath string) error {
	if filePath == "" {
		return ErrFilePathEmpty
	}

	fInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if fInfo.IsDir() {
		return ErrFilePathIsDir
	}
	if fInfo.Mode().String() != dataFileMode {
		return ErrFileModeInvalid
	}

	return nil
}
