package main

import (
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
