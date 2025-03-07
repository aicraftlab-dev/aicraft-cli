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

var AnthropicProvider = LLMProvider{
    Name: "Anthropic (Claude)",
    Setup: func() error {
        fmt.Println("Please create an Anthropic API account at:")
        fmt.Println("https://www.anthropic.com/signup")
        fmt.Println("\nThen generate an API token at:")
        fmt.Println("https://console.anthropic.com/settings/keys")
        return nil
    },
    Connect: func() error {
        fmt.Print("Enter your Anthropic API token: ")
        var token string
        fmt.Scanln(&token)
        token = strings.TrimSpace(token)
        
        if token == "" {
            return fmt.Errorf("API token cannot be empty")
        }
        
        // Store token in config
        config := loadConfig(configPath)
        config.AI.Providers["anthropic"] = ProviderConfig{
            APIKey: token,
            Model:  "claude-2",
        }
        saveConfig(configPath, config)
        
        fmt.Println("Anthropic API token successfully configured!")
        return nil
    },
}