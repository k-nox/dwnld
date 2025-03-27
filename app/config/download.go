package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Download struct {
	OutputDirectory string     `json:"outputDirectory,omitempty" koanf:"outputDirectory"`
	OutputTemplate  string     `json:"outputTemplate,omitempty" koanf:"outputTemplate"`
	TargetResoltion Resolution `json:"targetResolution,omitempty" koanf:"targetResolution"`
}

func defaultDownloadDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find user directory: %w", err)
	}

	downloadDir := filepath.Join(homeDir, "Downloads")
	info, err := os.Stat(downloadDir)
	if err != nil {
		return "", fmt.Errorf("unable to find download directory: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("$HOME/Downloads is not a directory: %w", err)
	}

	return downloadDir, nil
}
