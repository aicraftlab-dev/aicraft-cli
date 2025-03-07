package cli

import (
    "os"
    "path/filepath"
    "testing"
)

func TestLLMCommand(t *testing.T) {
    tempDir := t.TempDir()
    configPath := filepath.Join(tempDir, "config.yaml")
    os.Setenv("HOME", tempDir)

    tests := []struct {
        name     string
        provider string
        setup    bool
        connect  bool
    }{
        {
            name:     "Deepseek setup",
            provider: "deepseek",
            setup:    true,
        },
        {
            name:     "Ollama setup",
            provider: "ollama",
            setup:    true,
        },
        {
            name:     "Deepseek connect",
            provider: "deepseek",
            connect:  true,
        },
        {
            name:     "Ollama connect",
            provider: "ollama",
            connect:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cmd := NewLLMCmd()
            
            if tt.setup {
                cmd.SetArgs([]string{"setup", "--provider", tt.provider})
                if err := cmd.Execute(); err != nil {
                    t.Errorf("Setup failed: %v", err)
                }
            }
            
            if tt.connect {
                cmd.SetArgs([]string{"connect", "--provider", tt.provider})
                if err := cmd.Execute(); err != nil {
                    t.Errorf("Connect failed: %v", err)
                }
                
                // Verify config
                config := loadConfig(configPath)
                if config.AI.Providers[tt.provider].APIKey == "" {
                    t.Error("Expected API key/address to be set")
                }
            }
        })
    }
}