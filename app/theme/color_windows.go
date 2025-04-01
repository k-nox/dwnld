package theme

import (
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	DarkForeground      = windows.RGB(248, 250, 252)
	DarkBorder          = windows.RGB(30, 41, 59)
	DarkMuted           = DarkBorder
	DarkMutedForeground = windows.RGB(148, 163, 184)

	LightForegroud       = windows.RGB(2, 8, 23)
	LightBorder          = windows.RGB(226, 232, 240)
	LightMuted           = windows.RGB(241, 245, 249)
	LightMutedForeground = windows.RGB(100, 116, 139)
)

func (s *Settings) DarkMode() {
	runtime.WindowSetDarkTheme(s.ctx)
}

func (s *Settings) LightMode() {
	runtime.WindowSetLightTheme(s.ctx)
}

// only available for Windows
func (s *Settings) ResetMode() {
	runtime.WindowSetSystemDefaultTheme(s.ctx)
}

func WindowsTheme() *windows.ThemeSettings {
	return &windows.ThemeSettings{
		// Active
		DarkModeTitleBar:  windows.RGB(DarkBackground.R, DarkBackground.G, DarkBackground.B),
		DarkModeTitleText: DarkForeground,
		DarkModeBorder:    DarkBorder,

		LightModeTitleBar:  windows.RGB(LightBackground.R, LightBackground.G, LightBackground.B),
		LightModeTitleText: LightForegroud,
		LightModeBorder:    LightBorder,

		// Inactive
		DarkModeTitleBarInactive:  DarkMuted,
		DarkModeTitleTextInactive: DarkMutedForeground,
		DarkModeBorderInactive:    DarkBorder,

		LightModeTitleBarInactive:  LightMuted,
		LightModeTitleTextInactive: LightMutedForeground,
		LightModeBorderInactive:    LightBorder,
	}
}
