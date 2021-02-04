package main

import (
	"flag"

	"github.com/harmony-development/legato/cmd"
	"github.com/harmony-development/legato/server"

	"github.com/sirupsen/logrus"

	_ "github.com/harmony-development/legato/server/db/backends/postgres"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	var genKey bool
	flag.BoolVar(&genKey, "genkey", false, "generates a key pair for federation")
	flag.BoolVar(&genKey, "g", false, "generates a key pair for federation")
	var genData bool
	flag.BoolVar(&genData, "gendata", false, "generates testing data")
	flag.BoolVar(&genData, "d", false, "generates testing data")
	flag.Parse()

	if genKey {
		cmd.GenKeys()
		return
	}
	if genData {
		cmd.GenData()
		return
	}

	logrus.Info("Server starting")
	s := new(server.Instance)
	s.Start()
}
