package flatfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/harmony-development/legato/server/config"
)

// Dependencies lists the deps of the flatfile backend
type Dependencies struct {
	Config *config.Config
}

// Backend is the flatfile implementation of a file storage backend
type Backend struct {
	Dependencies
}

// FileData represents the data stored by a file
type FileData struct {
	ContentType string
	Filename    string
}

// Serialize serializes the file data into a byte array
func (f FileData) Serialize() []byte {
	data, _ := json.Marshal(f)
	return data
}

// SaveFile saves file
func (b *Backend) SaveFile(name, contentType string, r io.Reader) (id string, err error) {
	data := FileData{
		ContentType: contentType,
		Filename:    name,
	}
	fileID := uuid.New().String()

	filedata, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(path.Join(b.Config.Server.FlatfileMediaPath, fileID), filedata, 0o660)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(path.Join(b.Config.Server.FlatfileMediaPath, fmt.Sprintf("%s.data", fileID)), data.Serialize(), 0o660)
	if err != nil {
		return "", err
	}

	return fileID, nil
}

// ReadFile readsfile
func (b *Backend) ReadFile(id string) (contentType, filename string, r io.ReadCloser, err error) {
	data, err := ioutil.ReadFile(path.Join(b.Config.Server.FlatfileMediaPath, fmt.Sprintf("%s.data", id)))
	if err != nil {
		return
	}

	fileData := FileData{}
	err = json.Unmarshal(data, &fileData)
	if err != nil {
		return
	}

	contentType = fileData.ContentType
	filename = fileData.Filename

	r, err = os.Open(path.Join(b.Config.Server.FlatfileMediaPath, id))

	return
}
