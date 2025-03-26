package app

import (
	"context"

	"github.com/k-nox/dwnld/app/config"
	"github.com/k-nox/dwnld/app/download"
	"github.com/k-nox/dwnld/app/theme"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func Menu(app *App) *menu.Menu {
	aboutItem := menu.Text("About dwnld", nil, func(cd *menu.CallbackData) { app.about() })

	settingsItem := menu.Text("Settings", keys.CmdOrCtrl(","), func(cd *menu.CallbackData) { app.settings() })

	hideItem := menu.Text("Hide dwnld", keys.CmdOrCtrl("h"), func(cd *menu.CallbackData) { runtime.Hide(app.ctx) })

	showAllItem := menu.Text("Show All", nil, func(cd *menu.CallbackData) { runtime.Show(app.ctx) })

	quitItem := menu.Text("Quit dwnld", keys.CmdOrCtrl("q"), func(cd *menu.CallbackData) { runtime.Quit(app.ctx) })

	item := &menu.MenuItem{
		Role: menu.AppMenuRole,
		SubMenu: menu.NewMenuFromItems(
			aboutItem,
			menu.Separator(),
			settingsItem,
			menu.Separator(),
			hideItem,
			showAllItem,
			menu.Separator(),
			quitItem,
		),
	}

	return menu.NewMenuFromItems(item, menu.EditMenu(), menu.WindowMenu())
}

func (a *App) about() {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Message: "A dead-simple GUI for yt-dlp, powered by Wails and SvelteKit",
	})
}

func (a *App) settings() {
	runtime.EventsEmit(a.ctx, "openSettings")
}
