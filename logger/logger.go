// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package logger

import (
	"bufio"
	"io"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/discard"
)

type Logger struct {
	inputReader bufio.Reader
	log.Interface
}

func New(input io.ReadWriter) *Logger {
	return &Logger{
		Interface: &log.Logger{
			Handler: cli.New(input),
			Level:   log.DebugLevel,
		},
		inputReader: *bufio.NewReader(input),
	}
}

func NewNoop() log.Interface {
	return &log.Logger{
		Level:   log.DebugLevel,
		Handler: discard.New(),
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
