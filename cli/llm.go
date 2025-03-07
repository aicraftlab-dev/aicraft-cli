package cli

import (
    "fmt"
    "os"
    "path/filepath"
    "plugin"

    "github.com/spf13/cobra"
)

func loadProviders() error {
    // Get the absolute path to the providers directory
    providersDir := filepath.Join(filepath.Dir(os.Args[0]), "providers")
    
    // Read all .so files in the providers directory
    files, err := filepath.Glob(filepath.Join(providersDir, "*.so"))
    if err != nil {
        return fmt.Errorf("failed to read providers directory: %v", err)
    }

    // Load each provider plugin
    for _, file := range files {
        p, err := plugin.Open(file)
        if err != nil {
            return fmt.Errorf("failed to load provider plugin %s: %v", file, err)
        }

        // Look up the provider symbol
        providerSymbol, err := p.Lookup("Provider")
        if err != nil {
            return fmt.Errorf("failed to lookup provider symbol in %s: %v", file, err)
        }

        // Cast the symbol to LLMProvider
        provider, ok := providerSymbol.(*LLMProvider)
        if !ok {
            return fmt.Errorf("invalid provider type in %s", file)
        }

        // Add the provider to the providers map
        providers[provider.Name] = *provider
    }

    return nil
}

func NewLLMCmd() *cobra.Command {
    var provider string

    // Load providers dynamically
    if err := loadProviders(); err != nil {
        fmt.Printf("Error loading providers: %v\n", err)
        os.Exit(1)
    }

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