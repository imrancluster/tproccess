package monitoring

import (
	"log"
	"time"

	"github.com/imrancluster/techmongo/fyne/tprocess/internal/network"

	"github.com/shirou/gopsutil/process"
)

// StartMonitoring handles the monitoring of system processes
func StartMonitoring() error {
	for {
		runningProcesses, err := getRunningProcesses()
		if err != nil {
			return err
		}

		// Filter processes for specific software like TeamViewer, AnyDesk
		filteredProcesses := filterProcesses(runningProcesses)

		// Mock sending the filtered processes to the server
		network.SendProcessesToServer(filteredProcesses)
		time.Sleep(20 * time.Second) // Monitor every 20 seconds
	}
}

// getRunningProcesses returns a list of running processes on the system
func getRunningProcesses() ([]string, error) {
	procs, err := process.Processes()
	if err != nil {
		log.Printf("Error fetching processes: %v", err)
		return nil, err
	}

	var processNames []string
	for _, proc := range procs {
		name, err := proc.Name()
		if err == nil {
			processNames = append(processNames, name)
		}
	}
	return processNames, nil
}
