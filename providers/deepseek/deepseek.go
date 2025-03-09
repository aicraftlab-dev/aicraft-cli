package deepseek

import (
	"fmt"
	"os/exec"

	"github.com/aicraftlab-dev/aicraft-cli/types"
)

type DeepSeekProvider struct{}

func init() {
	types.Providers["deepseek"] = DeepSeekProvider{}
}

func (d DeepSeekProvider) Generate(modelName string, prompt string, apiKey string, host string) (string, error) {
	//Deepseek logic to get the result from the model.
	fmt.Println("Generate in DeepSeek provider")
	return "Result DeepSeek from " + modelName + " with prompt: " + prompt, nil
}
func (d DeepSeekProvider) GetModels(apiKey string, host string) ([]string, error) {
	//TODO: add real logic to get the models from the provider.
	//The apiKey and host are only to simulate that it is used.
	if apiKey == "" || host == "" {
		return []string{}, fmt.Errorf("apiKey and host needed")
	}
	return []string{"deepseek-chat", "deepseek-coder"}, nil
}

func runCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	return command.Start()
}
