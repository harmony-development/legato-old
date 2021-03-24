package test

import (
	"bytes"
	"errors"
	"io"

	"github.com/google/uuid"
)

type Attachment struct {
	data        []byte
	contentType string
	fileName    string
	size        int32
}

type MockAttachments struct {
	files map[string]*Attachment
}

func NewMockAttachmentsBackend() *MockAttachments {
	return &MockAttachments{
		files: map[string]*Attachment{},
	}
}

func (m *MockAttachments) SaveFile(name, contentType string, r io.Reader) (id string, err error) {
	fileID := uuid.New().String()

	data, err := io.ReadAll(r)

	if err != nil {
		return "", err
	}

	m.files[fileID] = &Attachment{
		data:        data,
		contentType: contentType,
		fileName:    name,
		size:        int32(len(data)),
	}

	return fileID, nil
}

func (m *MockAttachments) GetMetadata(id string) (contentType, fileName string, size int32, err error) {
	if f, ok := m.files[id]; !ok {
		return "", "", 0, errors.New("attachment not found")
	} else {
		return f.contentType, f.fileName, f.size, nil
	}
}

func (m *MockAttachments) ReadFile(id string) (contentType, filename string, size int32, r io.ReadCloser, err error) {
	if f, ok := m.files[id]; !ok {
		return "", "", 0, nil, errors.New("attachment not found")
	} else {
		return f.contentType, f.fileName, f.size, io.NopCloser(bytes.NewReader(f.data)), nil
	}
}
