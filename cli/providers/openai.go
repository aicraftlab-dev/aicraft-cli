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

var OpenAIProvider = LLMProvider{
    Name: "OpenAI",
    Setup: func() error {
        fmt.Println("Please create an OpenAI API account at:")
        fmt.Println("https://platform.openai.com/signup")
        fmt.Println("\nThen generate an API token at:")
        fmt.Println("https://platform.openai.com/api-keys")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter your OpenAI API token: ")
        var token string
        fmt.Scanln(&token)
        token = strings.TrimSpace(token)
        
        if token == "" {
            return fmt.Errorf("API token cannot be empty")
        }
        
        // Store token in config
        config := loadConfig(configPath)
        config.AI.Providers["openai"] = ProviderConfig{
            APIKey: token,
            Model:  "gpt-4",
        }
        saveConfig(configPath, config)
        
        fmt.Println("OpenAI API token successfully configured!")
        return nil
    },
}