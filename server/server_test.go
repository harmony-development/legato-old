package server_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/server"
	"github.com/harmony-development/legato/test"
	"github.com/stretchr/testify/assert"
)

// TestSigInt checks if a SIGINT will exit the server.
func TestSigInt(t *testing.T) {
	t.Parallel()

	notifiedExit := false
	notifiedStart := false

	logger := &test.MockLogger{
		LogCB: func(s string) {
			if strings.HasPrefix(s, "Received ") {
				notifiedExit = true
			} else if strings.HasPrefix(s, "Legato started") {
				notifiedStart = true
			}
		},
	}

	a := assert.New(t)
	i := server.New(&config.Config{}, logger)
	errChan := make(chan error)
	terminateChan := make(chan os.Signal, 1)
	terminateChan <- os.Interrupt
	a.Nil(i.Run(errChan, terminateChan))
	a.True(notifiedStart)
	a.True(notifiedExit)
}

// TestError checks if an error would result in the server exiting with error contents.
func TestError(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	errorReceived := false

	logger := &test.MockLogger{
		LogCB: func(s string) {
			if strings.HasPrefix(s, "Fatal") {
				errorReceived = true
			}
		},
	}

	err := errors.New("server died of cringe")

	i := server.New(&config.Config{}, logger)
	errChan := make(chan error, 1)
	terminateChan := make(chan os.Signal, 1)
	errChan <- err
	a.ErrorIs(i.Run(errChan, terminateChan), err)
	a.True(errorReceived)
}
