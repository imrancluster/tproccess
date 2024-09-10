package main

import (
	"fmt"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/app"
)

func main() {
	// Start the monitoring app
	app.StartApp()

	// Event RunLoop
	// this function will excecute after quit your Gui App
	tidyUp()
}

func tidyUp() {
	fmt.Println("Program Exited!")
}
