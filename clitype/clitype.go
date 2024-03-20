package clitype

//=====================================
//******** Types Only Section *********
//=====================================

import "github.com/alecthomas/kong"

// Profile allows options beyond command line switches.
type Profile struct {
	// TODO
	Name string `yaml:"name"`
}

type Config kong.ConfigFlag

// The global switches used by the application.
type Globals struct {
	Debug      bool             `help:"Enable debug mode (includes panic tracing)." short:"d"`
	ProFile    Profile          `help:"Use PROFILE file (*.yaml) of ${appName}." type:"yamlfile" short:"p"`
	Quiet      bool             `help:"Enable quiet mode errors (overrides some of -d)." short:"q"`
	Rollback   bool             `help:"Enable rollback mode on fatal errors." short:"r"`
	SysLog     bool             `help:"Enable syslog output." short:"s"`
	TempConfig Config           `help:"Override configuration CONFIG file (*.yaml) of ${appName} (usually ${localConfig})." short:"t"`
	Version    kong.VersionFlag `help:"Show ${appName} version (${version})." short:"v"`
	Wrong      bool             `help:"Enable fail on first error wrong mode." short:"x"`
}

// A python code file type.
type PyFile struct {
	PyFile string `arg:"" help:"The <py-file> to execute (- is STDIN)." type:"existingfile"`
}

// An input file type.
type InputFile struct {
	Expand    bool   `help:"Expand with gzip the <input-file>." short:"e"`
	InputFile string `arg:"" help:"The <input-file> to ${appName} (- is STDIN)." type:"existingfile"`
}

// An io file type.
type IoFile struct {
	Compand bool   `help:"Compress and expand with gzip the <io-file>." short:"c"`
	Group   bool   `help:"The <io-file> is restricted to user and group access permissions." short:"g"`
	IoFile  string `arg:"" help:"The <io-file> to ${appName} (implies -f)." type:"existingfile"`
	Write   bool   `help:"The <io-file> gains group write access permission" short:"w"`
}

// An output file type.
type OutputFile struct {
	Compress   bool   `help:"Compress with gzip the <output-file>." short:"c"`
	Force      bool   `help:"Force overwriting of an existing <output-file>." short:"f"`
	Group      bool   `help:"The <output-file> is restricted to user and group access permissions." short:"g"`
	OutputFile string `arg:"" help:"The <output-file> from ${appName} (- is STDOUT maybe use -q)." type:"path"`
	Write      bool   `help:"The <output-file> gains group write access permission." short:"w"`
}

// A pair of files for IO.
type StreamFilter struct {
	// special flags?
	InputFile
	OutputFile
}
