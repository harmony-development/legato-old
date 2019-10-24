package main

import (
	"log"
)

func main() {
	port := 3000
	startServer(port, func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	})
}
