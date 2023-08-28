package logger

import (
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	myLogger = newLogger()
	once     sync.Once
)

// logger is a type that contains a zerolog.Logger type.
// This is the logger that we'll use to log messages.
type logger struct {
	logger *zerolog.Logger
}

// Log returns the pointer of the logger object
func Log() *zerolog.Logger {
	return myLogger.logger
}

// newLogger creates a new logger that writes to the console.
// The instance can be configured only once. Singleton type.
func newLogger() *logger {
	var zlogger zerolog.Logger
	once.Do(func() {
		var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stderr}

		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zlogger = zerolog.New(writer).With().Timestamp().Logger()
	})
	return &logger{logger: &zlogger}
}
