package cli

import (
	"fmt"
	"os"

	"github.com/aicraftlab-dev/aicraft-cli/types"
	"github.com/spf13/cobra"
)

func NewLLMCmd() *cobra.Command {
	var prompt string
	var configName string //Changed from providerName to configName
	var modelName string

	cmd := &cobra.Command{
		Use:   "llm",
		Short: "Interact with LLM providers",
		Long:  `Send prompts to LLM providers and receive responses`,
		Run: func(cmd *cobra.Command, args []string) {
			config := loadConfig(configPath)
			if config.Providers == nil {
				fmt.Println("No provider configured. Run `aicraft config setup` first.")
				os.Exit(1)
			}

			var selectedProviderConfig types.ProviderConfig
			var providerFound bool
			var provider types.Provider
			//Find the provider in the list of providers.
			for _, providerConfig := range config.Providers {
				if providerConfig.Name == configName {
					selectedProviderConfig = providerConfig
					// Now we use the provider name to look in the types.Providers map.
					var ok bool
					provider, ok = types.Providers[providerConfig.Provider]
					if !ok {
						fmt.Printf("Unknown provider: %s\n", providerConfig.Provider)
						os.Exit(1)
					}
					providerFound = true
					break
				}
			}
			if !providerFound {
				fmt.Printf("Configuration %s not found. Run `aicraft config setup`.\n", configName)
				os.Exit(1)
			}

			var host string
			if isOllamaConfig(selectedProviderConfig.Name) { //Now we use the name instead of provider.
				host = selectedProviderConfig.Host
			}

			response, err := provider.Generate(modelName, prompt, selectedProviderConfig.APIKey, host)
			if err != nil {
				fmt.Printf("Error generating response: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(response)
		},
	}

	cmd.Flags().StringVarP(&configName, "config", "c", "", "Configuration name") //Change from provider to config
	cmd.Flags().StringVarP(&modelName, "model", "m", "", "LLM model name")
	cmd.Flags().StringVarP(&prompt, "prompt", "q", "", "Prompt to send to the model")
	cmd.MarkFlagRequired("config") //changed from provider to config
	cmd.MarkFlagRequired("model")
	cmd.MarkFlagRequired("prompt")

	return cmd
}
