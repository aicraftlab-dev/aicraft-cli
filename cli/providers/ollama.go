package providers

import (
    "fmt"
    "strings"
)

type ProviderConfig struct {
    APIKey string
    Model  string
}

type LLMProvider struct {
    Name    string
    Setup   func() error
    Connect func() error
}

var OllamaProvider = LLMProvider{
    Name: "Local Ollama",
    Setup: func() error {
        fmt.Println("To use local Ollama, you'll need to:")
        fmt.Println("1. Install Ollama: https://github.com/jmorganca/ollama")
        fmt.Println("2. Run a model locally using: ollama run <model-name>")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter Ollama server address (default: http://localhost:11434): ")
        var address string
        fmt.Scanln(&address)
        address = strings.TrimSpace(address)
        
        if address == "" {
            address = "http://localhost:11434"
        }
        
        // Store connection in config
        config := loadConfig(configPath)
        config.AI.Providers["ollama"] = ProviderConfig{
            APIKey: address,
            Model:  "llama2",
        }
        saveConfig(configPath, config)
        
        fmt.Println("Ollama connection successfully configured!")
        return nil
    },
}