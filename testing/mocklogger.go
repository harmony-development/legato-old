package testing

import "fmt"

type MockLogger struct{}

// CheckException logs an exception if it's defined
func (l MockLogger) CheckException(err error) {
	if err == nil {
		return
	}
	l.Exception(err)
}

// Exception logs an exception
func (l MockLogger) Exception(err error) {
	fmt.Println(err)
}

// Fatal logs an exception and then aborts
func (l MockLogger) Fatal(err error) {
	fmt.Println(err)
}
