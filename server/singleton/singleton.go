package singleton

import (
	"log"

	"github.com/OsakiTsukiko/frogpond/server/config"
)

// global scope variables
var CFG config.Config

// init is run when the program starts
func init() {
	log.Println("🐸 Initializing FrogPond Singleton")
	CFG = config.LoadConfig() // load config from environment variables
}
