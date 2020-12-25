package data

import "net/http"

//go:generate esc -o data_gen.go -pkg data ../../sql

// AssertFile asserts that the given operation succeeds
func AssertFile(file http.File, err error) http.File {
	if err != nil {
		panic(err)
	}
	return file
}

// AssertByteArray asserts that the given operation succeeds
func AssertByteArray(d []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return d
}
