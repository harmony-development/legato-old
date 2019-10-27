package handler

import (
	"log"
)

// PingHandler handles the ping socket event
func PingHandler(data interface{}) {
	log.Print("ping req received")
}
