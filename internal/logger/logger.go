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

// `logger` is a type that contains a `zerolog.Logger` type.
//
//	@property	logger - This is the logger that we'll use to log messages.
type logger struct {
	logger *zerolog.Logger
}

// Log returns a pointer of the logger object
func Log() *zerolog.Logger {
	return myLogger.logger
}

// It creates a new logger that writes to the console
func newLogger() *logger {
	var zlogger zerolog.Logger

	// Singleton - configure logger only once
	once.Do(func() {
		var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stderr}

		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zlogger = zerolog.New(writer).With().Timestamp().Logger()
	})

	return &logger{logger: &zlogger}
}
