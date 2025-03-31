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
	EmbedSubtitles  bool       `json:"embedSubtitles" koanf:"embedSubtitles"`
	WriteInfoJSON   bool       `json:"writeInfoJSON" koanf:"writeInfoJSON"`
}

func defaultDownloadSettings() Download {
	return Download{
		OutputTemplate:  "%(title)s [%(id)s].%(ext)s",
		TargetResoltion: Resolution720,
		EmbedSubtitles:  false,
		WriteInfoJSON:   false,
	}
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
