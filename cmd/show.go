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
		} else {
			fmt.Println("Status: Azure Foundry DISABLED")
		}

		// Always show environment variable values
		fmt.Println("\nEnvironment Variables:")
		fmt.Printf("  CLAUDE_CODE_USE_FOUNDRY:        %s\n", formatValue(cfg.UseFoundry))

		if cfg.Resource != "" {
			fmt.Printf("  ANTHROPIC_FOUNDRY_RESOURCE:     %s\n", formatEnvValue(cfg.Resource))
			fmt.Printf("  ANTHROPIC_FOUNDRY_BASE_URL:     %s (auto-generated)\n", formatEnvValue(cfg.BaseURL))
		} else {
			fmt.Printf("  ANTHROPIC_FOUNDRY_RESOURCE:     %s\n", formatEnvValue(""))
			fmt.Printf("  ANTHROPIC_FOUNDRY_BASE_URL:     %s\n", formatEnvValue(cfg.BaseURL))
		}

		if cfg.APIKey != "" {
			fmt.Printf("  ANTHROPIC_FOUNDRY_API_KEY:      %s... (masked)\n", maskAPIKey(cfg.APIKey))
		} else {
			fmt.Printf("  ANTHROPIC_FOUNDRY_API_KEY:      %s\n", formatEnvValue(""))
		}

		fmt.Printf("\nModel Deployments:\n")
		fmt.Printf("  ANTHROPIC_DEFAULT_SONNET_MODEL: %s\n", formatEnvValue(cfg.SonnetModel))
		fmt.Printf("  ANTHROPIC_DEFAULT_HAIKU_MODEL:  %s\n", formatEnvValue(cfg.HaikuModel))
		fmt.Printf("  ANTHROPIC_DEFAULT_OPUS_MODEL:   %s\n", formatEnvValue(cfg.OpusModel))

		if !cfg.UseFoundry {
			fmt.Println("\nNote: Using default Anthropic direct API configuration.")
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

func formatValue(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func formatEnvValue(value string) string {
	if value == "" {
		return "(not set)"
	}
	return value
}

func init() {
	rootCmd.AddCommand(showCmd)
}
