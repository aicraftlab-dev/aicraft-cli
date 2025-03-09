package types

type ProviderConfig struct {
	Provider string `yaml:"provider"`
	Name     string `yaml:"name"`
	APIKey   string `yaml:"apiKey"`
	Host     string `yaml:"host"`
}
type Agent struct {
	Name               string            `yaml:"name"`
	Model              string            `yaml:"model"`
	ProviderConfigName string            `yaml:"provider"`
	Options            map[string]string `yaml:"options,omitempty"`
}

type Config struct {
	Providers []ProviderConfig `yaml:"providers"`
	Agents    []Agent          `yaml:"agents,omitempty"`
}

type Provider interface {
	GetModels(apiKey string, host string) ([]string, error)
	Generate(model string, prompt string, apiKey string, host string) (string, error)
}

var Providers = make(map[string]Provider)
