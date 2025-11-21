package cmd

import (
	"fmt"
	"os"

	"github.com/gilbe/claude-foundry-manager/internal/backup"
	"github.com/gilbe/claude-foundry-manager/internal/config"
	"github.com/spf13/cobra"
)

var (
	resource    string
	baseURL     string
	apiKey      string
	sonnetModel string
	haikuModel  string
	opusModel   string
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Azure Foundry settings",
	Long: `Configure Claude Code to use Azure AI Foundry with the specified resource and models.

You can configure using either:
  1. --resource (resource name) - auto-generates the base URL
  2. --base-url (full URL) - provide the complete base URL

If --api-key is not provided, the tool will configure for Entra ID authentication.

Examples:
  # Configure with resource name (recommended)
  claude-foundry-manager configure --resource=my-foundry --api-key=sk-xxx

  # Configure with full base URL
  claude-foundry-manager configure --base-url=https://my-foundry.services.ai.azure.com --api-key=sk-xxx

  # Configure with Entra ID (no API key)
  claude-foundry-manager configure --resource=my-foundry

  # Configure with custom model deployments
  claude-foundry-manager configure --resource=my-foundry --sonnet-model=claude-4-5 --haiku-model=claude-haiku`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate that either resource or base-url is provided (but not both)
		if resource == "" && baseURL == "" {
			return fmt.Errorf("either --resource or --base-url is required")
		}
		if resource != "" && baseURL != "" {
			return fmt.Errorf("cannot specify both --resource and --base-url, choose one")
		}

		// Set defaults for model names if not provided
		if sonnetModel == "" {
			sonnetModel = "claude-sonnet-4-5"
		}
		if haikuModel == "" {
			haikuModel = "claude-haiku-4-5"
		}
		if opusModel == "" {
			opusModel = "claude-opus-4-1"
		}

		cfg := &config.FoundryConfig{
			Resource:    resource,
			BaseURL:     baseURL,
			APIKey:      apiKey,
			SonnetModel: sonnetModel,
			HaikuModel:  haikuModel,
			OpusModel:   opusModel,
		}

		// Create backup before making changes
		if err := backup.CreateAutoBackup("Before configuring Azure Foundry"); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to create backup: %v\n", err)
		}

		// Apply configuration
		if err := config.ApplyFoundryConfig(cfg); err != nil {
			return fmt.Errorf("failed to apply configuration: %w", err)
		}

		fmt.Println("\n✓ Azure Foundry configuration applied successfully!")
		fmt.Println("\nPlease restart your terminal for the changes to take effect.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVar(&resource, "resource", "", "Azure Foundry resource name (mutually exclusive with --base-url)")
	configureCmd.Flags().StringVar(&baseURL, "base-url", "", "Full Azure Foundry base URL (mutually exclusive with --resource)")
	configureCmd.Flags().StringVar(&apiKey, "api-key", "", "Azure Foundry API key (optional, uses Entra ID if not provided)")
	configureCmd.Flags().StringVar(&sonnetModel, "sonnet-model", "", "Sonnet model deployment name (default: claude-sonnet-4-5)")
	configureCmd.Flags().StringVar(&haikuModel, "haiku-model", "", "Haiku model deployment name (default: claude-haiku-4-5)")
	configureCmd.Flags().StringVar(&opusModel, "opus-model", "", "Opus model deployment name (default: claude-opus-4-1)")

	configureCmd.MarkFlagsOneRequired("resource", "base-url")
	configureCmd.MarkFlagsMutuallyExclusive("resource", "base-url")
}
