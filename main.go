package main

import (
	"fmt"
	"os"

	"github.com/aicraftlab-dev/aicraft-cli/cli"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aicraft-cli",
		Short: "aicraft-cli is a CLI for interacting with AI providers",
		Long: `A command-line interface for interacting with various AI providers,
                including LLM interactions, web intelligence, and code execution.`,
	}

	cli.InitRootCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
