package database_attachments_backend

import (
	"crypto/sha1"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
)

// Dependencies lists the deps of the flatfile backend
type Dependencies struct {
	Config *config.Config
	DB     db.IHarmonyDB
}

// Backend is the flatfile implementation of a file storage backend
type Backend struct {
	Dependencies
}

// SaveFile saves file
func (b *Backend) SaveFile(name, contentType string, r io.Reader) (id string, err error) {
	fileID := uuid.New().String()

	filedata, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filepath.Join(b.Config.Flatfile.MediaPath, fileID), filedata, 0o660)
	if err != nil {
		return "", err
	}

	err = b.DB.SetFileMetadata(fileID, contentType, name, int32(len(filedata)))
	if err != nil {
		return "", err
	}
	hash := sha1.Sum(filedata)
	err = b.DB.AddFileHash(fileID, hash[:]) // any way to use an array directly?

	return fileID, nil
}

// ReadFile readsfile
func (b *Backend) ReadFile(id string) (contentType, fileName string, size int32, r io.ReadCloser, err error) {
	baseFileName := filepath.Base(id)
	res, err := b.DB.GetFileMetadata(id)
	if err != nil {
		return
	}

	r, err = os.Open(path.Join(b.Config.Flatfile.MediaPath, baseFileName))

	return res.ContentType, res.Name, res.Size, r, err
}

func (b *Backend) GetMetadata(id string) (contentType, fileName string, size int32, err error) {
	res, err := b.DB.GetFileMetadata(id)
	return res.ContentType, res.Name, res.Size, err
}
