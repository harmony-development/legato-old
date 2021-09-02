package logger

import (
	"os"
	"strings"

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

func Indent(level log.Level, s string, finish string) string {
	var output strings.Builder
	for _, line := range strings.Split(s, "\n") {
		output.WriteRune('\n')
		output.WriteString(cli.Colors[level].Sprintf("   ║  ") + line)
	}
	output.WriteString(cli.Colors[level].Sprintf("\n   ╚  ") + finish)
	return output.String()
}
