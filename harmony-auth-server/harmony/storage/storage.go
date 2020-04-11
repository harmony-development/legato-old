package storage

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type Manager struct {
	DeleteQueue chan string
}

// New creates a new storage manager
func New() *Manager {
	return &Manager{
		DeleteQueue: make(chan string, 128),
	}
}

// DeleteAvatar adds an avatar to delete to the queue
func (h Manager) DeleteAvatar(id string) {
	h.DeleteQueue <- id
}

// DeleteRoutine is a function that deletes avatars that are being queued
func (h Manager) DeleteRoutine(avatarPath string) {
	for {
		id := <-h.DeleteQueue
		if err := os.Remove(path.Join(avatarPath, id)); err != nil {
			logrus.Warn("Error deleting avatar", id)
		}
	}
}
