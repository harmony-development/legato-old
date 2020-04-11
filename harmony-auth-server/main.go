package main

import (
	"harmony-auth-server/harmony"
	"os"
)

func main() {
	_ = os.Mkdir("./avatars", 0777)
	s := new(harmony.Instance)
	s.Start()
}
