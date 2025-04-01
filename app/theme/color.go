//go:build !windows

package theme

import (
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *Settings) DarkMode() {
	runtime.WindowSetBackgroundColour(s.ctx, DarkBackground.R, DarkBackground.G, DarkBackground.B, DarkBackground.A)
}

func (s *Settings) LightMode() {
	runtime.WindowSetBackgroundColour(s.ctx, LightBackground.R, LightBackground.G, LightBackground.B, LightBackground.A)
}

// only available for Windows
func (s *Settings) ResetMode() {}

func WindowsTheme() *windows.ThemeSettings {
	return nil
}
