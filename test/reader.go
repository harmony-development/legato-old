package test

import "errors"

type ErrReader int

func (ErrReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
