package theme

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
)

var (
	DarkBackground  = &options.RGBA{R: 2, G: 8, B: 23, A: 255}
	LightBackground = &options.RGBA{R: 255, G: 255, B: 255, A: 255}
)

type Settings struct {
	ctx context.Context
}

func (s *Settings) Startup(ctx context.Context) {
	s.ctx = ctx
}
