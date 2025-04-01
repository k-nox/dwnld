//go:build !darwin

package app

import "github.com/wailsapp/wails/v2/pkg/menu"

func Menu(app *App) *menu.Menu {
	return nil
}
