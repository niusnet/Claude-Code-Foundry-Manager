//go:build !windows

package config

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	markerBegin = "# >>> Claude Foundry Manager - BEGIN >>>"
	markerEnd   = "# <<< Claude Foundry Manager - END <<<"
)

// getEnvVarUnix reads an environment variable from the current process environment
func getEnvVarUnix(key string) (string, error) {
	// On Unix, we read from the profile files, not from the current environment
	// because we need to persist changes across sessions
	value := os.Getenv(key)
	return value, nil
}

// setEnvVarUnix writes environment variables to shell profile files
func setEnvVarUnix(key, value string) error {
	// Get all variables to write them together
	vars := getAllVarsFromProfile()
	vars[key] = value

	return writeVarsToProfile(vars)
}

// deleteEnvVarUnix removes an environment variable from shell profile files
func deleteEnvVarUnix(key string) error {
	// Get all variables except the one to delete
	vars := getAllVarsFromProfile()
	delete(vars, key)

	if len(vars) == 0 {
		// If no variables left, remove the entire block
		return removeBlockFromProfile()
	}

	return writeVarsToProfile(vars)
}

// getAllVarsFromProfile reads all Claude Foundry variables from the profile
func getAllVarsFromProfile() map[string]string {
	vars := make(map[string]string)
	profilePath := getProfilePath()

	file, err := os.Open(profilePath)
	if err != nil {
		return vars // Return empty map if file doesn't exist
	}
	defer file.Close()

	inBlock := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, markerBegin) {
			inBlock = true
			continue
		}
		if strings.Contains(line, markerEnd) {
			break
		}

		if inBlock && strings.HasPrefix(strings.TrimSpace(line), "export ") {
			// Parse: export KEY="VALUE"
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(strings.TrimPrefix(parts[0], "export"))
				value := strings.Trim(parts[1], `"`)
				vars[key] = value
			}
		}
	}

	return vars
}

// writeVarsToProfile writes variables to the shell profile
func writeVarsToProfile(vars map[string]string) error {
	profilePath := getProfilePath()

	// Read existing content
	content := []string{}
	if file, err := os.Open(profilePath); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		inBlock := false

		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, markerBegin) {
				inBlock = true
				continue
			}
			if strings.Contains(line, markerEnd) {
				inBlock = false
				continue
			}

			if !inBlock {
				content = append(content, line)
			}
		}
	}

	// Append our block
	content = append(content, "")
	content = append(content, markerBegin)
	content = append(content, "# Claude Code Azure Foundry Configuration")
	content = append(content, "# Managed by claude-foundry-manager - DO NOT EDIT MANUALLY")

	for key, value := range vars {
		content = append(content, fmt.Sprintf(`export %s="%s"`, key, value))
	}

	content = append(content, markerEnd)
	content = append(content, "")

	// Write back
	return os.WriteFile(profilePath, []byte(strings.Join(content, "\n")), 0644)
}

// removeBlockFromProfile removes the entire Claude Foundry block
func removeBlockFromProfile() error {
	profilePath := getProfilePath()

	// Read existing content
	content := []string{}
	file, err := os.Open(profilePath)
	if err != nil {
		return nil // If file doesn't exist, nothing to remove
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inBlock := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, markerBegin) {
			inBlock = true
			continue
		}
		if strings.Contains(line, markerEnd) {
			inBlock = false
			continue
		}

		if !inBlock {
			content = append(content, line)
		}
	}

	// Write back
	return os.WriteFile(profilePath, []byte(strings.Join(content, "\n")), 0644)
}

// getProfilePath determines which shell profile file to use
func getProfilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("failed to get home directory")
	}

	// Detect shell
	shell := os.Getenv("SHELL")

	if strings.Contains(shell, "zsh") {
		return filepath.Join(home, ".zshrc")
	} else if strings.Contains(shell, "bash") {
		// Check if .bash_profile exists (macOS prefers this)
		bashProfile := filepath.Join(home, ".bash_profile")
		if _, err := os.Stat(bashProfile); err == nil {
			return bashProfile
		}
		return filepath.Join(home, ".bashrc")
	} else if strings.Contains(shell, "fish") {
		configDir := filepath.Join(home, ".config", "fish")
		os.MkdirAll(configDir, 0755)
		return filepath.Join(configDir, "config.fish")
	}

	// Default to .profile (POSIX standard)
	return filepath.Join(home, ".profile")
}

// getCurrentShell returns the current shell name
func getCurrentShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		// Try to get from ps
		cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", os.Getppid()), "-o", "comm=")
		output, err := cmd.Output()
		if err == nil {
			shell = strings.TrimSpace(string(output))
		}
	}
	return filepath.Base(shell)
}
