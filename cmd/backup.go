package cmd

import (
	"fmt"
	"os"

	"github.com/gilbe/claude-foundry-manager/internal/backup"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Manage configuration backups",
	Long: `Manage configuration backups for Claude Code settings.

Subcommands:
  list    - List all available backups
  create  - Create a manual backup
  restore - Restore from a specific backup

Examples:
  claude-foundry-manager backup list
  claude-foundry-manager backup create "My manual backup"
  claude-foundry-manager backup restore backup_20240115_143022.json`,
}

var backupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available backups",
	RunE: func(cmd *cobra.Command, args []string) error {
		backups, err := backup.ListBackups()
		if err != nil {
			return fmt.Errorf("failed to list backups: %w", err)
		}

		if len(backups) == 0 {
			fmt.Println("\nNo backups found.")
			fmt.Printf("Backup location: %s\n", backup.GetBackupDir())
			return nil
		}

		fmt.Printf("\n=== Available Backups (%d total) ===\n\n", len(backups))
		for i, b := range backups {
			fmt.Printf("[%d] %s\n", i+1, b.Filename)
			fmt.Printf("    Created: %s\n", b.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("    Description: %s\n", b.Description)
			if b.UseFoundry {
				fmt.Printf("    Resource: %s\n", b.Resource)
			} else {
				fmt.Printf("    Resource: (default Anthropic)\n")
			}
			fmt.Println()
		}

		fmt.Printf("Backup location: %s\n\n", backup.GetBackupDir())
		return nil
	},
}

var backupCreateCmd = &cobra.Command{
	Use:   "create [description]",
	Short: "Create a manual backup",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := args[0]

		filename, err := backup.CreateManualBackup(description)
		if err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}

		fmt.Printf("\n✓ Backup created successfully: %s\n", filename)
		fmt.Printf("Location: %s\n\n", backup.GetBackupDir())
		return nil
	},
}

var backupRestoreCmd = &cobra.Command{
	Use:   "restore [filename]",
	Short: "Restore from a specific backup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		// Create a backup before restoring (in case user wants to undo)
		if err := backup.CreateAutoBackup("Before restore operation"); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to create pre-restore backup: %v\n", err)
		}

		if err := backup.RestoreBackup(filename); err != nil {
			return fmt.Errorf("failed to restore backup: %w", err)
		}

		fmt.Printf("\n✓ Configuration restored from: %s\n", filename)
		fmt.Println("\nPlease restart your terminal for the changes to take effect.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.AddCommand(backupListCmd)
	backupCmd.AddCommand(backupCreateCmd)
	backupCmd.AddCommand(backupRestoreCmd)
}
