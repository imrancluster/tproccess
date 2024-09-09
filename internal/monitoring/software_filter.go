package monitoring

import "strings"

// filterProcesses filters out specific screen-sharing software like TeamViewer, AnyDesk
func filterProcesses(processes []string) []string {
	var filtered []string
	for _, proc := range processes {
		if isScreenSharingSoftware(proc) {
			filtered = append(filtered, proc)
		}
	}
	return filtered
}

// isScreenSharingSoftware checks if the process is a known screen-sharing software
func isScreenSharingSoftware(proc string) bool {
	knownScreenSharingSoftware := []string{
		"TeamViewer",
		"AnyDesk",
		"Zoom",
		"GoToMeeting",
		"Slack",
		"Microsoft Teams",
		"iTerm2",
	}

	for _, software := range knownScreenSharingSoftware {
		if strings.Contains(proc, software) {
			return true
		}
	}
	return false
}
