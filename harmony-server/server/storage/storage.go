package storage

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/getsentry/sentry-go"
)

// Manager manages storage operations
type Manager struct {
	ImageDeleteQueue        chan string
	GuildPictureDeleteQueue chan string
	ImagePath               string
	GuildPicturePath        string
}

// DeleteImage adds an image to delete to the queue
func (m Manager) DeleteImage(id string) {
	m.ImageDeleteQueue <- id
}

// DeleteGuildPicture adds a guild picture to delete to the queue
func (m Manager) DeleteGuildPicture(id string) {
	m.GuildPictureDeleteQueue <- id
}

// DeleteRoutine is a function that deletes images that are being queued
func (m Manager) DeleteRoutine() {
	for {
		select {
		case target := <-m.ImageDeleteQueue:
			{
				if err := os.Remove(path.Join(m.ImagePath, target)); err != nil {
					sentry.CaptureException(err)
				}
			}
		case target := <-m.GuildPictureDeleteQueue:
			{
				if err := os.Remove(path.Join(m.GuildPicturePath, target)); err != nil {
					sentry.CaptureException(err)
				}
			}
		}
	}
}

// AddImage writes an image to storage
func (m Manager) AddImage(id string, image []byte) error {
	return ioutil.WriteFile(path.Join(m.ImagePath, id), image, 0666)
}

// AddGuildPicture adds an image for a guild to storage
func (m Manager) AddGuildPicture(id string, image []byte) error {
	return ioutil.WriteFile(path.Join(m.GuildPicturePath, id), image, 0666)
}
