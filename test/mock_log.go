package test

import "fmt"

type MockLogger struct {
	LogCB func(string)
}

func (m *MockLogger) Log(v ...interface{}) {
	if m.LogCB != nil {
		m.LogCB(
			fmt.Sprintln(v...),
		)
	}
}

func (m *MockLogger) Logf(format string, v ...interface{}) {
	if m.LogCB != nil {
		m.LogCB(
			fmt.Sprintf(format, v...),
		)
	}
}
