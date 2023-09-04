package repository

import (
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

// parseUnixEpoch converts the given string time (unix time in seconds) to time.Time.
// t should be in the format "1650536555.203"; decimal portion is optional.
func parseUnixEpoch(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, ErrTimeStringEmpty
	}
	timeStr := strings.SplitN(t, ".", 2)
	if length := len(timeStr); length == 0 || length > 2 {
		return time.Time{}, ErrInvalidTimeString
	}

	sec, err := strconv.ParseInt(timeStr[0], 10, 64)
	if err != nil {
		return time.Time{}, ErrInvalidTimeString
	}

	var nsec int64
	if len(timeStr) == 2 {
		fraction, err := strconv.ParseFloat(timeStr[1], 64)
		if err != nil {
			return time.Time{}, ErrInvalidTimeString
		}
		fractionScale := math.Pow10(9 - len(timeStr[1]))
		nsec = int64(fraction * fractionScale)
	}

	return time.Unix(sec, nsec).UTC(), nil
}
