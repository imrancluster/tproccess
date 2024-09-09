package app

import (
	"sync"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var once sync.Once
var instance fyne.App

// GetAppInstance returns a singleton instance of the fyne App
func GetAppInstance() fyne.App {
	once.Do(func() {
		instance = app.New()
	})
	return instance
}

// StartApp initializes and runs the application
func StartApp() {
	GetAppInstance()

	ui.NewWindow().ShowAndRun()
}
