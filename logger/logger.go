package logger

import (
	"bufio"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

type Logger struct {
	inputReader bufio.Reader
	log.Interface
}

func New(input *os.File) *Logger {
	return &Logger{
		Interface: &log.Logger{
			Handler: cli.New(input),
			Level:   log.DebugLevel,
		},
		inputReader: *bufio.NewReader(input),
	}
}

func Indent(level log.Level, s string, finish *string) string {
	var output strings.Builder
	for _, line := range strings.Split(s, "\n") {
		output.WriteRune('\n')
		output.WriteString(cli.Colors[level].Sprintf("   ║  ") + line)
	}
	if finish != nil {
		output.WriteString(cli.Colors[level].Sprintf("\n   ╚  ") + *finish)
	}
	return output.String()
}
