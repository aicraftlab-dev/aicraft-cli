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