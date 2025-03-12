package app

import (
	"context"

	"github.com/k-nox/dwnld/app/theme"
)

type App struct {
	ctx       context.Context
	Downloder *Downloader
	Settings  *theme.Settings
}

func New() *App {
	return &App{
		Downloder: &Downloader{},
		Settings:  &theme.Settings{},
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Settings.Ctx = ctx
	a.Downloder.ctx = ctx
	a.Downloder.ensureInstalled()
}
