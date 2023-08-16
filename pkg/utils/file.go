package utils

import (
	"os"
)

const dataFileMode = "-rw-------"

func ValidateDataFile(filePath string) error {
	if filePath == "" {
		return ErrDataFilePathEmpty
	}

	fInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if fInfo.IsDir() {
		return ErrDataFileIsDir
	}
	if fInfo.Mode().String() != dataFileMode {
		return ErrDataFileInvalidMode
	}

	return nil
}
