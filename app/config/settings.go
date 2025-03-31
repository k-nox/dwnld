package config

import (
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

type Manager struct {
	conf         *koanf.Koanf
	fileProvider *file.File
	yamlParser   *yaml.YAML
	filePath     string
	settings     *Settings
}

type Settings struct {
	Download Download `json:"download" koanf:"download"`
}

func defaultSettings() Settings {
	return Settings{
		Download: Download{
			OutputTemplate:  "%(title)s [%(id)s].%(ext)s",
			TargetResoltion: Resolution720,
		},
	}
}

func New(configFilePath string) (*Manager, error) {
	return &Manager{
		conf:         koanf.New("."),
		fileProvider: file.Provider(configFilePath),
		yamlParser:   yaml.Parser(),
		filePath:     configFilePath,
	}, nil
}

func (m *Manager) Startup() error {
	loaded, err := m.load()
	if err != nil {
		return err
	}
	m.settings = loaded
	return nil
}

func (m *Manager) Settings() *Settings {
	return m.settings
}

func (m *Manager) Update(updated Settings) error {
	if err := m.conf.Load(structs.Provider(updated, "koanf"), nil); err != nil {
		return fmt.Errorf("error updating settings: %w", err)
	}

	var merged Settings
	if err := m.conf.Unmarshal("", &merged); err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}

	bytes, err := m.conf.Marshal(m.yamlParser)
	if err != nil {
		return fmt.Errorf("error updating settings: %w", err)
	}

	err = os.WriteFile(m.filePath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	m.settings = &merged

	return nil
}

func (m *Manager) load() (*Settings, error) {
	defaults := defaultSettings()
	if err := m.conf.Load(structs.Provider(defaults, "koanf"), nil); err != nil {
		return nil, fmt.Errorf("error loading default settings: %w", err)
	}

	if err := m.conf.Load(m.fileProvider, m.yamlParser); err != nil {
		return nil, fmt.Errorf("error loading config file: %w", err)
	}

	var out Settings
	if err := m.conf.Unmarshal("", &out); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if out.Download.OutputDirectory == "" {
		defaultDir, err := defaultDownloadDirectory()
		if err != nil {
			return nil, fmt.Errorf("error finding default download directory: %w", err)
		}
		out.Download.OutputDirectory = defaultDir
	}
	return &out, nil
}
