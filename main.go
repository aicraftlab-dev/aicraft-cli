package main

import (
    "fmt"
    "os"

    "github.com/aicraftlab-dev/aicraft-cli/cli"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "aicraft",
    Short: "AICraft CLI - Unified AI Toolkit",
    Long: `AICraft CLI provides a unified interface for AI inference,
code execution, and web intelligence operations.`,
}

func main() {
    cli.InitRootCmd(rootCmd)
    
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}