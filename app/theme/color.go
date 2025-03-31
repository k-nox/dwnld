package theme

import (
	"context"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	DarkBackground      = &options.RGBA{R: 2, G: 8, B: 23, A: 255}
	DarkForeground      = windows.RGB(248, 250, 252)
	DarkBorder          = windows.RGB(30, 41, 59)
	DarkMuted           = DarkBorder
	DarkMutedForeground = windows.RGB(148, 163, 184)

	LightBackground      = &options.RGBA{R: 255, G: 255, B: 255, A: 255}
	LightForegroud       = windows.RGB(2, 8, 23)
	LightBorder          = windows.RGB(226, 232, 240)
	LightMuted           = windows.RGB(241, 245, 249)
	LightMutedForeground = windows.RGB(100, 116, 139)
)

type Settings struct {
	ctx context.Context
}

func (s *Settings) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *Settings) DarkMode() {
	if goruntime.GOOS == "windows" {
		runtime.WindowSetDarkTheme(s.ctx)
	}
	runtime.WindowSetBackgroundColour(s.ctx, DarkBackground.R, DarkBackground.G, DarkBackground.B, DarkBackground.A)
}

func (s *Settings) LightMode() {
	if goruntime.GOOS == "windows" {
		runtime.WindowSetLightTheme(s.ctx)
	}
	runtime.WindowSetBackgroundColour(s.ctx, LightBackground.R, LightBackground.G, LightBackground.B, LightBackground.A)
}

// only available for Windows
func (s *Settings) ResetMode() {
	if goruntime.GOOS == "windows" {
		runtime.WindowSetSystemDefaultTheme(s.ctx)
	}
}
