package cli

type ProviderConfig struct {
    APIKey string
    Model  string
}

type LLMProvider struct {
    Name    string
    Setup   func() error
    Connect func() error
}

var providers = make(map[string]LLMProvider)
var configPath = "~/.aicraft/config.yaml"