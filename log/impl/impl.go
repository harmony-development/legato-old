package impl

import (
	"fmt"

	"github.com/harmony-development/legato/config"
)

// TODO: come up with a better package name
// Logger is a implementation of the log interface, exists as the primary implementation.
type Logger struct {
	c *config.Config
}

func New(cfg *config.Config) *Logger {
	return &Logger{
		c: cfg,
	}
}

func (l *Logger) Log(args ...interface{}) {
	fmt.Println(args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
