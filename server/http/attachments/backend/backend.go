package backend

import (
	"errors"
	"io"
)

var NotFound = errors.New("Not Found")

type AttachmentBackend interface {
	SaveFile(name, contentType string, r io.Reader) (id string, err error)
	ReadFile(id string) (contentType, filename string, r io.ReadCloser, err error)
}
