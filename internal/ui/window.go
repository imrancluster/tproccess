package ui

import (
	"image/color"
	"log"
	"os"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/monitoring"
	"github.com/imrancluster/techmongo/fyne/tprocess/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// NewWindow creates a new Fyne window with the UI components
func NewWindowOld() fyne.Window {
	app := fyne.CurrentApp()
	myWindow := app.NewWindow("Candidate Monitoring App")
	myWindow.Resize(fyne.NewSize(400, 250)) // Set the standard window size

	// Background Image
	// Load the SVG image from file
	svgFile := "tt_v2_logo.svg"
	if _, err := os.Stat(svgFile); os.IsNotExist(err) {
		myWindow.SetContent(canvas.NewText("SVG file not found!", theme.ForegroundColor()))
		myWindow.ShowAndRun()
	}

	// Create an image canvas from the SVG file
	bgImage := canvas.NewImageFromFile(svgFile)
	bgImage.FillMode = canvas.ImageFillStretch

	// Set a light theme
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	label2 := makeLabel("Not Connected", color.RGBA{R: 0, G: 0, B: 0, A: 255}, false)
	label1 := makeLabel("Connected", color.RGBA{R: 178, G: 222, B: 39, A: 255}, true)

	// Input Field to get Candiate Slug
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Candidate ID")

	// Connect button
	btn := widget.NewButton("Connect", func() {
		// Input text processing
		utils.Logger().Println("Input Text: " + input.Text)
		if input.Text == "" {
			dialog.ShowInformation("Error", "Please enter your candidate ID. You should get the ID on the browser!", myWindow)
			return
		}

		// status.SetText("Connected")
		utils.Logger().Println("Button clicked. Monitoring started...")

		label2.Hidden = true
		label1.Hidden = false

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
		bgImage,
		input,
		btn,
		// status,
		label1,
		label2,
	)

	myWindow.SetContent(content)
	return myWindow
}

func NewWindow() fyne.Window {
	app := fyne.CurrentApp()
	myWindow := app.NewWindow("Candidate Monitoring App")
	myWindow.Resize(fyne.NewSize(400, 250)) // Set the standard window size

	// Load the SVG image from file
	svgFile := "tt_v2_logo.png"
	if _, err := os.Stat(svgFile); os.IsNotExist(err) {
		myWindow.SetContent(canvas.NewText("SVG file not found!", theme.ForegroundColor()))
		myWindow.ShowAndRun()
		return myWindow // Return to avoid further execution
	}

	// Create an image canvas from the SVG file
	bgImage := canvas.NewImageFromFile(svgFile)
	bgImage.FillMode = canvas.ImageFillOriginal // Use original size for bottom-centered positioning

	// Set a light theme
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	label2 := makeLabel("Not Connected", color.RGBA{R: 0, G: 0, B: 0, A: 255}, false)
	label1 := makeLabel("Connected", color.RGBA{R: 178, G: 222, B: 39, A: 255}, true)

	// Input Field to get Candidate Slug
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Candidate ID")

	// Imran Sarder 
	// skdjfksdlj Imran kimran kjskdf 
	// Imran 

	// Connect button
	btn := widget.NewButton("Connect", btnFunc(input, label2, label1, myWindow))

	// Arrange components in a vertical layout
	uiContent := container.NewVBox(
		input,
		btn,
		label1,
		label2,
	)

	// Center the logo at the bottom
	logoContent := container.NewVBox(
		container.NewCenter(bgImage), // Center the logo horizontally
	)

	// Combine UI content and logo at the bottom
	finalContent := container.NewBorder(uiContent, logoContent, nil, nil)

	myWindow.SetContent(finalContent)
	return myWindow
}

// makeLabel
func makeLabel(name string, color color.RGBA, visibility bool) *canvas.Text {
	label := canvas.NewText(name, color)
	label.TextSize = 15
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	label.Hidden = visibility

	return label
}

func btnFunc(input *widget.Entry, label2 *canvas.Text, label1 *canvas.Text, myWindow fyne.Window) func() {
	return func() {
		// Input text processing
		utils.Logger().Println("Input Text: " + input.Text)
		if input.Text == "" {
			dialog.ShowInformation("Error", "Please enter your candidate ID. You should get the ID on the browser!", myWindow)
			return
		}

		utils.Logger().Println("Button clicked. Monitoring started...")

		label2.Hidden = true
		label1.Hidden = false

		go func() {
			err := monitoring.StartMonitoring()
			if err != nil {
				log.Printf("Monitoring error: %v", err)
			}
		}()

		// Start keyboard monitoring in a separate goroutine
		go monitoring.StartKeyboardMonitoring()
	}
}
