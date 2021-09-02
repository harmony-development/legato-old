package logger

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

type Logger struct {
	log.Interface
}

func New() *Logger {
	return &Logger{
		Interface: &log.Logger{
			Handler: cli.New(os.Stdout),
			Level:   log.DebugLevel,
		},
	}
}
