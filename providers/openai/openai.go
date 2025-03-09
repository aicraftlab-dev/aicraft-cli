package openai

import (
	"fmt"
	"os/exec"

	"github.com/aicraftlab-dev/aicraft-cli/types"
)

type OpenAiProvider struct{}

func init() {
	types.Providers["openai"] = OpenAiProvider{}
}

func (o OpenAiProvider) Generate(modelName string, prompt string, apiKey string, host string) (string, error) {
	//OpenAI logic to get the result from the model.
	fmt.Println("Generate in OpenAi provider")
	return "Result OpenAI from " + modelName + " with prompt: " + prompt, nil
}
func (o OpenAiProvider) GetModels(apiKey string, host string) ([]string, error) {
	//TODO: add real logic to get the models from the provider.
	//The apiKey and host are only to simulate that it is used.
	if apiKey == "" || host == "" {
		return []string{}, fmt.Errorf("apiKey and host needed")
	}
	return []string{"gpt-3.5-turbo", "gpt-4"}, nil
}

func runCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	return command.Start()
}
