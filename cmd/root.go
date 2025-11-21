package cmd

import (
	"fmt"
	"os"

	"github.com/gilbe/claude-foundry-manager/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "claude-foundry-manager",
	Short: "Claude Code - Azure Foundry Configuration Manager",
	Long: `A cross-platform CLI tool to manage Claude Code configuration between Azure AI Foundry and direct Anthropic service.

This tool helps you easily switch between providers by managing environment variables across Windows, Linux, and macOS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, run interactive mode
		if err := ui.RunInteractive(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
