package app

import (
	"context"

	"github.com/k-nox/dwnld/app/config"
	"github.com/k-nox/dwnld/app/download"
	"github.com/k-nox/dwnld/app/theme"
)

type App struct {
	ctx           context.Context
	Downloader    *download.Downloader
	Theme         *theme.Settings
	ConfigManager *config.Manager
}

func New(configFilePath string) (*App, error) {
	cfgMgr, err := config.New(configFilePath)
	if err != nil {
		return nil, err
	}
	return &App{
		Theme:         &theme.Settings{},
		Downloader:    &download.Downloader{},
		ConfigManager: cfgMgr,
	}, nil
}

func Startup(app *App) func(context.Context) {
	return func(ctx context.Context) {
		app.ctx = ctx
		settings, err := app.ConfigManager.Load()
		if err != nil {
			// TODO: handle
			panic(err)
		}
		theme.Startup(ctx, app.Theme)
		download.Startup(ctx, app.Downloader, settings.Download)
	}
}
