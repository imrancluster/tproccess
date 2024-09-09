package ui

import (
	"log"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/monitoring"
	"github.com/imrancluster/techmongo/fyne/tprocess/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// NewWindow creates a new Fyne window with the UI components
func NewWindow() fyne.Window {
	app := fyne.CurrentApp()
	myWindow := app.NewWindow("Candidate Monitoring App")
	myWindow.Resize(fyne.NewSize(400, 300)) // Set the standard window size

	// Set a light theme
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	// Status indicator label
	status := widget.NewLabel("Not Connected")

	// Connect button
	btn := widget.NewButton("Connect", func() {
		status.SetText("Connected")
		utils.Logger().Println("Button clicked. Monitoring started...")

		go func() {
			err := monitoring.StartMonitoring()
			if err != nil {
				log.Printf("Monitoring error: %v", err)
			}
		}()

		// Start keyboard monitoring in a separate goroutine
		go monitoring.StartKeyboardMonitoring()
	})

	// Arrange components in a vertical layout
	content := container.NewVBox(
		btn,
		status,
	)

	myWindow.SetContent(content)
	return myWindow
}
