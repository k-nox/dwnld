package download

import "slices"

type Resolution string

const (
	Resolution2160 Resolution = "2160"
	Resolution1440 Resolution = "1440"
	Resolution1080 Resolution = "1080"
	Resolution720  Resolution = "720"
	Resoultion480  Resolution = "480"
	Resolution360  Resolution = "360"
)

var Resolutions = []struct {
	Value  Resolution
	TSName string
}{
	{Resolution2160, "RESOLUTION_2160"},
	{Resolution1440, "RESOLUTION_1440"},
	{Resolution1080, "RESOLUTION_1080"},
	{Resolution720, "RESOLUTION_720"},
	{Resoultion480, "RESOLUTION_480"},
	{Resolution360, "RESOLUTION_360"},
}

type Options struct {
	URL             string     `json:"url"`
	OutputDir       string     `json:"outputDir,omitempty"`
	OutputTempl     string     `json:"outputTempl,omitempty"`
	TargetResoltion Resolution `json:"targetResolution,omitempty"`
}

func (res Resolution) valid() bool {
	return slices.ContainsFunc(Resolutions, func(r struct {
		Value  Resolution
		TSName string
	},
	) bool {
		return r.Value == res
	})
}
