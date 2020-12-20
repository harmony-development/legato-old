package backend

import (
	"errors"
	"io"
)

var NotFound = errors.New("Not Found")

type AttachmentBackend interface {
	SaveFile(name, contentType string, r io.Reader) (id string, err error)
	GetMetadata(id string) (contentType, fileName string, size int32, err error)
	ReadFile(id string) (contentType, filename string, size int32, r io.ReadCloser, err error)
}
