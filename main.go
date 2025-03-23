package main

import (
	"fmt"

	"./config"
	"./services"
)

func main() {
	cfg := config.LoadConfig()

	service := services.NewService(cfg)

	if err := service.ProcessData(); err != nil {
		fmt.Printf("Error processing data: %v\n", err)
		return
	}
}
