package cli

import (
    "testing"

    "github.com/spf13/cobra"
)

func TestRootCommand(t *testing.T) {
    rootCmd := &cobra.Command{}
    InitRootCmd(rootCmd)

    if len(rootCmd.Commands()) != 3 {
        t.Errorf("Expected 3 commands, got %d", len(rootCmd.Commands()))
    }
}

func TestAICmd(t *testing.T) {
    aiCmd := NewAICmd()
    if aiCmd.Use != "ai" {
        t.Errorf("Expected command use to be 'ai', got %s", aiCmd.Use)
    }
}

func TestCodeCmd(t *testing.T) {
    codeCmd := NewCodeCmd()
    if codeCmd.Use != "code" {
        t.Errorf("Expected command use to be 'code', got %s", codeCmd.Use)
    }
}

func TestWebCmd(t *testing.T) {
    webCmd := NewWebCmd()
    if webCmd.Use != "web" {
        t.Errorf("Expected command use to be 'web', got %s", webCmd.Use)
    }
}