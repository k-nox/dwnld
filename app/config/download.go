package config

type Download struct {
	OutputDirectory string     `json:"outputDirectory,omitempty" koanf:"outputDirectory"`
	OutputTemplate  string     `json:"outputTemplate,omitempty" koanf:"outputTemplate"`
	TargetResoltion Resolution `json:"targetResolution,omitempty" koanf:"targetResolution"`
}
