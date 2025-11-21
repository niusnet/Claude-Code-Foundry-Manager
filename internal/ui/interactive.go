package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gilbe/claude-foundry-manager/internal/backup"
	"github.com/gilbe/claude-foundry-manager/internal/config"
)

// Color codes for terminal output
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

var reader = bufio.NewReader(os.Stdin)

// RunInteractive starts the interactive menu
func RunInteractive() error {
	for {
		showBanner()
		showMenu()

		choice, err := readInput("Enter your choice (1-7): ")
		if err != nil {
			return err
		}

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			if err := handleConfigure(); err != nil {
				printError(fmt.Sprintf("Configuration failed: %v", err))
			}
		case "2":
			if err := handleRollback(); err != nil {
				printError(fmt.Sprintf("Rollback failed: %v", err))
			}
		case "3":
			if err := handleShowConfig(); err != nil {
				printError(fmt.Sprintf("Failed to show configuration: %v", err))
			}
		case "4":
			if err := handleListBackups(); err != nil {
				printError(fmt.Sprintf("Failed to list backups: %v", err))
			}
		case "5":
			if err := handleRestoreBackup(); err != nil {
				printError(fmt.Sprintf("Failed to restore backup: %v", err))
			}
		case "6":
			if err := handleCreateBackup(); err != nil {
				printError(fmt.Sprintf("Failed to create backup: %v", err))
			}
		case "7", "q", "quit", "exit":
			printInfo("\nGoodbye!")
			return nil
		default:
			printError("Invalid choice. Please enter 1-7.")
		}

		fmt.Println("\nPress Enter to continue...")
		reader.ReadString('\n')
	}
}

func showBanner() {
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println(colorCyan + colorBold + "=" + strings.Repeat("=", 68) + colorReset)
	fmt.Println(colorCyan + colorBold + "    CLAUDE CODE - AZURE FOUNDRY CONFIGURATION MANAGER (Go Edition)" + colorReset)
	fmt.Println(colorCyan + colorBold + "=" + strings.Repeat("=", 68) + colorReset)
	fmt.Println()
}

func showMenu() {
	fmt.Println("Please select an option:")
	fmt.Println()
	fmt.Println("  " + colorGreen + "[1]" + colorReset + " Configure Azure Foundry")
	fmt.Println("  " + colorYellow + "[2]" + colorReset + " Rollback to Default (Direct Anthropic)")
	fmt.Println("  " + colorBlue + "[3]" + colorReset + " View Current Configuration")
	fmt.Println("  " + colorCyan + "[4]" + colorReset + " List Available Backups")
	fmt.Println("  " + colorCyan + "[5]" + colorReset + " Restore from Backup")
	fmt.Println("  " + colorCyan + "[6]" + colorReset + " Save Manual Backup")
	fmt.Println("  " + colorRed + "[7]" + colorReset + " Exit")
	fmt.Println()
}

func handleConfigure() error {
	printInfo("\n=== Configure Azure Foundry ===\n")

	resource, err := readInput("Enter Azure Foundry Resource name: ")
	if err != nil {
		return err
	}
	resource = strings.TrimSpace(resource)
	if resource == "" {
		return fmt.Errorf("resource name is required")
	}

	apiKey, err := readInput("Enter API Key (leave empty for Entra ID): ")
	if err != nil {
		return err
	}
	apiKey = strings.TrimSpace(apiKey)

	sonnetModel, err := readInputWithDefault("Sonnet model deployment name", "claude-sonnet-4-5")
	if err != nil {
		return err
	}

	haikuModel, err := readInputWithDefault("Haiku model deployment name", "claude-haiku-4-5")
	if err != nil {
		return err
	}

	opusModel, err := readInputWithDefault("Opus model deployment name", "claude-opus-4-1")
	if err != nil {
		return err
	}

	// Show summary
	fmt.Println("\n" + colorYellow + "Configuration Summary:" + colorReset)
	fmt.Printf("  Resource: %s\n", resource)
	if apiKey != "" {
		fmt.Printf("  API Key: %s... (masked)\n", maskAPIKey(apiKey))
	} else {
		fmt.Println("  API Key: (using Entra ID)")
	}
	fmt.Printf("  Sonnet Model: %s\n", sonnetModel)
	fmt.Printf("  Haiku Model: %s\n", haikuModel)
	fmt.Printf("  Opus Model: %s\n", opusModel)

	confirm, err := readInput("\nApply this configuration? (y/n): ")
	if err != nil {
		return err
	}

	if !strings.EqualFold(strings.TrimSpace(confirm), "y") {
		printInfo("Configuration cancelled.")
		return nil
	}

	// Create backup
	if err := backup.CreateAutoBackup("Before configuring Azure Foundry"); err != nil {
		printWarning(fmt.Sprintf("Failed to create backup: %v", err))
	}

	// Apply configuration
	cfg := &config.FoundryConfig{
		Resource:    resource,
		APIKey:      apiKey,
		SonnetModel: sonnetModel,
		HaikuModel:  haikuModel,
		OpusModel:   opusModel,
	}

	if err := config.ApplyFoundryConfig(cfg); err != nil {
		return err
	}

	printSuccess("\n✓ Azure Foundry configuration applied successfully!")
	printInfo("\nPlease restart your terminal for the changes to take effect.")

	return nil
}

func handleRollback() error {
	printWarning("\n=== Rollback to Default Configuration ===\n")
	printWarning("This will remove all Azure Foundry settings and return to direct Anthropic API.\n")

	confirm, err := readInput("Are you sure? (y/n): ")
	if err != nil {
		return err
	}

	if !strings.EqualFold(strings.TrimSpace(confirm), "y") {
		printInfo("Rollback cancelled.")
		return nil
	}

	// Create backup
	if err := backup.CreateAutoBackup("Before rollback to default"); err != nil {
		printWarning(fmt.Sprintf("Failed to create backup: %v", err))
	}

	// Rollback
	if err := config.RollbackToDefault(); err != nil {
		return err
	}

	printSuccess("\n✓ Successfully rolled back to default Anthropic configuration!")
	printInfo("\nPlease restart your terminal for the changes to take effect.")

	return nil
}

func handleShowConfig() error {
	cfg, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	fmt.Println("\n" + colorCyan + colorBold + "=== Current Configuration ===" + colorReset + "\n")

	if cfg.UseFoundry {
		printSuccess("Status: Azure Foundry ENABLED")
		fmt.Printf("\n  Resource:    %s\n", cfg.Resource)
		fmt.Printf("  Base URL:    %s\n", cfg.BaseURL)

		if cfg.APIKey != "" {
			fmt.Printf("  API Key:     %s... (masked)\n", maskAPIKey(cfg.APIKey))
		} else {
			fmt.Println("  API Key:     (using Entra ID)")
		}

		fmt.Println("\n  Model Deployments:")
		fmt.Printf("    Sonnet: %s\n", cfg.SonnetModel)
		fmt.Printf("    Haiku:  %s\n", cfg.HaikuModel)
		fmt.Printf("    Opus:   %s\n", cfg.OpusModel)
	} else {
		printInfo("Status: Azure Foundry DISABLED")
		fmt.Println("\nUsing default Anthropic direct API configuration.")
	}

	fmt.Println()
	return nil
}

func handleListBackups() error {
	backups, err := backup.ListBackups()
	if err != nil {
		return err
	}

	if len(backups) == 0 {
		printInfo("\nNo backups found.")
		fmt.Printf("Backup location: %s\n", backup.GetBackupDir())
		return nil
	}

	fmt.Printf("\n" + colorCyan + colorBold + "=== Available Backups (%d total) ===" + colorReset + "\n\n", len(backups))

	for i, b := range backups {
		fmt.Printf(colorYellow+"[%d]"+colorReset+" %s\n", i+1, b.Filename)
		fmt.Printf("    Created: %s\n", b.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Printf("    Description: %s\n", b.Description)
		if b.UseFoundry {
			fmt.Printf("    Resource: %s\n", b.Resource)
		} else {
			fmt.Printf("    Resource: (default Anthropic)\n")
		}
		fmt.Println()
	}

	fmt.Printf("Backup location: %s\n", backup.GetBackupDir())
	return nil
}

func handleRestoreBackup() error {
	backups, err := backup.ListBackups()
	if err != nil {
		return err
	}

	if len(backups) == 0 {
		printInfo("\nNo backups available to restore.")
		return nil
	}

	// Show list
	fmt.Println("\n" + colorCyan + "=== Select Backup to Restore ===" + colorReset + "\n")
	for i, b := range backups {
		fmt.Printf("[%d] %s - %s\n", i+1, b.Timestamp.Format("2006-01-02 15:04:05"), b.Description)
	}

	input, err := readInput("\nEnter backup number (or 'c' to cancel): ")
	if err != nil {
		return err
	}

	input = strings.TrimSpace(input)
	if strings.EqualFold(input, "c") {
		printInfo("Restore cancelled.")
		return nil
	}

	var selection int
	if _, err := fmt.Sscanf(input, "%d", &selection); err != nil || selection < 1 || selection > len(backups) {
		return fmt.Errorf("invalid selection")
	}

	selectedBackup := backups[selection-1]

	// Confirm
	confirm, err := readInput(fmt.Sprintf("\nRestore from '%s'? (y/n): ", selectedBackup.Filename))
	if err != nil {
		return err
	}

	if !strings.EqualFold(strings.TrimSpace(confirm), "y") {
		printInfo("Restore cancelled.")
		return nil
	}

	// Create backup before restoring
	if err := backup.CreateAutoBackup("Before restore operation"); err != nil {
		printWarning(fmt.Sprintf("Failed to create backup: %v", err))
	}

	// Restore
	if err := backup.RestoreBackup(selectedBackup.Filename); err != nil {
		return err
	}

	printSuccess(fmt.Sprintf("\n✓ Configuration restored from: %s", selectedBackup.Filename))
	printInfo("\nPlease restart your terminal for the changes to take effect.")

	return nil
}

func handleCreateBackup() error {
	fmt.Println("\n" + colorCyan + "=== Create Manual Backup ===" + colorReset + "\n")

	description, err := readInput("Enter backup description: ")
	if err != nil {
		return err
	}

	description = strings.TrimSpace(description)
	if description == "" {
		description = "Manual backup"
	}

	filename, err := backup.CreateManualBackup(description)
	if err != nil {
		return err
	}

	printSuccess(fmt.Sprintf("\n✓ Backup created successfully: %s", filename))
	fmt.Printf("Location: %s\n", backup.GetBackupDir())

	return nil
}

// Helper functions

func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readInputWithDefault(prompt, defaultValue string) (string, error) {
	input, err := readInput(fmt.Sprintf("%s [%s]: ", prompt, defaultValue))
	if err != nil {
		return "", err
	}
	if input == "" {
		return defaultValue, nil
	}
	return input, nil
}

func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return "***"
	}
	return key[:8] + "***"
}

func printSuccess(msg string) {
	fmt.Println(colorGreen + msg + colorReset)
}

func printError(msg string) {
	fmt.Println(colorRed + "Error: " + msg + colorReset)
}

func printWarning(msg string) {
	fmt.Println(colorYellow + "Warning: " + msg + colorReset)
}

func printInfo(msg string) {
	fmt.Println(colorCyan + msg + colorReset)
}
