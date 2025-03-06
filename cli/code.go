package cli

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewCodeCmd() *cobra.Command {
    var lang string
    var sandbox string

    cmd := &cobra.Command{
        Use:   "code",
        Short: "Code execution operations",
        Long:  `Execute code in various languages with sandboxing options`,
    }

    runCmd := &cobra.Command{
        Use:   "run [file]",
        Short: "Run a code file",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Running %s code in %s sandbox\nFile: %s\n", lang, sandbox, args[0])
        },
    }

    runCmd.Flags().StringVarP(&lang, "lang", "l", "python", "Programming language")
    runCmd.Flags().StringVarP(&sandbox, "sandbox", "s", "docker", "Sandbox environment")
    cmd.AddCommand(runCmd)

    return cmd
}