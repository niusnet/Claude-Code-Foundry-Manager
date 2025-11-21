package cmd

import (
	"fmt"
	"os"

	"github.com/gilbe/claude-foundry-manager/internal/backup"
	"github.com/gilbe/claude-foundry-manager/internal/config"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback to default Anthropic configuration",
	Long: `Remove all Azure Foundry environment variables and return to the default Anthropic direct configuration.

This command will:
  1. Create a backup of current settings
  2. Remove all Azure Foundry environment variables
  3. Restore default Claude Code behavior (direct Anthropic API)

Example:
  claude-foundry-manager rollback`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create backup before rolling back
		if err := backup.CreateAutoBackup("Before rollback to default"); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to create backup: %v\n", err)
		}

		// Remove all Foundry configuration
		if err := config.RollbackToDefault(); err != nil {
			return fmt.Errorf("failed to rollback: %w", err)
		}

		fmt.Println("\n✓ Successfully rolled back to default Anthropic configuration!")
		fmt.Println("\nPlease restart your terminal for the changes to take effect.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
