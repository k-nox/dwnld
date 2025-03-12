package app

import "context"

type App struct {
	ctx       context.Context
	Downloder *Downloader
}

func New() *App {
	return &App{
		Downloder: &Downloader{},
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Downloder.ctx = ctx
	a.Downloder.ensureInstalled()
}
