package main

import (
	"embed"
	"os"
	"slices"

	"github.com/k-nox/dwnld/app"
	"github.com/k-nox/dwnld/app/config"
	"github.com/k-nox/dwnld/app/theme"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the a structure
	isDevMode := slices.Contains(os.Args, "dev")
	devModeFileLogger := isDevMode && slices.Contains(os.Args, "filelogger")

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
		OnStartup:        app.Startup(a, isDevMode),
		Bind: []interface{}{
			a,
		},
		EnumBind: []interface{}{
			config.Resolutions,
		},
		EnableDefaultContextMenu: true,
		Menu:                     app.Menu(a),
		Logger:                   customLogger(isDevMode, devModeFileLogger),
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   "dwnld",
				Message: "A dead-simple GUI for yt-dlp, powered by Wails and SvelteKit",
			},
		},
		Windows: &windows.Options{
			DisableWindowIcon: true,
			CustomTheme:       theme.WindowsTheme(),
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}

func customLogger(isDevMode, devModeFileLogger bool) logger.Logger {
	if isDevMode && !devModeFileLogger {
		return logger.NewDefaultLogger()
	}

	path, err := config.LogFile(isDevMode)
	if err != nil {
		panic(err)
	}

	return logger.NewFileLogger(path)
}
