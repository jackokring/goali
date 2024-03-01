//=====================================
//******* Packaging Section ***********
//=====================================

package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"

	py "github.com/jackokring/cpy3"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"

	"log"
	"log/syslog"
)

// # Application name
const AppName = "goali"
const AppDescription = "The " + AppName + " ball saving all in one app."
const maxVerbose = 3

//=====================================
//****** TUI structure section ********
//=====================================

type model struct {
	status int
	err    error
}

func checkServer() tea.Msg {

	// Create an HTTP client and make a GET request.
	/* c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		// There was an error making our request. Wrap the error we received
		// in a message and return it.
		return errMsg{err}
	}
	// We received a response from the server. Return the HTTP status code
	// as a message.
	return statusMsg(res.StatusCode) */
	return statusMsg(0)
}

type statusMsg int

type errMsg struct{ err error }

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		// The server returned a status message. Save it to our model. Also
		// tell the Bubble Tea runtime we want to exit because we have nothing
		// else to do. We'll still be able to render a final view with our
		// status message.
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		// Ctrl+c exits. Even with short running programs it's good to have
		// a quit key, just in case your logic is off. Users will be very
		// annoyed if they can't exit.
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	// If we happen to get any other messages, don't do anything.
	return m, nil
}

func (m model) View() string {
	// If there's an error, print it out and don't do anything else.
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	// Tell the user we're doing something.
	s := fmt.Sprintf("Checking %s ... ", "??") //url)

	// When the server responds with a status, add it to the current line.
	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	// Send off whatever we came up with above for rendering.
	return "\n" + s + "\n\n"
}

//=====================================
//****** CLI Structure Section ********
//=====================================

// optional:""
// type:"path"	A path. ~ expansion is applied. - is accepted for stdout, and will be passed unaltered.
// type:"existingfile"	An existing file. ~ expansion is applied. - is accepted for stdin, and will be passed unaltered.
// type:"existingdir"	An existing directory. ~ expansion is applied.
// type:"counter"	Increment a numeric field. Useful for -vvv. Can accept -s, --long or --long=N.
// type:"filecontent"	Read the file at path into the field. ~ expansion is applied. - is accepted for stdin, and will be passed unaltered.
// env:"X,Y,..."	Specify envars to use for default value. The envs are resolved in the declared order. The first value found is used.
// default:"X"	Default value.
// short:"X"	Short name, if flag.
// negatable:""	If present on a bool field, supports prefixing a flag with --no- to invert the default value.
// envprefix:"X"	Envar prefix for all sub-flags.
// passthrough:""	If present on a positional argument, it stops flag parsing when encountered, as if -- was processed before. Useful for external command wrappers, like exec. On a command it requires that the command contains only one argument of type []string which is then filled with everything following the command, unparsed.

// Flag allocation matrix fio and goa add to any tool with file IO and global options
//flg  goa fio uni mic
// a
// b
// c		x	x
// d	x		x	x
// e		x	x
// f		x	x	x
// g		x	x
// h	x 		x	x
// i
// j
// k
// l
// m
// n
// o
// p	x		x	x
// q	x		x	x
// r
// s	x		x	x
// t
// u
// v	x		x	x
// w		x	x
// x	x		x	x
// y
// z

// TODO
type profile struct {
	Name string
}

type inputFile struct {
	Expand    bool   `help:"Expand with gzip the <input-file>" short:"e"`
	InputFile string `arg:"" help:"The <input-file> to ${appName} (- is STDIN)" type:"existingfile"`
}

type outputFile struct {
	Compress   bool   `help:"Compress with gzip the <output-file>" short:"c"`
	Force      bool   `help:"Force overwriting of an existing <output-file>" short:"f"`
	Group      bool   `help:"The <output-file> is restricted to user and group access permissions" short:"g"`
	OutputFile string `arg:"" help:"The <output-file> from ${appName} (- is STDOUT maybe -q)" type:"path"`
	Write      bool   `help:"The <output-file> gains group write access permission" short:"w"`
}

type streamFilter struct {
	// special flags?
	inputFile
	outputFile
}

type guiCommand struct {
	FullScreen bool `help:"Use full screen for GUI window." short:"f"`
}

func (c *guiCommand) Help() string {
	return `For a more graphical user experience.`
}

func (c *guiCommand) Run(p *kong.Context) error {
	// mickey command hook
	Notify(p.Command())
	// interactive GUI mode
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetFullScreen(c.FullScreen)
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
	return nil
}

type unicornCommand struct {
	streamFilter // embedded type .. => .
}

func (c *unicornCommand) Help() string {
	return `Unicode mangling depending on the flags.
UTF-8 errors are marked to recover data.`
}

func (c *unicornCommand) Run(p *kong.Context) error {
	// unicorn command hook
	Notify(p.Command())
	fmt.Println(c.InputFile)
	fmt.Println(c.OutputFile)
	return nil
}

type detail int

// Alphabetic sorting
var cli struct {
	Debug   bool    `help:"Enable debug mode (includes panic tracing)" short:"d"`
	ProFile profile `help:"Configuration PROFILE of ${appName}" type:"yamlfile" short:"p"`
	Quiet   bool    `help:"Enable quiet mode errors (overrides -v)" short:"q"`
	SysLog  bool    `help:"Enable syslog output" short:"s"`
	Verbose detail  `help:"Enable verbose mode detail (1 to ${maxVerbose})" short:"v" type:"counter"`
	Wrong   bool    `help:"Enable fail on first error wrong mode" short:"x"`
	// a classic start
	Mickey  guiCommand     `cmd:"" help:"GUI launcher"`
	Unicorn unicornCommand `cmd:"" help:"Unicode mangler"`
}

//=====================================
//****** Error Handler Section ********
//=====================================

// Notify the current logger writer.
func Notify(s any) {
	if cli.Quiet {
		return
	}
	if cli.SysLog { // external so OK as no collide TUI
		log.Print(s)
	} else { // external not used
		// - STDOUT target situation barrier
		// auto quiet
		if os.Stderr != os.Stdout {
			log.Print(s)
		}
	}
}

// Error not nil checker syntax sugar
func Error(e error) bool {
	// better naming but should var args limit a callback?
	// indicated in a function signature by
	// func (rx type) name (args) { body } {}
	// you know, as an empty {} ... filled on use?
	//
	if e != nil {
		Notify(e) // {} here handler
		return true
	}
	return false
}

// Fatal error logging
func Fatal(e error) {
	if Error(e) {
		if cli.Debug {
			// this should always drop somewhere
			log.Panic(e)
		}
		//
		// No Notify() proxy as serious terminal error
		log.Fatal("FATAL: ", e)
	}
}

// Hard error check logging
//
// Check to see if the hard error flag is set
func Hard(e error) bool {
	if cli.Wrong {
		Fatal(e)
		return false
	}
	return Error(e)
}

// Notify a debug message to the current logger writer.
func Debug(s any) {
	if cli.Debug {
		Notify(s)
	}
}

// Verbose measure of output status to show
func Verbose() int {
	if cli.Quiet { // quiet or STDOUT priority?
		return 0
	}
	v := int(cli.Verbose)
	if v < 0 {
		v = 0
	}
	if v > maxVerbose {
		v = maxVerbose
	}
	return v
}

//=====================================
//**** File Abstraction Section *******
//=====================================

// GetIO
//
// Open reader before writer for error
// interaction effect with STDOUT by "-".
func GetIO(i string, expand bool,
	o string, compress bool, force bool, group bool, write bool) (FilterReadCloser, FilterWriteCloser) {
	return GetReader(i, expand), GetWriter(o, compress, force, group, write)
}

// FilterReadCloser is an abstraction to allow the wrapped
// unfiltered streams to be closed possibly by cascade calling.
type FilterReadCloser interface {
	// io.EOF
	Read(b []byte) (n int)
	CloseWrapped(closeWrapped bool)
	EOF() bool
}

// A concrete GZip FilterReadCloser
type GReader struct {
	this io.ReadCloser
	// is it a Closer => this == nil
	// and some funny business
	funny io.Reader
	// the wrapped or inner Closer
	wrap io.ReadCloser
	// requires pointer receiver
	// so all instances must be
	// by address &.
	eof bool
}

func (r GReader) CloseWrapped(closeWrapped bool) {
	situation := closeWrapped && r.wrap != nil
	if r.this != nil {
		// if there is an error is it required to close wrapped?
		if situation {
			Error(r.this.Close())
		} else {
			Fatal(r.this.Close())
		}
	}
	// does a wrapped need closing?
	if situation {
		Fatal(r.wrap.Close())
	}
}

// N.B. Due to needing to alter the EOF state
// the *GReader becomes required. This causes
// a cascade to need all interface instances
// to need pointer to value by &. Otherwise
// how would the pointer refer to that
// which is to be modified?
func (r *GReader) Read(b []byte) (n int) {
	if r.this == nil {
		// EOF suitable for while not EOF ...
		n, e := r.funny.Read(b)
		if e == io.EOF {
			r.eof = true
		}
		return n
	}
	n2, e2 := r.this.Read(b)
	if e2 == io.EOF {
		r.eof = true
	}
	return n2
}

func (r GReader) EOF() bool {
	return r.eof
}

// FilterWriteCloser is an abstraction to allow the wrapped
// unfiltered streams to be closed possibly by cascade calling.
type FilterWriteCloser interface {
	// io.EOF? on writing?
	Write(b []byte)
	CloseWrapped(closeWrapped bool)
}

// A concrete GZip FilterWriteCloser
type GWriter struct {
	this io.WriteCloser
	// is it a Closer => this == nil
	// and some funny business
	funny io.Writer
	// the wrapped or inner Closer
	wrap io.WriteCloser
}

func (r GWriter) CloseWrapped(closeWrapped bool) {
	situation := closeWrapped && r.wrap != nil
	if r.this != nil {
		// if there is an error is it required to close wrapped?
		if situation {
			Error(r.this.Close())
		} else {
			Fatal(r.this.Close())
		}
	}
	// does a wrapped need closing?
	if situation {
		Fatal(r.wrap.Close())
	}
}

func (r GWriter) Write(b []byte) {
	if r.this == nil {
		_, e := r.funny.Write(b)
		Fatal(e)
		return
	}
	_, e := r.this.Write(b)
	Fatal(e)
}

// Get reader
func GetReader(s string, expand bool) FilterReadCloser {
	if s == "-" {
		in := os.Stdin
		nin, e := os.Open(os.DevNull)
		Fatal(e)
		os.Stdin = nin
		return &GReader{in, nil, nil, false}
	}
	f, err := os.Open(s)
	Fatal(err)
	if expand {
		f2, err2 := gzip.NewReader(f)
		Fatal(err2)
		return &GReader{f2, nil, f, false}
	}
	return &GReader{nil, bufio.NewReader(f), nil, false}
}

// Get writer
func GetWriter(s string, compress bool, force bool, group bool, write bool) FilterWriteCloser {
	if s == "-" {
		out := os.Stdout
		// Handle TUI expectations
		os.Stdout = os.Stderr
		// already -q as command may have Notify()
		// on logger mixing
		return GWriter{out, nil, nil}
	}
	if force {
		os.Remove(s) // delete to force
	}
	// create if not exist <- N.B.
	var perms fs.FileMode = 0644
	if write && !group {
		perms = 0664 // give group permissive permissions
	}
	if group && !write {
		perms = 0640 // remove everybody permissions
	}
	if group && write {
		perms = 0660
	}
	f, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE|os.O_EXCL, perms)
	Fatal(err)
	if compress {
		f2, err2 := gzip.NewWriterLevel(f, gzip.BestCompression)
		Fatal(err2)
		return GWriter{f2, nil, f}
	}
	return GWriter{nil, bufio.NewWriter(f), nil}
}

//=====================================
//********** Main Section *************
//=====================================

// Find the system configuration directory
func SystemConfigDir() string { // Linux
	// yes you're crazy configuration sets are a NO!
	// i guess the first is the best as per $PATH
	// and it's not an over merge apply of last to first
	systemConfig := strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
	if len(systemConfig) == 0 { // Windows
		systemConfig = []string{os.Getenv("PROGRAMDATA")}
		if len(systemConfig[0]) == 0 { // MacOS
			systemConfig[0] = "/Library/Application Support"
		}
	}
	// should be fine on a well configured system
	return systemConfig[0]
}

// # Main Entry Point
func main() {
	// full config loading
	// the pro-file sub tree can be supplied from a file on the CLI
	//globalConfig := "/etc/" + AppName + "/config.yaml"
	dir, err := os.UserConfigDir() // {
	// should the error handler go here syntax wise??
	// tuple implicit?
	// }
	if err != nil { // Error(err) not used as not critical
		dir2, err2 := os.UserHomeDir()
		// pretty critical to have a home directory?
		// maybe some sort of demon process?
		// Fatal(err2)
		dir = dir2
		if err2 != nil {
			// pretends to be Darwin on failing
			dir = SystemConfigDir()
		}
	}
	localConfig := filepath.Join(dir, "."+AppName+".yaml")
	// Now we can parse
	ctx := kong.Parse(&cli,
		kong.Configuration(kongyaml.Loader /* globalConfig, */, localConfig),
		kong.Vars{
			// ${<name>} in `tags`
			//"globalConfig": globalConfig,
			"localConfig": localConfig,
			"appName":     AppName,
			"maxVerbose":  strconv.Itoa(maxVerbose),
		},
		// loading defaults for flags and options
		kong.NamedMapper("yamlfile", kongyaml.YAMLFileMapper),
		kong.Description(AppDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
	)
	log.SetOutput(os.Stderr)
	debug := 0
	if cli.Debug {
		debug += log.Lshortfile | log.Lmicroseconds
	}
	log.SetFlags(log.LstdFlags | log.LUTC | debug)
	//Error(errors.New("Error test"))
	if cli.SysLog {
		// Configure logger to write to the syslog.
		logwriter, e := syslog.New(syslog.LOG_NOTICE, AppName)
		Fatal(e)
		log.SetOutput(logwriter)
	}
	// Call the Run() method of the selected parsed command.
	// Extra context arg as not cast to command
	Fatal(ctx.Run(&ctx))
	// So you've found an Error?
	// Have you considered using the functions:
	//
	// 		func Error(err) bool // in an if test handler
	//		func Fatal(err) // anywhere
	//		func Hard(err)	bool // anywhere
	//
	// As these will provide panic info with the -d option.
	//
	//		func (*command)Run(*kong.Context) error
	//
	// Yes, returning a nil is an option. Source code error?

	defer py.Py_Finalize()
	py.Py_Initialize()
	py.PyRun_SimpleString("print('hello world')")

	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
