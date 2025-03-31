package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func File(devMode bool) (string, error) {
	if devMode {
		err := check("config.yaml")
		return "config.yaml", err
	}

	path, err := xdg.ConfigFile(filepath.Join("dwnld", "dwnld.yaml"))
	if err != nil {
		return path, fmt.Errorf("error looking for config file: %w", err)
	}

	return path, check(path)
}

func create(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer file.Close()
	return nil
}

func check(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return create(path)
	}

	if err != nil {
		return fmt.Errorf("error checking for config file: %w", err)
	}

	return nil
}
