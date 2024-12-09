package main

import (
	"log"

	"github.com/OsakiTsukiko/frogpond/server"
	"github.com/OsakiTsukiko/frogpond/server/singleton"
)

func main() {
	log.Println("🐸 The Pond is Forming!")
	log.Printf("🐸 Running on domain: %q and port: %q\n", singleton.CFG.Server.Domain, singleton.CFG.Server.Port)
	server.Run() // run server
}
