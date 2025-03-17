package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

type Manager struct {
	k      *koanf.Koanf
	file   *file.File
	parser *yaml.YAML
}

type Settings struct {
	Download Download `json:"download" koanf:"download"`
}

func New(configFilePath string) (*Manager, error) {
	return &Manager{
		k:      koanf.New("."),
		file:   file.Provider(configFilePath),
		parser: yaml.Parser(),
	}, nil
}

func (m *Manager) Load() (Settings, error) {
	var out Settings
	if err := m.k.Load(m.file, m.parser); err != nil {
		return out, fmt.Errorf("error loading config file: %w", err)
	}
	if err := m.k.Unmarshal("", &out); err != nil {
		return out, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return out, nil
}

func (m *Manager) Update(updated Settings) error {
	if err := m.k.Load(structs.Provider(updated, "koanf"), nil); err != nil {
		return fmt.Errorf("error updating settings: %w", err)
	}
	if _, err := m.k.Marshal(m.parser); err != nil {
		return fmt.Errorf("error updating settings: %w", err)
	}
	return nil
}
