package consts

//=====================================
//********* Global Constants **********
//=====================================

import "github.com/charmbracelet/lipgloss"

// The application name.
const AppName = "goali"

// The application description.
const AppDescription = "The " + AppName + " ball saving all in one app."

// The application version (pseudo constant).
var Version = "dev.feature.alpha"

// The application time stamp.
var BuildTime = "9999-13-32T24:60:60UTC"

// The date time format string.
const DateTimeFormat = "2006-01-02 Mon 15:04:05 MST"

// Default CLI style (pseudo constant).
var Gloss = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))

// The application build dynamic library poke (pseudo constant).
var Dynamic = "static"
