package consts

//=====================================
//********* Global Constants **********
//=====================================

import "github.com/charmbracelet/lipgloss"

// The application name.
const AppName = "goali"

// The application description.
const AppDescription = "The " + AppName + " ball saving all in one app."

// The application version.
const Version = "0.1.0"

// The date time format string.
const DateTimeFormat = "2006-01-02 Mon 15:04:05 MST"

// Default CLI styles (pseudo constants)
var Gloss = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))
