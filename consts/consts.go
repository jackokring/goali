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

//=====================================
//******* Debug Level Section *********
//=====================================

type DebugSource int
type DebugLevel int

const (
	// application debug source origin
	APP_DEBUG_SOURCE DebugSource = iota
	MAX_DEBUG_SOURCE
)

// Level of debug at which appearing in log happens
var DebugMin [MAX_DEBUG_SOURCE]DebugLevel = [MAX_DEBUG_SOURCE]DebugLevel{
	1, // APP
}

//=====================================
//******** Exit Codes Section *********
//=====================================

// The exit code type
type ExitCode int // uint8

// The primitive exit codes
//
// Some combinations are used as unique goali and TUI errors
const (
	// general error (used by kong errors, and other library errors)
	ERR_GENERAL ExitCode = 1 << iota
	// all fatal errors have this set (all non-fatal errors become this with -x option)
	ERR_FATAL
	// 4 bits of error code with flags in lover mask
	ERR_0
	//
	ERR_1
	//
	ERR_2
	//
	ERR_3
	// for some shell errors lower than 128 (64+n)
	ERR_SHELL
	// internal code (128+n) for signals like SIGHUP, SIGTERM etc.
	ERR_SIGNAL_HANDLER
	// (> uint8) maybe used but, some applications and the shell might not support the code
	ERR_RANGE_PLUS_ONE
)

const ( // maximum of 16 possible bit patterns before "shell" overflow
	// basic error code space for more specific errors
	// this one is special as it clears the lower 2 bits when used on its own
	// xx are bits ERR_FATAL ERR_GENERAL

	// ERR_RESET_UNCLASSIFIED error (in combination with:)
	// 00 - SUCCESS
	// 01 - General error
	// 10 - Fatal general error
	// 11 - WRONG option -x error
	ERR_RESET_UNCLASSIFIED ExitCode = iota << 2
	
	// ERR_STREAM File IO error (in combination with:)
	// 00 - TUI gin error
	// 01 - goali error log
	// 10 - Fatal stream error (Normal)
	// 11 - goali error RunAfter
	ERR_STREAM
	
	// ERR_PYTHON Snake error (not file IO related with:)
	// 00 - Unmodified SUCCESS in error error
	// 01 - 
	// 10 - Fatal python error (Normal)
	// 11 - 
	ERR_PYTHON
	
	E_03
	// ...
	E_10
	E_11
	E_12
	E_13
	//
	E_20
	E_21
	E_22
	E_23
	//
	E_30
	E_31
	E_32
	E_33
)

// The combination exit codes useful named list
const (
	ERR_MINUS_ONE     ExitCode = ExitCode(^0)                   // two's complement inversion
	ERR_SIGNAL_CTRL_C          = ERR_SIGNAL_HANDLER | ERR_FATAL // for example
	ERR_RANGE                  = ERR_RANGE_PLUS_ONE - 1         // 255 (technically also ERR_MINUS_ONE)
	ERR_WRONG                  = ERR_FATAL | ERR_GENERAL        // both as general non fatal made fatal
	ERR_SHELL_CMPLX            = ERR_SIGNAL_HANDLER | ERR_SHELL // both for some complex shell errors (196+n)
)
