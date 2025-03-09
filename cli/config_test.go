package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	_ "github.com/aicraftlab-dev/aicraft-cli/providers/deepseek"
	"github.com/aicraftlab-dev/aicraft-cli/providers/ollama"
	_ "github.com/aicraftlab-dev/aicraft-cli/providers/ollama"
	_ "github.com/aicraftlab-dev/aicraft-cli/providers/openai"
	"github.com/aicraftlab-dev/aicraft-cli/types"
)

var ProvidersDir = "providers"
var configPath = filepath.Join(os.Getenv("HOME"), ".aicraft", "config.yaml")

// Config is the main configuration struct.
type Config = types.Config

type Provider struct {
	Name string
	URL  string
}

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure AI providers",
		Long:  `Set up and manage AI provider configurations`,
	}

	providerCmd := &cobra.Command{
		Use:   "provider",
		Short: "Set up an AI provider",
		Run:   runProvider,
	}
	agentCmd := &cobra.Command{
		Use:   "agent",
		Short: "Set up an AI agent",
		Run:   runAgent,
	}

	cmd.AddCommand(providerCmd)
	cmd.AddCommand(agentCmd)
	return cmd
}

func runProvider(cmd *cobra.Command, args []string) {
	config := loadConfig(configPath)
	providers, err := loadProviders(ProvidersDir)
	if err != nil {
		fmt.Printf("Error loading providers: %v\n", err)
		return
	}

	provider, err := selectProvider(providers)
	if err != nil {
		fmt.Printf("Error selecting provider: %v\n", err)
		return
	}

	err = openURL(provider.URL)
	if err != nil {
		fmt.Printf("Error opening provider URL: %v\n", err)
	}

	var apiKey string
	if provider.Name != "ollama" {
		apiKey, err = promptForKey()
		if err != nil {
			fmt.Printf("Error getting API key: %v\n", err)
			return
		}
	}

	var ollamaHost string
	if provider.Name == "ollama" {
		ollamaHost, err = promptForOllamaHost()
		if err != nil {
			fmt.Printf("Error getting Ollama host: %v\n", err)
			return
		}
	}

	defaultName := fmt.Sprintf("%s-config", provider.Name)
	configName, err := promptForConfigName(defaultName)
	if err != nil {
		fmt.Printf("Error getting config name: %v\n", err)
		return
	}
	var providerConfig types.ProviderConfig

	if provider.Name == "ollama" {
		providerConfig = types.ProviderConfig{
			Name:   configName,
			APIKey: "", // Ollama doesn't use API keys
			Host:   ollamaHost,
		}
	} else {
		providerConfig = types.ProviderConfig{
			Name:   configName,
			APIKey: apiKey,
			Host:   provider.URL,
		}
	}

	if config.Providers == nil {
		config.Providers = make(map[string][]types.ProviderConfig)
	}
	var isPresent bool
	for i, configProvider := range config.Providers[provider.Name] {
		if configProvider.Name == providerConfig.Name {
			isPresent = true
			updated := false
			if provider.Name == "ollama" && configProvider.Host != providerConfig.Host {
				updated = true
			}
			if provider.Name != "ollama" && configProvider.APIKey != providerConfig.APIKey {
				updated = true
			}
			if updated {
				update, err := promptForUpdate()
				if err != nil {
					fmt.Printf("Error getting provider update option: %v\n", err)
					return
				}
				if !update {
					configName, err := promptForConfigName(defaultName)
					if err != nil {
						fmt.Printf("Error getting config name: %v\n", err)
						return
					}
					providerConfig.Name = configName
				}
			} else {
				fmt.Printf("The config name %s already exists.\n", providerConfig.Name)
				return
			}
			config.Providers[provider.Name][i].APIKey = providerConfig.APIKey
			config.Providers[provider.Name][i].Host = providerConfig.Host
			config.Providers[provider.Name][i].Name = providerConfig.Name
			fmt.Printf("The config name %s was updated.\n", providerConfig.Name)
		}
	}
	if !isPresent {
		config.Providers[provider.Name] = append(config.Providers[provider.Name], providerConfig)
		fmt.Printf("The config name %s was added.\n", providerConfig.Name)
	}
	saveConfig(configPath, config)
}

func runAgent(cmd *cobra.Command, args []string) {
	config := loadConfig(configPath)
	if len(config.Providers) == 0 {
		fmt.Println("No providers configured. Run `aicraft config provider` first.")
		os.Exit(1)
	}

	selectedProviderConfig, err := selectProviderConfig(config)
	if err != nil {
		fmt.Printf("Error selecting provider: %v\n", err)
		return
	}
	var models []string
	//Now we use the name instead of selectedProviderConfig.Name
	providerName := getProviderName(selectedProviderConfig.Name)
	provider, ok := types.Providers[providerName]
	if !ok {
		fmt.Printf("Unknown provider: %s\n", providerName)
		return
	}
	if !isOllamaConfig(selectedProviderConfig.Name) && selectedProviderConfig.Name != "" {

		models, err = provider.GetModels(selectedProviderConfig.APIKey, "")
		if err != nil {
			fmt.Printf("Error getting models: %v\n", err)
			return
		}

	} else {
		models, err = getOllamaModels(selectedProviderConfig.Host)
		if err != nil {
			fmt.Printf("Error getting models: %v\n", err)
			return
		}
	}

	selectedModel, err := selectModel(models)
	if err != nil {
		fmt.Printf("Error selecting model: %v\n", err)
		return
	}
	agentName, err := promptForAgentName()
	if err != nil {
		fmt.Printf("Error getting agent name: %v\n", err)
		return
	}
	config.Agents = append(config.Agents, types.Agent{
		Name:               agentName,
		Model:              selectedModel,
		ProviderConfigName: selectedProviderConfig.Name,
	})
	saveConfig(configPath, config)

}

func getOllamaModels(host string) ([]string, error) {
	provider := ollama.OllamaProvider{}
	return provider.GetModels("", host)
}
