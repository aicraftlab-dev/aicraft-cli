package cli

import (
    "os"
    "path/filepath"
    "testing"
)

func TestNewConfigCmd(t *testing.T) {
    cmd := NewConfigCmd()
    if cmd.Use != "config" {
        t.Errorf("Expected command use to be 'config', got %s", cmd.Use)
    }
}

func TestConfigSetup(t *testing.T) {
    tempDir := t.TempDir()
    configPath := filepath.Join(tempDir, "config.yaml")
    os.Setenv("HOME", tempDir)

    cmd := NewConfigCmd()
    setupCmd := cmd.Commands()[0]

    // Test with required flags
    setupCmd.SetArgs([]string{"--provider", "deepseek", "--api-key", "test-key"})
    if err := setupCmd.Execute(); err != nil {
        t.Errorf("Setup command failed: %v", err)
    }

    // Verify config file
    config := loadConfig(configPath)
    if config.AI.Providers == nil {
        t.Error("Expected providers map to be initialized")
    }
    provider, exists := config.AI.Providers["deepseek"]
    if !exists {
        t.Error("Expected deepseek provider to exist")
    }
    if provider.APIKey != "test-key" {
        t.Errorf("Expected API key 'test-key', got %s", provider.APIKey)
    }
}

func TestLoadConfig(t *testing.T) {
    tempDir := t.TempDir()
    configPath := filepath.Join(tempDir, "config.yaml")
    os.Setenv("HOME", tempDir)

    // Test loading non-existent config
    config := loadConfig(configPath)
    if config.AI.Providers != nil {
        t.Error("Expected empty providers map for new config")
    }

    // Test loading existing config
    testConfig := Config{
        AI: struct {
            Providers map[string]ProviderConfig `yaml:"providers"`
        }{
            Providers: map[string]ProviderConfig{
                "test": {APIKey: "test-key"},
            },
        },
    }
    saveConfig(configPath, testConfig)

    loadedConfig := loadConfig(configPath)
    if loadedConfig.AI.Providers["test"].APIKey != "test-key" {
        t.Errorf("Expected API key 'test-key', got %s", loadedConfig.AI.Providers["test"].APIKey)
    }
}