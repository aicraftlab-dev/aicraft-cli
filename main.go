package main

import (
	"fmt"

	"github.com/aicraftlab-dev/aicraft-cli/config"
	"github.com/aicraftlab-dev/aicraft-cli/services"
)

func main() {
	cfg := config.LoadConfig()

	service := services.NewService(cfg)

	if err := service.ProcessData(); err != nil {
		fmt.Printf("Error processing data: %v\n", err)
		return
	}
}
