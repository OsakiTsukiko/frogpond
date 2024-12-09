package main

import (
	"log"

	"github.com/OsakiTsukiko/frogpond/server"
	"github.com/OsakiTsukiko/frogpond/server/singleton"
)

func main() {
	log.Println("🐸 The Pond is Forming!")
	log.Printf("🐸 Running on domain: %q\n", singleton.CFG.Server.Domain)
	server.Run()
}
