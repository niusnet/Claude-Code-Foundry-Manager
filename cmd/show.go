package cmd

import (
	"fmt"

	"github.com/gilbe/claude-foundry-manager/internal/config"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long: `Display the current Claude Code configuration, including all Azure Foundry environment variables.

This command shows:
  - Whether Azure Foundry is enabled
  - Azure Foundry resource name
  - Base URL
  - API key status (masked for security)
  - Model deployment names

Example:
  claude-foundry-manager show`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetCurrentConfig()
		if err != nil {
			return fmt.Errorf("failed to read configuration: %w", err)
		}

		fmt.Println("\n=== Current Claude Code Configuration ===")

		if cfg.UseFoundry {
			fmt.Println("Status: Azure Foundry ENABLED")
			fmt.Printf("Resource: %s\n", cfg.Resource)
			fmt.Printf("Base URL: %s\n", cfg.BaseURL)

			if cfg.APIKey != "" {
				fmt.Printf("API Key: %s... (masked)\n", maskAPIKey(cfg.APIKey))
			} else {
				fmt.Println("API Key: Not set (using Entra ID)")
			}

			fmt.Printf("\nModel Deployments:\n")
			fmt.Printf("  Sonnet: %s\n", cfg.SonnetModel)
			fmt.Printf("  Haiku:  %s\n", cfg.HaikuModel)
			fmt.Printf("  Opus:   %s\n", cfg.OpusModel)
		} else {
			fmt.Println("Status: Azure Foundry DISABLED")
			fmt.Println("\nUsing default Anthropic direct API configuration.")
		}

		fmt.Println()
		return nil
	},
}

func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return "***"
	}
	return key[:8] + "***"
}

func init() {
	rootCmd.AddCommand(showCmd)
}
