package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// tag config for envconfig
// require binding doesnt work
// in env config sadly
type Config struct {
	Server struct {
		Port string `envconfig:"FP_PORT"`
		// TODO: ADD SECURE PORT and NOT SECURE PORT
		JWTSecretKey string `envconfig:"FP_JWT_SECRET_KEY"`
		Domain       string `envconfig:"FP_DOMAIN"`

		DefaultRedirect string
		UseHTTPS        bool
	}

	DataBase struct {
		Host     string `envconfig:"FP_DB_HOST"`
		Port     string `envconfig:"FP_DB_PORT"`
		Username string `envconfig:"FP_DB_USERNAME"`
		Password string `envconfig:"FP_DB_PASSWORD"`
		Database string `envconfig:"FP_DB_DATABASE"`
	}
}

func LoadConfig() Config {
	var cfg Config
	// load config from environment based on tags
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("🚩 Error loading configuration: %v", err)
	}

	cfg.Server.DefaultRedirect = "/"

	// check if the fields are set
	// TODO: AUTOMATE THIS (maybe)
	if cfg.Server.Port == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_PORT")
	}

	if cfg.Server.JWTSecretKey == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_JWT_SECRET_KEY")
	}

	if cfg.Server.Domain == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DOMAIN")
	}

	if cfg.DataBase.Host == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DB_HOST")
	}

	if cfg.DataBase.Port == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DB_PORT")
	}

	if cfg.DataBase.Username == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DB_USERNAME")
	}

	if cfg.DataBase.Password == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DB_PASSWORD")
	}

	if cfg.DataBase.Database == "" {
		log.Fatalf("🚩 Environment variable %q is empty!", "FP_DB_DATABASE")
	}

	if os.Getenv("FP_USE_HTTPS") == "true" {
		cfg.Server.UseHTTPS = true
	} else if os.Getenv("FP_USE_HTTPS") == "false" {
		cfg.Server.UseHTTPS = false
	} else {
		log.Fatalf("🚩 Environment variable %q is empty or invalid! (true/false)", "FP_USE_HTTPS")
	}

	return cfg
}
