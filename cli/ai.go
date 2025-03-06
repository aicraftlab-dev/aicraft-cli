package cli

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewAICmd() *cobra.Command {
    var model string

    cmd := &cobra.Command{
        Use:   "ai",
        Short: "AI inference operations",
        Long:  `Perform AI inference operations using various models and providers`,
    }

    generateCmd := &cobra.Command{
        Use:   "generate [prompt]",
        Short: "Generate text using AI",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Generating text with model %s\nPrompt: %s\n", model, args[0])
        },
    }

    generateCmd.Flags().StringVarP(&model, "model", "m", "ollama:llama2", "AI model to use")
    cmd.AddCommand(generateCmd)

    return cmd
}