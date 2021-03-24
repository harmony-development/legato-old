package test

import (
	"testing"

	"github.com/harmony-development/legato/server/logger"
)

type MockLogger struct {
	T testing.TB
}

func (m MockLogger) CheckException(err error) {
	panic("unimplemented")
}
func (m MockLogger) Exception(err error) {
	panic("unimplemented")
}
func (m MockLogger) Debug(d logger.DebugScope, v ...interface{}) {
	panic("unimplemented")
}
func (m MockLogger) Verbose(d logger.DebugScope, format string, v ...interface{}) {
	m.T.Logf(format, v...)
}
func (m MockLogger) Fatal(err error) {
	panic("unimplemented")
}
func (m MockLogger) Warn(s string, v ...interface{}) {
	panic("unimplemented")
}
