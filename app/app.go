package app

import (
	"context"
	"fmt"
	"os"

	"dario.cat/mergo"
	"github.com/k-nox/dwnld/app/config"
	"github.com/k-nox/dwnld/app/download"
	"github.com/k-nox/dwnld/app/theme"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	theme      *theme.Settings
	downloader *download.Downloader
	cfgMgr     *config.Manager
}

func New() *App {
	return &App{
		theme:      &theme.Settings{},
		downloader: &download.Downloader{},
	}
}

func Startup(app *App, isDevMode bool) func(context.Context) {
	return func(ctx context.Context) {
		app.ctx = ctx

		configFilePath, err := config.File(isDevMode)
		if err != nil {
			app.logErrAndQuit(err)
		}

		app.cfgMgr = config.New(configFilePath)

		if err := app.cfgMgr.Startup(); err != nil {
			app.logErrAndQuit(err)
		}

		app.theme.Startup(ctx)
		app.downloader.Startup(ctx)
	}
}

func Menu(app *App) *menu.Menu {
	aboutItem := menu.Text("About dwnld", nil, func(cd *menu.CallbackData) {
		runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Message: "A dead-simple GUI for yt-dlp, powered by Wails and SvelteKit",
		})
	})

	settingsItem := menu.Text("Settings", keys.CmdOrCtrl(","), func(cd *menu.CallbackData) { runtime.EventsEmit(app.ctx, "openSettings") })

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

func (a *App) Settings() *config.Settings {
	return a.cfgMgr.Settings()
}

func (a *App) UpdateSettings(updated config.Settings) error {
	if err := a.cfgMgr.Update(updated); err != nil {
		return a.logErrAndReturn(err)
	}
	return nil
}

func (a *App) Download(url string, opts config.Download) error {
	err := mergo.Merge(&opts, a.Settings().Download)
	if err != nil {
		return a.logErrAndReturn(fmt.Errorf("error merging opts with default opts: %w", err))
	}

	err = a.downloader.Download(url, opts)
	if err != nil {
		return a.logErrAndReturn(err)
	}
	return nil
}

func (a *App) DarkMode() {
	a.theme.DarkMode()
}

func (a *App) LightMode() {
	a.theme.LightMode()
}

func (a *App) ResetMode() {
	a.theme.ResetMode()
}

func (a *App) ChooseDirectory() (string, error) {
	opts := runtime.OpenDialogOptions{
		Title:                "Choose directory to save downloads",
		CanCreateDirectories: true,
	}

	curr := a.Settings().Download.OutputDirectory
	if curr != "" && isDir(curr) {
		opts.DefaultDirectory = curr
	}

	dir, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		return "", a.logErrAndReturn(fmt.Errorf("error choosing directory: %w", err))
	}
	return dir, nil
}

func isDir(dir string) bool {
	info, err := os.Stat(dir)
	return err == nil && info.IsDir()
}

func (a *App) logErrAndReturn(err error) error {
	runtime.LogError(a.ctx, err.Error())
	return err
}

func (a *App) logErrAndQuit(err error) {
	runtime.LogError(a.ctx, err.Error())
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Unable to start app due to error",
		Message: err.Error(),
	})
	runtime.Quit(a.ctx)
}
