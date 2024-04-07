package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct{ zerolog.Logger }

var log *Logger

func Get() *Logger {
	if log == nil {
		log = &Logger{zerolog.New(os.Stdout)}
	}
	return log
}
