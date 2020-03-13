package util

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
)

// GetFile takes in a form and an entry to read, and returns the first file received
func GetFile(form *multipart.Form, entry string) ([]byte, error) {
	files := form.File[entry]
	if len(files) == 0 {
		return nil, errors.New("no file received")
	}
	file, err := files[0].Open()
	if err != nil {
		return nil, errors.New("error reading file")
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("error reading file")
	}
	return fileBytes, nil
}
