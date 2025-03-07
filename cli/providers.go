package cli

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

var providers = map[string]LLMProvider{
    "deepseek": {
        Name: "Deepseek",
        Setup: func() error {
            fmt.Println("Please create a Deepseek API account at:")
            fmt.Println("https://platform.deepseek.com/signup")
            fmt.Println("\nThen generate an API token at:")
            fmt.Println("https://platform.deepseek.com/api-keys")
            return nil
        },
        Connect: func() error {
            fmt.Print("Enter your Deepseek API token: ")
            var token string
            fmt.Scanln(&token)
            token = strings.TrimSpace(token)
            
            if token == "" {
                return fmt.Errorf("API token cannot be empty")
            }
            
            // Store token in config
            config := loadConfig(configPath)
            config.AI.Providers["deepseek"] = ProviderConfig{
                APIKey: token,
                Model:  "deepseek-chat",
            }
            saveConfig(configPath, config)
            
            fmt.Println("Deepseek API token successfully configured!")
            return nil
        },
    },
    "openai": {
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
    },
    // Add other providers here...
}