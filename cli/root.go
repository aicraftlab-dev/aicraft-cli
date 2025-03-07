package cli

import (
    "github.com/spf13/cobra"
)

func InitRootCmd(rootCmd *cobra.Command) {
    rootCmd.AddCommand(NewAICmd())
    rootCmd.AddCommand(NewCodeCmd())
    rootCmd.AddCommand(NewWebCmd())
    rootCmd.AddCommand(NewConfigCmd())
    rootCmd.AddCommand(NewLLMCmd())
}