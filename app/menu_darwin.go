package app

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
