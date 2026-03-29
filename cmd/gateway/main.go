package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NewHorizonIT/rategate/internal/config"
	"github.com/NewHorizonIT/rategate/internal/server"
)

func main() {
	fmt.Println("Hello, Gateway!")
	// Load configuration
	cfg := config.SetupConfig()

	// Log Configuration
	log.Printf("Configuration loaded: %+v", cfg)

	// Setup route health check
	r := server.SetupRoutes()

	// Initialize and start the gateway
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Gateway is running on port ", cfg.Server.Port)

}
