package main

import (
	"log"

	"github.com/mic-paytrix/modlib/server"
)

func main() {
	log.Print("Starting server")
	server.GinServer()
}
