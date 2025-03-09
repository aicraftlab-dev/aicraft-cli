package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aicraftlab-dev/aicraft-cli/types"
)

type OllamaProvider struct{}

func init() {
	types.Providers["ollama"] = OllamaProvider{}
}

func (o OllamaProvider) Generate(modelName string, prompt string, apiKey string, host string) (string, error) {
	//Ollama logic to get the result from the model.
	fmt.Println("Generate in Ollama provider")
	//Create the url for the API.
	url := fmt.Sprintf("%s/api/generate", host)
	//Create the request body.
	requestBody, err := json.Marshal(map[string]string{
		"model":  modelName,
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}
	//Create the request to the API.
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//Read the response from the API.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Return the response.
	return string(body), nil
}

func (o OllamaProvider) GetModels(apiKey string, host string) ([]string, error) {
	//Simulate that the ollama client is installed and a call to it.
	url := fmt.Sprintf("%s/api/tags", host)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string][]map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	models := []string{}
	if modelsData, ok := data["models"]; ok {
		for _, modelData := range modelsData {
			if modelName, ok := modelData["name"].(string); ok {
				models = append(models, modelName)
			}
		}
	}

	return models, nil
}
