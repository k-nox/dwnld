package app

import (
	"context"

	"github.com/k-nox/dwnld/app/download"
	"github.com/k-nox/dwnld/app/theme"
)

type App struct {
	ctx       context.Context
	Downloder *download.Downloader
	Theme     *theme.Settings
}

func New() *App {
	return &App{
		Theme:     &theme.Settings{},
		Downloder: &download.Downloader{},
	}
}

func Startup(app *App) func(context.Context) {
	return func(ctx context.Context) {
		app.ctx = ctx
		theme.Startup(ctx, app.Theme)
		download.Startup(ctx, app.Downloder)
	}
}
