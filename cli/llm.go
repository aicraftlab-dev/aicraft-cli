package cli

import (
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

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
    
    "anthropic": {
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
    },
    "google": {
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
    },
    "cohere": {
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
    },
    "huggingface": {
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
    },
    "ollama": {
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
    "anthropic": {
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
    },
    "google": {
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
    },
    "cohere": {
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
    },
    "huggingface": {
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
    },
}
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
    "ollama": {
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
    },
}

func NewLLMCmd() *cobra.Command {
    var provider string
    
    cmd := &cobra.Command{
        Use:   "llm",
        Short: "Configure LLM providers",
        Long:  `Set up and connect to different LLM providers`,
    }
    
    setupCmd := &cobra.Command{
        Use:   "setup",
        Short: "Setup instructions for LLM provider",
        Run: func(cmd *cobra.Command, args []string) {
            p, exists := providers[provider]
            if !exists {
                fmt.Printf("Unknown provider: %s\n", provider)
                return
            }
            
            if err := p.Setup(); err != nil {
                fmt.Printf("Error: %v\n", err)
                os.Exit(1)
            }
        },
    }
    
    connectCmd := &cobra.Command{
        Use:   "connect",
        Short: "Connect to LLM provider",
        Run: func(cmd *cobra.Command, args []string) {
            p, exists := providers[provider]
            if !exists {
                fmt.Printf("Unknown provider: %s\n", provider)
                return
            }
            
            if err := p.Connect(); err != nil {
                fmt.Printf("Error: %v\n", err)
                os.Exit(1)
            }
        },
    }
    
    setupCmd.Flags().StringVarP(&provider, "provider", "p", "", "LLM provider name")
    connectCmd.Flags().StringVarP(&provider, "provider", "p", "", "LLM provider name")
    
    cmd.AddCommand(setupCmd)
    cmd.AddCommand(connectCmd)
    return cmd
}