package main

import (
	"harmony-server/server"
	"os"
)

func main() {
	_ = os.Mkdir("./filestore", 0777)
	s := new(server.Instance)
	s.Start()
}
