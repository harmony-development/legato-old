package main

import (
	"flag"

	"github.com/harmony-development/legato/cmd"
	"github.com/harmony-development/legato/server"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	var genKey bool
	flag.BoolVar(&genKey, "genkey", false, "generates a key pair for federation")
	flag.BoolVar(&genKey, "g", false, "generates a key pair for federation")
	flag.Parse()

	if genKey {
		cmd.GenKeys()
		return
	}
	logrus.Info("Server starting")
	s := new(server.Instance)
	s.Start()
}
