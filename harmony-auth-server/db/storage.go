package db

import (
	"fmt"
	"github.com/kataras/golog"
	"os"
)

func DeleteFromAvatars(fileid string) {
	err := os.Remove(fmt.Sprintf("./filestore/%v", fileid))
	if err != nil {
		golog.Warnf("Error deleting from filestore : %v", err)
	}
}