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

var GoogleProvider = LLMProvider{
    Name: "Google Gemini",
    Setup: func() error {
        fmt.Println("Please create a Google Cloud account at:")
        fmt.Println("https://cloud.google.com/")
        fmt.Println("\nThen enable the Gemini API and generate an API token at:")
        fmt.Println("https://console.cloud.google.com/apis/credentials")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter your Google API token: ")
        var token string
        fmt.Scanln(&token)
        token = strings.TrimSpace(token)
        
        if token == "" {
            return fmt.Errorf("API token cannot be empty")
        }
        
        // Store token in config
        config := loadConfig(configPath)
        config.AI.Providers["google"] = ProviderConfig{
            APIKey: token,
            Model:  "gemini-pro",
        }
        saveConfig(configPath, config)
        
        fmt.Println("Google API token successfully configured!")
        return nil
    },
}