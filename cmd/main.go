package main

import (
	"log"

	"github.com/ij4l/go-htmx/apps"
	"github.com/ij4l/go-htmx/internal/config"
)

func main() {
	config, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Cannot load configuration: %v", err)
	}

	server, err := apps.NewServer()
	if err != nil {
		log.Fatalf("Cannot create server: %v", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
