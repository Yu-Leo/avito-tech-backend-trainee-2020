package main

import (
	"log"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/app"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	l := logger.NewLogger(cfg.Logger.Level)

	app.Run(cfg, l)
}
