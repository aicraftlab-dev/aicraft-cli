package cli

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewWebCmd() *cobra.Command {
    var engine string

    cmd := &cobra.Command{
        Use:   "web",
        Short: "Web intelligence operations",
        Long:  `Perform web searches and URL content extraction`,
    }

    searchCmd := &cobra.Command{
        Use:   "search [query]",
        Short: "Search the web",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Searching web with %s engine\nQuery: %s\n", engine, args[0])
        },
    }

    searchCmd.Flags().StringVarP(&engine, "engine", "e", "google", "Search engine")
    cmd.AddCommand(searchCmd)

    return cmd
}