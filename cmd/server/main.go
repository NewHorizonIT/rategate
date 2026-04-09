package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NewHorizonIT/rategate/internal/config"
	"github.com/NewHorizonIT/rategate/internal/infra/redis"
	"github.com/NewHorizonIT/rategate/internal/server"
)

func main() {
	fmt.Println("Hello, Gateway!")
	// Load configuration
	cfg := config.SetupConfig()

	// Log Configuration
	log.Printf("Configuration loaded: %+v", cfg)

	// Connect to redis
	redisClient := redis.NewClient(cfg.Redis)
	// Test Redis connection
	if err := redis.Ping(context.Background(), redisClient); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("[[Successfully connected to Redis]]")

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
