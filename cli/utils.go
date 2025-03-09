package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/aicraftlab-dev/aicraft-cli/types"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
)

func loadConfig(path string) types.Config {
	config := types.Config{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return config
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return config
	}
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		fmt.Printf("Error parsing config: %v\n", err)
		return config
	}
	// Check if the providers are in the old map format
	if providersMap, ok := rawConfig.(map[string]interface{})["providers"].(map[string]interface{}); ok {
		fmt.Println("Migrating config.yaml to the new format...")
		for providerType, configs := range providersMap {
			for _, configData := range configs.([]interface{}) {
				configMap := configData.(map[string]interface{})
				providerConfig := types.ProviderConfig{
					Provider: providerType,
					Name:     configMap["name"].(string),
					APIKey:   configMap["apiKey"].(string),
					Host:     configMap["host"].(string),
				}
				config.Providers = append(config.Providers, providerConfig)
			}
		}
		//Check if there is a agent section in the config file.
		if agents, ok := rawConfig.(map[string]interface{})["agents"].([]interface{}); ok {
			for _, agentData := range agents {
				agentMap := agentData.(map[string]interface{})
				agent := types.Agent{
					Name:               agentMap["name"].(string),
					Model:              agentMap["model"].(string),
					ProviderConfigName: agentMap["provider"].(string),
				}
				config.Agents = append(config.Agents, agent)
			}
		}

		saveConfig(path, config) // Save the migrated data
	} else {
		// If not in the old format, try to unmarshal directly to the new struct
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Printf("Error parsing config: %v\n", err)
		}
	}

	return config
}

func saveConfig(path string, config types.Config) {
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

func loadProviders(dir string) ([]Provider, error) {
	// Construct the absolute path to the providers directory
	providersDir := filepath.Join(ProvidersDir) // 'dir' will be "providers" in this case

	entries, err := os.ReadDir(providersDir)
	if err != nil {
		return nil, fmt.Errorf("error reading providers directory '%s': %v", providersDir, err)
	}

	var providers []Provider
	for _, entry := range entries {
		if entry.IsDir() {
			providerName := entry.Name()
			urlPath := filepath.Join(providersDir, providerName, "url.txt")

			if _, err := os.Stat(urlPath); os.IsNotExist(err) {
				fmt.Printf("url.txt not found for provider %s. Skipping.\n", providerName)
				continue
			}

			urlBytes, err := os.ReadFile(urlPath)
			if err != nil {
				fmt.Printf("Error reading url for provider %s: %v. Skipping.\n", providerName, err)
				continue
			}
			url := strings.TrimSpace(string(urlBytes))

			providers = append(providers, Provider{
				Name: getProviderName(providerName),
				URL:  url,
			})
		}
	}
	return providers, nil
}

func selectProvider(providers []Provider) (Provider, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ .Name | cyan }} ({{ .URL | yellow }})",
		Inactive: "  {{ .Name | cyan }} ({{ .URL | yellow }})",
		Selected: "\U000027A4 {{ .Name | red | cyan }}",
		Details: `
--------- Provider ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "URL:" | faint }}	{{ .URL }}`,
	}

	prompt := promptui.Select{
		Label:     "Select an AI provider",
		Items:     providers,
		Templates: templates,
		Size:      8,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return Provider{}, err
	}

	return providers[i], nil
}

func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return runCommand(cmd, args...)
}

func promptForKey() (string, error) {
	validate := func(input string) error {
		if len(input) < 10 {
			return fmt.Errorf("API Key must have at least 10 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter API Key",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func selectModel(models []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select a model",
		Items: models,
	}

	_, model, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return model, nil
}

func runCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	return command.Start()
}

func promptForConfigName(defaultName string) (string, error) {
	prompt := promptui.Prompt{
		Label:   "Enter a name for this configuration",
		Default: defaultName,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}

	return result, nil
}

func promptForOllamaHost() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Enter Ollama Host (optional)",
		Default: "http://localhost:11434",
	}
	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}

	return result, nil
}

func promptForAgentName(config types.Config) (string, error) {
	for {
		validate := func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("agent name cannot be empty")
			}
			return nil
		}
		prompt := promptui.Prompt{
			Label:    "Enter a name for this agent",
			Validate: validate,
		}

		result, err := prompt.Run()
		if err != nil {
			return "", fmt.Errorf("prompt failed %v", err)
		}

		for _, agent := range config.Agents {
			if agent.Name == result {
				fmt.Printf("Agent name %s already exists. Please choose a different name.\n", result)
				goto nextPrompt
			}
		}
		return result, nil
	nextPrompt:
	}
}

func promptForUpdate() (bool, error) {
	prompt := promptui.Select{
		Label: "What do you want to do?",
		Items: []string{"Update", "Change name"},
	}
	i, _, err := prompt.Run()
	if err != nil {
		return false, fmt.Errorf("prompt failed %v", err)
	}
	return i == 0, nil
}

func getProviderName(name string) string {
	return strings.Split(name, " ")[0]
}

func isOllamaConfig(name string) bool {
	return strings.Contains(name, "ollama")
}

func selectProviderConfig(config types.Config) (types.ProviderConfig, error) {
	var providerNames []string
	var providerConfigs []types.ProviderConfig
	for _, configProvider := range config.Providers {
		providerNames = append(providerNames, configProvider.Provider+" "+configProvider.Name)
		providerConfigs = append(providerConfigs, configProvider)
	}

	if len(providerNames) == 0 {
		return types.ProviderConfig{}, fmt.Errorf("no providers configured")
	}

	prompt := promptui.Select{
		Label: "Select a provider configuration",
		Items: providerNames,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return types.ProviderConfig{}, fmt.Errorf("prompt failed %v", err)
	}

	return providerConfigs[i], nil
}
