package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	if strings.ToLower(provider.Name) == "ollama" {
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
	providerConfig := types.ProviderConfig{
		Name:     configName,
		Provider: provider.Name,
		APIKey:   apiKey,
		Host:     provider.URL,
	}
	if strings.ToLower(provider.Name) == "ollama" {
		providerConfig.APIKey = "" // Ollama doesn't use API keys
		providerConfig.Host = ollamaHost
	}

	var isPresent bool
	for i, configProvider := range config.Providers {
		if configProvider.Name == providerConfig.Name && configProvider.Provider == providerConfig.Provider {
			isPresent = true
			updated := false
			if strings.ToLower(provider.Name) == "ollama" && configProvider.Host != providerConfig.Host {
				updated = true
			}
			if strings.ToLower(provider.Name) != "ollama" && configProvider.APIKey != providerConfig.APIKey {
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
				fmt.Printf("The config %s for %s already exists.\n", providerConfig.Name, providerConfig.Provider)
				return
			}
			config.Providers[i].APIKey = providerConfig.APIKey
			config.Providers[i].Host = providerConfig.Host
			config.Providers[i].Name = providerConfig.Name
			fmt.Printf("The config name %s was updated.\n", providerConfig.Name)
			break //Exit from the loop.
		}
	}
	if !isPresent {
		config.Providers = append(config.Providers, providerConfig)
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
	//Find the provider
	var provider types.Provider
	var ok bool
	for _, configProvider := range config.Providers {
		if configProvider.Name == selectedProviderConfig.Name {
			provider, ok = types.Providers[configProvider.Provider]
			selectedProviderConfig = configProvider
			break
		}
	}

	if !ok {
		fmt.Printf("Unknown provider: %s\n", selectedProviderConfig.Name)
		return
	}

	if !isOllamaConfig(selectedProviderConfig.Name) {
		//Use the host and apikey from selectedProviderConfig
		models, err = provider.GetModels(selectedProviderConfig.APIKey, selectedProviderConfig.Host)
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

	for {
		agentName, err := promptForAgentName(config)
		if err != nil {
			fmt.Printf("Error getting agent name: %v\n", err)
			return
		}
		var agentToUpdate *types.Agent
		var agentIndex int
		for i, agent := range config.Agents {
			if agent.Name == agentName {
				agentToUpdate = &agent
				agentIndex = i
				break
			}
		}

		if agentToUpdate != nil {
			update, err := promptForUpdate()
			if err != nil {
				fmt.Printf("Error getting agent update option: %v\n", err)
				return
			}

			if update {
				config.Agents[agentIndex].Model = selectedModel
				config.Agents[agentIndex].ProviderConfigName = selectedProviderConfig.Name
				config.Agents[agentIndex].Name = agentName
				fmt.Printf("The agent %s was updated.\n", agentName)
				saveConfig(configPath, config)
				return
			}
		} else {
			config.Agents = append(config.Agents, types.Agent{
				Name:               agentName,
				Model:              selectedModel,
				ProviderConfigName: selectedProviderConfig.Name,
			})
			fmt.Printf("The agent %s was created.\n", agentName)
			saveConfig(configPath, config)
			return
		}
	}
}

func getOllamaModels(host string) ([]string, error) {
	provider := ollama.OllamaProvider{}
	return provider.GetModels("", host)
}
