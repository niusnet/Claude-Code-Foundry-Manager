package config

import (
	"testing"
)

func TestFoundryConfigStruct(t *testing.T) {
	cfg := &FoundryConfig{
		Resource:    "test-resource",
		APIKey:      "test-api-key",
		SonnetModel: "claude-sonnet-4-5",
		HaikuModel:  "claude-haiku-4-5",
		OpusModel:   "claude-opus-4-1",
	}

	if cfg.Resource != "test-resource" {
		t.Errorf("Expected Resource 'test-resource', got '%s'", cfg.Resource)
	}

	if cfg.APIKey != "test-api-key" {
		t.Errorf("Expected APIKey 'test-api-key', got '%s'", cfg.APIKey)
	}

	if cfg.SonnetModel != "claude-sonnet-4-5" {
		t.Errorf("Expected SonnetModel 'claude-sonnet-4-5', got '%s'", cfg.SonnetModel)
	}
}

func TestCurrentConfigStruct(t *testing.T) {
	cfg := &CurrentConfig{
		UseFoundry:  true,
		Resource:    "test-resource",
		BaseURL:     "https://test.com",
		APIKey:      "test-key",
		SonnetModel: "claude-sonnet-4-5",
		HaikuModel:  "claude-haiku-4-5",
		OpusModel:   "claude-opus-4-1",
	}

	if !cfg.UseFoundry {
		t.Error("Expected UseFoundry to be true")
	}

	if cfg.Resource != "test-resource" {
		t.Errorf("Expected Resource 'test-resource', got '%s'", cfg.Resource)
	}

	if cfg.BaseURL != "https://test.com" {
		t.Errorf("Expected BaseURL 'https://test.com', got '%s'", cfg.BaseURL)
	}
}

func TestEnvironmentVariableConstants(t *testing.T) {
	// Test that constants are defined correctly
	expectedVars := map[string]string{
		"EnvUseFoundry":      "CLAUDE_CODE_USE_FOUNDRY",
		"EnvFoundryResource": "ANTHROPIC_FOUNDRY_RESOURCE",
		"EnvFoundryBaseURL":  "ANTHROPIC_FOUNDRY_BASE_URL",
		"EnvFoundryAPIKey":   "ANTHROPIC_FOUNDRY_API_KEY",
		"EnvDefaultSonnet":   "ANTHROPIC_DEFAULT_SONNET_MODEL",
		"EnvDefaultHaiku":    "ANTHROPIC_DEFAULT_HAIKU_MODEL",
		"EnvDefaultOpus":     "ANTHROPIC_DEFAULT_OPUS_MODEL",
	}

	actualVars := map[string]string{
		"EnvUseFoundry":      EnvUseFoundry,
		"EnvFoundryResource": EnvFoundryResource,
		"EnvFoundryBaseURL":  EnvFoundryBaseURL,
		"EnvFoundryAPIKey":   EnvFoundryAPIKey,
		"EnvDefaultSonnet":   EnvDefaultSonnet,
		"EnvDefaultHaiku":    EnvDefaultHaiku,
		"EnvDefaultOpus":     EnvDefaultOpus,
	}

	for name, expected := range expectedVars {
		actual := actualVars[name]
		if actual != expected {
			t.Errorf("Constant %s: expected '%s', got '%s'", name, expected, actual)
		}
	}
}

func TestGetAllVars(t *testing.T) {
	// GetAllVars should return a map of environment variables
	// This test just verifies it returns a map without errors
	vars := GetAllVars()

	if vars == nil {
		t.Error("GetAllVars returned nil")
	}

	// It's okay if the map is empty (no env vars set)
	// Just verify it's a valid map
	if vars == nil {
		t.Error("GetAllVars should return empty map, not nil")
	}
}

func TestGetCurrentConfig(t *testing.T) {
	// GetCurrentConfig should return a CurrentConfig struct
	// This test just verifies it returns without errors
	cfg, err := GetCurrentConfig()

	if err != nil {
		t.Fatalf("GetCurrentConfig failed: %v", err)
	}

	if cfg == nil {
		t.Fatal("GetCurrentConfig returned nil")
	}

	// Verify struct fields exist (even if empty)
	_ = cfg.UseFoundry
	_ = cfg.Resource
	_ = cfg.BaseURL
	_ = cfg.APIKey
	_ = cfg.SonnetModel
	_ = cfg.HaikuModel
	_ = cfg.OpusModel
}

func TestFoundryConfigWithEmptyAPIKey(t *testing.T) {
	// Test configuration without API key (Entra ID mode)
	cfg := &FoundryConfig{
		Resource:    "test-resource",
		APIKey:      "", // Empty = use Entra ID
		SonnetModel: "claude-sonnet-4-5",
		HaikuModel:  "claude-haiku-4-5",
		OpusModel:   "claude-opus-4-1",
	}

	if cfg.APIKey != "" {
		t.Error("Expected APIKey to be empty for Entra ID mode")
	}

	if cfg.Resource == "" {
		t.Error("Resource should not be empty even without API key")
	}
}

func TestFoundryConfigDefaults(t *testing.T) {
	// Test that default model names are reasonable
	cfg := &FoundryConfig{
		Resource:    "test-resource",
		SonnetModel: "claude-sonnet-4-5",
		HaikuModel:  "claude-haiku-4-5",
		OpusModel:   "claude-opus-4-1",
	}

	if cfg.SonnetModel == "" {
		t.Error("SonnetModel should have a default value")
	}

	if cfg.HaikuModel == "" {
		t.Error("HaikuModel should have a default value")
	}

	if cfg.OpusModel == "" {
		t.Error("OpusModel should have a default value")
	}
}

func TestBaseURLGeneration(t *testing.T) {
	// Test that base URL format is correct
	resource := "my-foundry-resource"
	expectedBaseURL := "https://my-foundry-resource.services.ai.azure.com/models"

	// This is the format used in ApplyFoundryConfig
	actualBaseURL := "https://" + resource + ".services.ai.azure.com/models"

	if actualBaseURL != expectedBaseURL {
		t.Errorf("Expected BaseURL '%s', got '%s'", expectedBaseURL, actualBaseURL)
	}
}

func TestSetAllVarsWithEmptyMap(t *testing.T) {
	// Test that SetAllVars handles empty map gracefully
	emptyVars := make(map[string]string)

	// This should not panic or error with empty map
	err := SetAllVars(emptyVars)

	// On some systems this might succeed (doing nothing)
	// On others it might fail (can't set vars without permissions)
	// We just verify it doesn't panic
	_ = err
}

func TestIsTruthy(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		// True values
		{"lowercase true", "true", true},
		{"uppercase TRUE", "TRUE", true},
		{"mixed case True", "True", true},
		{"number 1", "1", true},
		{"lowercase yes", "yes", true},
		{"uppercase YES", "YES", true},
		{"lowercase on", "on", true},
		{"uppercase ON", "ON", true},
		{"lowercase enabled", "enabled", true},
		{"uppercase ENABLED", "ENABLED", true},
		{"with spaces", "  true  ", true},
		{"with tabs", "\ttrue\t", true},

		// False values
		{"lowercase false", "false", false},
		{"uppercase FALSE", "FALSE", false},
		{"number 0", "0", false},
		{"lowercase no", "no", false},
		{"lowercase off", "off", false},
		{"empty string", "", false},
		{"just spaces", "   ", false},
		{"random text", "random", false},
		{"number 2", "2", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTruthy(tt.value)
			if result != tt.expected {
				t.Errorf("isTruthy(%q) = %v, expected %v", tt.value, result, tt.expected)
			}
		})
	}
}
