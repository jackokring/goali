package clitype

//=====================================
//******** Types Only Section *********
//=====================================

// Profile allows options beyond command line switches.
type Profile struct {
	// TODO
	Name string
}

// The type decides what the assigned thing is called in the help
// hence level of DETAIL
type detail int

// The global switches used by the application.
type Globals struct {
	Debug   bool    `help:"Enable debug mode (includes panic tracing)" short:"d"`
	ProFile Profile `help:"Configuration PROFILE of ${appName}" type:"yamlfile" short:"p"`
	Quiet   bool    `help:"Enable quiet mode errors (overrides -v)" short:"q"`
	SysLog  bool    `help:"Enable syslog output" short:"s"`
	Verbose detail  `help:"Enable verbose mode detail (1 to ${maxVerbose})" short:"v" type:"counter"`
	Wrong   bool    `help:"Enable fail on first error wrong mode" short:"x"`
}

// An input file type.
type InputFile struct {
	Expand    bool   `help:"Expand with gzip the <input-file>" short:"e"`
	InputFile string `arg:"" help:"The <input-file> to ${appName} (- is STDIN)" type:"existingfile"`
}

// An output file type.
type OutputFile struct {
	Compress   bool   `help:"Compress with gzip the <output-file>" short:"c"`
	Force      bool   `help:"Force overwriting of an existing <output-file>" short:"f"`
	Group      bool   `help:"The <output-file> is restricted to user and group access permissions" short:"g"`
	OutputFile string `arg:"" help:"The <output-file> from ${appName} (- is STDOUT maybe -q)" type:"path"`
	Write      bool   `help:"The <output-file> gains group write access permission" short:"w"`
}

// A pair of files for IO.
type StreamFilter struct {
	// special flags?
	InputFile
	OutputFile
}
