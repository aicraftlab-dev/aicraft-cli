package cli

import (
	"github.com/spf13/cobra"
)

func InitRootCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(NewConfigCmd())
	rootCmd.AddCommand(NewCodeCmd())
	rootCmd.AddCommand(NewWebCmd())
	rootCmd.AddCommand(NewLLMCmd()) //Added the llm command.

}
