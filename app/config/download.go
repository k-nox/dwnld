package config

type DownloadOptions struct {
	OutputDirectory string     `json:"outputDirectory,omitempty"`
	OutputTemplate  string     `json:"outputTemplate,omitempty"`
	TargetResoltion Resolution `json:"targetResolution,omitempty"`
}
