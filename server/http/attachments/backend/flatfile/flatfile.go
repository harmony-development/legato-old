package flatfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

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
	Size        int32
}

// Serialize serializes the file data into a byte array
func (f FileData) Serialize() []byte {
	data, _ := json.Marshal(f)
	return data
}

// SaveFile saves file
func (b *Backend) SaveFile(name, contentType string, r io.Reader) (id string, err error) {
	fileID := uuid.New().String()

	filedata, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filepath.Join(b.Config.Server.FlatfileMediaPath, fileID), filedata, 0o660)
	if err != nil {
		return "", err
	}

	data := FileData{
		ContentType: contentType,
		Filename:    name,
		Size:        int32(len(filedata)),
	}

	err = ioutil.WriteFile(filepath.Join(b.Config.Server.FlatfileMediaPath, fmt.Sprintf("%s.data", fileID)), data.Serialize(), 0o660)
	if err != nil {
		return "", err
	}

	return fileID, nil
}

// ReadFile readsfile
func (b *Backend) ReadFile(id string) (contentType, filename string, size int32, r io.ReadCloser, err error) {
	baseFileName := filepath.Base(id)
	data, err := ioutil.ReadFile(filepath.Join(b.Config.Server.FlatfileMediaPath, fmt.Sprintf("%s.data", baseFileName)))
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
	size = fileData.Size

	r, err = os.Open(path.Join(b.Config.Server.FlatfileMediaPath, baseFileName))

	return
}

func (b *Backend) GetMetadata(id string) (contentType, fileName string, size int32, err error) {
	baseFileName := filepath.Base(id)
	data, err := ioutil.ReadFile(filepath.Join(b.Config.Server.FlatfileMediaPath, fmt.Sprintf("%s.data", baseFileName)))
	if err != nil {
		return
	}

	fileData := FileData{}
	err = json.Unmarshal(data, &fileData)
	if err != nil {
		return
	}

	return fileData.ContentType, fileData.Filename, fileData.Size, err
}
