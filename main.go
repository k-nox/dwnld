package main

import (
	"embed"

	"github.com/k-nox/dwnld/app"
	"github.com/k-nox/dwnld/app/theme"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the a structure
	a := app.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "dwnld",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: theme.DarkBackground,
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a.Downloder,
			a.Settings,
		},
		EnumBind: []interface{}{
			app.Resolutions,
		},
		EnableDefaultContextMenu: true,
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   "dwnld",
				Message: "A dead-simple GUI for yt-dlp, powered by Wails and SvelteKit",
			},
		},
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings{
				// Active
				DarkModeTitleBar:  windows.RGB(theme.DarkBackground.R, theme.DarkBackground.G, theme.DarkBackground.B),
				DarkModeTitleText: theme.DarkForeground,
				DarkModeBorder:    theme.DarkBorder,

				LightModeTitleBar:  windows.RGB(theme.LightBackground.R, theme.LightBackground.G, theme.LightBackground.B),
				LightModeTitleText: theme.LightForegroud,
				LightModeBorder:    theme.LightBorder,

				// Inactive
				DarkModeTitleBarInactive:  theme.DarkMuted,
				DarkModeTitleTextInactive: theme.DarkMutedForeground,
				DarkModeBorderInactive:    theme.DarkBorder,

				LightModeTitleBarInactive:  theme.LightMuted,
				LightModeTitleTextInactive: theme.LightMutedForeground,
				LightModeBorderInactive:    theme.LightBorder,
			},
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
