package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct{ zerolog.Logger }

var log *Logger

func Initialize(isProduction bool) {
	if isProduction == true {
		log = &Logger{zerolog.New(os.Stdout).With().Timestamp().Logger()}
	} else {
		log = &Logger{zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()}
	}
}

func Get() *Logger {
	return log
}
