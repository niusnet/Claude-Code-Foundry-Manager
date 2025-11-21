package config

import (
	"fmt"
	"runtime"
)

// Environment variable names used by Claude Code
const (
	EnvUseFoundry       = "CLAUDE_CODE_USE_FOUNDRY"
	EnvFoundryResource  = "ANTHROPIC_FOUNDRY_RESOURCE"
	EnvFoundryBaseURL   = "ANTHROPIC_FOUNDRY_BASE_URL"
	EnvFoundryAPIKey    = "ANTHROPIC_FOUNDRY_API_KEY"
	EnvDefaultSonnet    = "ANTHROPIC_DEFAULT_SONNET_MODEL"
	EnvDefaultHaiku     = "ANTHROPIC_DEFAULT_HAIKU_MODEL"
	EnvDefaultOpus      = "ANTHROPIC_DEFAULT_OPUS_MODEL"
)

// FoundryConfig represents the Azure Foundry configuration
type FoundryConfig struct {
	Resource    string
	APIKey      string // Optional - if empty, use Entra ID
	SonnetModel string
	HaikuModel  string
	OpusModel   string
}

// CurrentConfig represents the current system configuration
type CurrentConfig struct {
	UseFoundry  bool
	Resource    string
	BaseURL     string
	APIKey      string
	SonnetModel string
	HaikuModel  string
	OpusModel   string
}

// ApplyFoundryConfig applies Azure Foundry configuration to the system
func ApplyFoundryConfig(cfg *FoundryConfig) error {
	baseURL := fmt.Sprintf("https://%s.services.ai.azure.com/models", cfg.Resource)

	vars := map[string]string{
		EnvUseFoundry:      "true",
		EnvFoundryResource: cfg.Resource,
		EnvFoundryBaseURL:  baseURL,
		EnvDefaultSonnet:   cfg.SonnetModel,
		EnvDefaultHaiku:    cfg.HaikuModel,
		EnvDefaultOpus:     cfg.OpusModel,
	}

	// Only set API key if provided (otherwise use Entra ID)
	if cfg.APIKey != "" {
		vars[EnvFoundryAPIKey] = cfg.APIKey
	}

	// Set all environment variables
	for key, value := range vars {
		if err := setEnvVar(key, value); err != nil {
			return fmt.Errorf("failed to set %s: %w", key, err)
		}
	}

	// Notify system of environment changes
	if err := notifyEnvironmentChange(); err != nil {
		return fmt.Errorf("failed to notify system of changes: %w", err)
	}

	return nil
}

// RollbackToDefault removes all Azure Foundry configuration
func RollbackToDefault() error {
	vars := []string{
		EnvUseFoundry,
		EnvFoundryResource,
		EnvFoundryBaseURL,
		EnvFoundryAPIKey,
		EnvDefaultSonnet,
		EnvDefaultHaiku,
		EnvDefaultOpus,
	}

	for _, key := range vars {
		if err := deleteEnvVar(key); err != nil {
			// Continue even if some deletions fail
			fmt.Printf("Warning: failed to delete %s: %v\n", key, err)
		}
	}

	// Notify system of environment changes
	if err := notifyEnvironmentChange(); err != nil {
		return fmt.Errorf("failed to notify system of changes: %w", err)
	}

	return nil
}

// GetCurrentConfig reads the current configuration from the system
func GetCurrentConfig() (*CurrentConfig, error) {
	cfg := &CurrentConfig{}

	useFoundry, _ := getEnvVar(EnvUseFoundry)
	cfg.UseFoundry = useFoundry == "true"

	cfg.Resource, _ = getEnvVar(EnvFoundryResource)
	cfg.BaseURL, _ = getEnvVar(EnvFoundryBaseURL)
	cfg.APIKey, _ = getEnvVar(EnvFoundryAPIKey)
	cfg.SonnetModel, _ = getEnvVar(EnvDefaultSonnet)
	cfg.HaikuModel, _ = getEnvVar(EnvDefaultHaiku)
	cfg.OpusModel, _ = getEnvVar(EnvDefaultOpus)

	return cfg, nil
}

// GetAllVars returns all environment variable values as a map
func GetAllVars() map[string]string {
	vars := make(map[string]string)

	keys := []string{
		EnvUseFoundry,
		EnvFoundryResource,
		EnvFoundryBaseURL,
		EnvFoundryAPIKey,
		EnvDefaultSonnet,
		EnvDefaultHaiku,
		EnvDefaultOpus,
	}

	for _, key := range keys {
		value, _ := getEnvVar(key)
		if value != "" {
			vars[key] = value
		}
	}

	return vars
}

// SetAllVars sets multiple environment variables from a map
func SetAllVars(vars map[string]string) error {
	for key, value := range vars {
		if err := setEnvVar(key, value); err != nil {
			return fmt.Errorf("failed to set %s: %w", key, err)
		}
	}

	// Notify system of environment changes
	if err := notifyEnvironmentChange(); err != nil {
		return fmt.Errorf("failed to notify system of changes: %w", err)
	}

	return nil
}

// Platform-specific implementations are in separate files:
// - manager_windows.go: Windows registry implementation
// - manager_unix.go: Linux/macOS shell profile implementation

// getEnvVar reads an environment variable from the system
func getEnvVar(key string) (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getEnvVarWindows(key)
	default:
		return getEnvVarUnix(key)
	}
}

// setEnvVar writes an environment variable to the system
func setEnvVar(key, value string) error {
	switch runtime.GOOS {
	case "windows":
		return setEnvVarWindows(key, value)
	default:
		return setEnvVarUnix(key, value)
	}
}

// deleteEnvVar removes an environment variable from the system
func deleteEnvVar(key string) error {
	switch runtime.GOOS {
	case "windows":
		return deleteEnvVarWindows(key)
	default:
		return deleteEnvVarUnix(key)
	}
}

// notifyEnvironmentChange notifies the system that environment variables have changed
func notifyEnvironmentChange() error {
	switch runtime.GOOS {
	case "windows":
		return notifyEnvironmentChangeWindows()
	default:
		// Unix systems don't need notification
		return nil
	}
}
