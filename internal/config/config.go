package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type TelemetryConfig struct {
	Session bool `json:"session"`
	Local   bool `json:"local"`
	Project bool `json:"project"`
}

type HookConfig struct {
	Post []string `json:"post"`
}

type Config struct {
	Name        string          `json:"name"`
	Identifier  string          `json:"identifier"`
	Description string          `json:"description"`
	Template    string          `json:"template"`
	Telemetry   TelemetryConfig `json:"telemetry"`
	Hooks       HookConfig      `json:"hooks"`
}

func LoadConfig(targetDir string) (Config, error) {
	var cfg Config

	// Defaults
	cfg.Telemetry.Session = true
	cfg.Telemetry.Local = false
	cfg.Telemetry.Project = false

	// User-level config: ~/.nexusv.json
	home, err := os.UserHomeDir()
	if err == nil {
		userCfg := filepath.Join(home, ".nexusv.json")
		if data, err := os.ReadFile(userCfg); err == nil {
			_ = json.Unmarshal(data, &cfg)
		}
	}

	// Project-level config: <targetDir>/nexusv.json
	projectCfg := filepath.Join(targetDir, "nexusv.json")
	if data, err := os.ReadFile(projectCfg); err == nil {
		_ = json.Unmarshal(data, &cfg)
	}

	return cfg, nil
}
