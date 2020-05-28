package main

import (
	"harmony-server/server"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	_ = os.Mkdir("./filestore", 0777)
	logrus.SetLevel(logrus.DebugLevel)
	s := new(server.Instance)
	s.Start()
}
