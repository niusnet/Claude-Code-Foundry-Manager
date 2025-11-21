package backup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gilbe/claude-foundry-manager/internal/config"
)

func TestGetBackupDir(t *testing.T) {
	dir := GetBackupDir()
	if dir == "" {
		t.Error("GetBackupDir returned empty string")
	}

	home, _ := os.UserHomeDir()
	expected := filepath.Join(home, ".claude-code-backups")
	if dir != expected {
		t.Errorf("Expected %s, got %s", expected, dir)
	}
}

func TestBackupStructure(t *testing.T) {
	// Test Backup struct JSON marshaling
	testVars := map[string]string{
		config.EnvUseFoundry:      "true",
		config.EnvFoundryResource: "test-resource",
	}

	backup := Backup{
		Timestamp:   time.Now(),
		Description: "Test backup",
		Variables:   testVars,
	}

	// Marshal to JSON
	data, err := json.Marshal(backup)
	if err != nil {
		t.Fatalf("Failed to marshal backup: %v", err)
	}

	// Unmarshal back
	var restored Backup
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("Failed to unmarshal backup: %v", err)
	}

	// Verify data
	if restored.Description != backup.Description {
		t.Errorf("Description mismatch: expected '%s', got '%s'", backup.Description, restored.Description)
	}

	if len(restored.Variables) != len(backup.Variables) {
		t.Errorf("Variables count mismatch: expected %d, got %d", len(backup.Variables), len(restored.Variables))
	}

	for key, value := range backup.Variables {
		if restored.Variables[key] != value {
			t.Errorf("Variable mismatch for %s: expected '%s', got '%s'", key, value, restored.Variables[key])
		}
	}
}

func TestBackupInfoStructure(t *testing.T) {
	info := BackupInfo{
		Filename:    "backup_20240115_143022.json",
		Timestamp:   time.Now(),
		Description: "Test backup",
		UseFoundry:  true,
		Resource:    "test-resource",
	}

	if info.Filename == "" {
		t.Error("BackupInfo filename is empty")
	}

	if info.Timestamp.IsZero() {
		t.Error("BackupInfo timestamp is zero")
	}
}

func TestListBackupsWithRealDirectory(t *testing.T) {
	// This test uses the real backup directory
	// It should not fail even if directory doesn't exist
	backups, err := ListBackups()
	if err != nil {
		t.Fatalf("Failed to list backups: %v", err)
	}

	// Should return a list (empty or with items)
	if backups == nil {
		t.Error("ListBackups returned nil instead of empty slice")
	}

	// If there are backups, verify structure
	for _, backup := range backups {
		if backup.Filename == "" {
			t.Error("Backup has empty filename")
		}

		if backup.Timestamp.IsZero() {
			t.Error("Backup has zero timestamp")
		}
	}
}

func TestBackupFilenameFormat(t *testing.T) {
	// Test that backup filenames follow expected format
	now := time.Now()
	expectedPrefix := "backup_" + now.Format("20060102")

	// Create a temporary backup to verify filename
	tmpDir := t.TempDir()

	backup := Backup{
		Timestamp:   now,
		Description: "Test filename",
		Variables:   make(map[string]string),
	}

	filename := filepath.Join(tmpDir, "backup_"+now.Format("20060102_150405")+".json")
	data, _ := json.MarshalIndent(backup, "", "  ")

	if err := os.WriteFile(filename, data, 0644); err != nil {
		t.Fatalf("Failed to write test backup: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Test backup file was not created")
	}

	// Verify filename contains expected prefix
	base := filepath.Base(filename)
	if len(base) < len(expectedPrefix) || base[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Filename '%s' doesn't start with expected prefix '%s'", base, expectedPrefix)
	}
}

func TestBackupJSONFormat(t *testing.T) {
	// Test that backup JSON has correct formatting
	testVars := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}

	backup := Backup{
		Timestamp:   time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC),
		Description: "Test JSON format",
		Variables:   testVars,
	}

	data, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal backup: %v", err)
	}

	// Verify it's valid JSON
	var check map[string]interface{}
	if err := json.Unmarshal(data, &check); err != nil {
		t.Fatalf("Marshaled data is not valid JSON: %v", err)
	}

	// Verify required fields exist
	if _, ok := check["timestamp"]; !ok {
		t.Error("JSON missing 'timestamp' field")
	}

	if _, ok := check["description"]; !ok {
		t.Error("JSON missing 'description' field")
	}

	if _, ok := check["variables"]; !ok {
		t.Error("JSON missing 'variables' field")
	}
}

func TestEnsureBackupDir(t *testing.T) {
	// Test that ensureBackupDir creates directory if needed
	dir := GetBackupDir()

	// Call ensureBackupDir
	err := ensureBackupDir()
	if err != nil {
		t.Fatalf("ensureBackupDir failed: %v", err)
	}

	// Verify directory exists
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		t.Error("Backup directory was not created")
	}

	if err == nil && !info.IsDir() {
		t.Error("Backup path exists but is not a directory")
	}
}
