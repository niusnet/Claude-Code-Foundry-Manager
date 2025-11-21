package backup

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gilbe/claude-foundry-manager/internal/config"
)

// Backup represents a saved configuration backup
type Backup struct {
	Timestamp   time.Time         `json:"timestamp"`
	Description string            `json:"description"`
	Variables   map[string]string `json:"variables"`
}

// BackupInfo represents metadata about a backup file
type BackupInfo struct {
	Filename    string
	Timestamp   time.Time
	Description string
	UseFoundry  bool
	Resource    string
}

// GetBackupDir returns the directory where backups are stored
func GetBackupDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("failed to get home directory")
	}
	return filepath.Join(home, ".claude-code-backups")
}

// ensureBackupDir creates the backup directory if it doesn't exist
func ensureBackupDir() error {
	dir := GetBackupDir()
	return os.MkdirAll(dir, 0755)
}

// CreateAutoBackup creates an automatic backup with a description
func CreateAutoBackup(description string) error {
	return createBackup(description)
}

// CreateManualBackup creates a manual backup with a user-provided description
func CreateManualBackup(description string) (string, error) {
	filename, err := createBackup(description)
	if err != nil {
		return "", err
	}
	return filepath.Base(filename), nil
}

// createBackup creates a backup file with current configuration
func createBackup(description string) (string, error) {
	if err := ensureBackupDir(); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Get current configuration
	vars := config.GetAllVars()

	backup := Backup{
		Timestamp:   time.Now(),
		Description: description,
		Variables:   vars,
	}

	// Generate filename with timestamp
	filename := fmt.Sprintf("backup_%s.json", time.Now().Format("20060102_150405"))
	filepath := filepath.Join(GetBackupDir(), filename)

	// Marshal to JSON
	data, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal backup: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write backup file: %w", err)
	}

	return filepath, nil
}

// ListBackups returns a list of all available backups, sorted by timestamp (newest first)
func ListBackups() ([]BackupInfo, error) {
	dir := GetBackupDir()

	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return []BackupInfo{}, nil
	}

	// Read directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	backups := []BackupInfo{}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		// Read backup file
		filepath := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(filepath)
		if err != nil {
			continue // Skip files that can't be read
		}

		var backup Backup
		if err := json.Unmarshal(data, &backup); err != nil {
			continue // Skip invalid JSON files
		}

		info := BackupInfo{
			Filename:    entry.Name(),
			Timestamp:   backup.Timestamp,
			Description: backup.Description,
			UseFoundry:  backup.Variables[config.EnvUseFoundry] == "true",
			Resource:    backup.Variables[config.EnvFoundryResource],
		}

		backups = append(backups, info)
	}

	// Sort by timestamp, newest first
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Timestamp.After(backups[j].Timestamp)
	})

	return backups, nil
}

// RestoreBackup restores configuration from a backup file
func RestoreBackup(filename string) error {
	filepath := filepath.Join(GetBackupDir(), filename)

	// Read backup file
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %w", err)
	}

	var backup Backup
	if err := json.Unmarshal(data, &backup); err != nil {
		return fmt.Errorf("failed to parse backup file: %w", err)
	}

	// First, rollback to default to clear all variables
	if err := config.RollbackToDefault(); err != nil {
		return fmt.Errorf("failed to clear existing configuration: %w", err)
	}

	// Then, restore variables from backup
	if len(backup.Variables) > 0 {
		if err := config.SetAllVars(backup.Variables); err != nil {
			return fmt.Errorf("failed to restore variables: %w", err)
		}
	}

	return nil
}

// DeleteBackup removes a backup file
func DeleteBackup(filename string) error {
	filepath := filepath.Join(GetBackupDir(), filename)
	return os.Remove(filepath)
}
