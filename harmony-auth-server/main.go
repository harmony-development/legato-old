package main

import (
	"harmony-auth-server/server"
	"os"
)

func main() {
	_ = os.Mkdir("./avatars", 0777)
	s := new(server.Instance)
	s.Start()
}
