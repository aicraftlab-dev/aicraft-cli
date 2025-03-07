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

var CohereProvider = LLMProvider{
    Name: "Cohere",
    Setup: func() error {
        fmt.Println("Please create a Cohere API account at:")
        fmt.Println("https://dashboard.cohere.com/signup")
        fmt.Println("\nThen generate an API token at:")
        fmt.Println("https://dashboard.cohere.com/api-keys")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter your Cohere API token: ")
        var token string
        fmt.Scanln(&token)
        token = strings.TrimSpace(token)
        
        if token == "" {
            return fmt.Errorf("API token cannot be empty")
        }
        
        // Store token in config
        config := loadConfig(configPath)
        config.AI.Providers["cohere"] = ProviderConfig{
            APIKey: token,
            Model:  "command",
        }
        saveConfig(configPath, config)
        
        fmt.Println("Cohere API token successfully configured!")
        return nil
    },
}