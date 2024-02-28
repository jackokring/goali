package main

import (
	"fmt"

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

//=====================================
// TUI structure section
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
// CLI structure section
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

// TODO
type profile struct {
	Name string
}

type streamFilter struct {
	// special flags?
	Force      bool   `help:"Force overwrite of an existing <output-file>" short:"f"`
	InputFile  string `arg:"" help:"Input file to ${appName} (- is STDIN)" type:"existingfile"`
	OutputFile string `arg:"" help:"Output file of ${appName} (- is STDOUT implies -q)" type:"path"`
}

type guiCommand struct {
	FullScreen bool `help:"Use full screen for GUI window." short:"f"`
}

func (c *guiCommand) Help() string {
	return `For a more graphical user experience.`
}

func (c *guiCommand) Run(p *kong.Context) error {
	// mickey command hook
	Notify(AppName + " used")
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

type unicornCommand streamFilter

func (c *unicornCommand) Help() string {
	return `Unicode mangling depending on the flags.
UTF-8 errors are marked to recover data.`
}

func (c *unicornCommand) Run(p *kong.Context) error {
	// unicorn command hook
	Notify(AppName + " used")
	fmt.Println(c.InputFile)
	fmt.Println(c.OutputFile)
	return nil
}

type detail int

var cli struct {
	Debug   bool    `help:"Enable debug mode" short:"d"`
	Used    bool    `help:"Enable logging when used" short:"u"`
	Quiet   bool    `help:"Enable quiet mode (overrides -v)" short:"q"`
	Verbose detail  `help:"Enable verbose mode" short:"v" type:"counter"`
	ProFile profile `help:"Configuration PROFILE of ${appName}" type:"yamlfile"`
	// a classic start
	Unicorn unicornCommand `cmd:"" help:"Unicode mangler"`
	Mickey  guiCommand     `cmd:"" help:"GUI launcher"`
}

// Notify the current logger writer.
func Notify(s string) {
	log.Print(s)
}

// # Main Entry Point
func main() {
	// Configure logger to write to the syslog.
	logwriter, e := syslog.New(syslog.LOG_NOTICE, AppName)
	if e == nil {
		log.SetOutput(logwriter)
	}

	// full config loading
	// the pro-file sub tree can be supplied from a file on the CLI
	globalConfig := "/etc/" + AppName + "/config.yaml"
	localConfig := "~/." + AppName + ".yaml"
	ctx := kong.Parse(&cli,
		kong.Configuration(kongyaml.Loader, globalConfig, localConfig),
		kong.Vars{
			// ${<name>} in `tags`
			"globalConfig": globalConfig,
			"localConfig":  localConfig,
			"appName":      AppName,
		},
		// loading defaults for flags and options
		kong.NamedMapper("yamlfile", kongyaml.YAMLFileMapper),
		kong.Description("The "+AppName+" ball saving all in one app."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}))
	// Call the Run() method of the selected parsed command.
	// Extra context arg as not cast to command
	err := ctx.Run(&ctx)
	ctx.FatalIfErrorf(err)

	defer py.Py_Finalize()
	py.Py_Initialize()
	py.PyRun_SimpleString("print('hello world')")

	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
