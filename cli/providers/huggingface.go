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

var HuggingFaceProvider = LLMProvider{
    Name: "Hugging Face",
    Setup: func() error {
        fmt.Println("Please create a Hugging Face account at:")
        fmt.Println("https://huggingface.co/join")
        fmt.Println("\nThen generate an API token at:")
        fmt.Println("https://huggingface.co/settings/tokens")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter your Hugging Face API token: ")
        var token string
        fmt.Scanln(&token)
        token = strings.TrimSpace(token)
        
        if token == "" {
            return fmt.Errorf("API token cannot be empty")
        }
        
        // Store token in config
        config := loadConfig(configPath)
        config.AI.Providers["huggingface"] = ProviderConfig{
            APIKey: token,
            Model:  "mistral-7b",
        }
        saveConfig(configPath, config)
        
        fmt.Println("Hugging Face API token successfully configured!")
        return nil
    },
}