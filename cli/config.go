package cli

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"

    "github.com/aicraftlab-dev/aicraft-cli/types"
)

type Config struct {
    AI struct {
        Providers map[string]types.ProviderConfig `yaml:"providers"`
    } `yaml:"ai"`
}

func NewConfigCmd() *cobra.Command {
    var provider string
    var apiKey string
    var model string

    cmd := &cobra.Command{
        Use:   "config",
        Short: "Configure AI providers",
        Long:  `Set up and manage AI provider configurations`,
    }

    setupCmd := &cobra.Command{
        Use:   "setup",
        Short: "Set up an AI provider",
        Run: func(cmd *cobra.Command, args []string) {
            configPath := filepath.Join(os.Getenv("HOME"), ".aicraft", "config.yaml")
            config := loadConfig(configPath)
            
            if config.AI.Providers == nil {
                config.AI.Providers = make(map[string]types.ProviderConfig)
            }

            config.AI.Providers[provider] = types.ProviderConfig{
                APIKey: apiKey,
                Model:  model,
            }

            saveConfig(configPath, config)
            fmt.Printf("Successfully configured %s provider\n", provider)
        },
    }

    setupCmd.Flags().StringVarP(&provider, "provider", "p", "", "AI provider name (e.g., deepseek)")
    setupCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "API key for the provider")
    setupCmd.Flags().StringVarP(&model, "model", "m", "deepseek-chat", "Default model to use")
    
    setupCmd.MarkFlagRequired("provider")
    setupCmd.MarkFlagRequired("api-key")

    cmd.AddCommand(setupCmd)
    return cmd
}

func loadConfig(path string) Config {
    config := Config{}
    
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return config
    }

    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Printf("Error reading config: %v\n", err)
        return config
    }

    if err := yaml.Unmarshal(data, &config); err != nil {
        fmt.Printf("Error parsing config: %v\n", err)
    }

    return config
}

func saveConfig(path string, config Config) {
    data, err := yaml.Marshal(&config)
    if err != nil {
        fmt.Printf("Error marshaling config: %v\n", err)
        return
    }

    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0755); err != nil {
        fmt.Printf("Error creating config directory %s: %v\n", dir, err)
        return
    }

    if err := os.WriteFile(path, data, 0600); err != nil {
        fmt.Printf("Error writing config to %s: %v\n", path, err)
    } else {
        fmt.Printf("Successfully wrote config to %s\n", path)
    }
}