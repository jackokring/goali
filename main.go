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
// prefix:"X"	Prefix for all sub-flags.
// envprefix:"X"	Envar prefix for all sub-flags.
// passthrough:""	If present on a positional argument, it stops flag parsing when encountered, as if -- was processed before. Useful for external command wrappers, like exec. On a command it requires that the command contains only one argument of type []string which is then filled with everything following the command, unparsed.

type profile struct {
	Yaml struct {
		Name string
	} `type:"yamlfile"`
}

type flagWithHelp bool

func (f *flagWithHelp) Help() string {
	return "🏁 no additional flag help"
}

type commandWithHelp struct {
	Msg argumentWithHelp `arg:"" help:"Regular argument help"`
}

func (c *commandWithHelp) Help() string {
	return "🚀 no additional command help"
}

func (c *commandWithHelp) Run(p *profile) error {
	fmt.Println(c.Msg.Msg)
	return nil
}

type argumentWithHelp struct {
	Msg string `arg:""`
}

func (f *argumentWithHelp) Help() string {
	return "📣 no additional argument help"
}

var CLI struct {
	Debug flagWithHelp `help:"Enable debug mode"`

	Flag flagWithHelp    `help:"Regular flag help"`
	Echo commandWithHelp `cmd:"" help:"Regular command help"`
}

func notify(s string) {
	// Now from anywhere else in your program, you can use this:
	log.Print(s)
}

func main() {
	// Configure logger to write to the syslog. You could do this in init(), too.
	logwriter, e := syslog.New(syslog.LOG_NOTICE, AppName)
	if e == nil {
		log.SetOutput(logwriter)
	}

	notify(AppName + " started")

	if len(os.Args) > 1 { // batch mode
		globalConfig := "/etc/" + AppName + "/config.yaml"
		localConfig := "~/." + AppName + ".yaml"
		ctx := kong.Parse(&CLI,
			kong.Configuration(kongyaml.Loader, globalConfig, localConfig),
			kong.Vars{
				// ${<name>} in `tags`
				"globalConfig": globalConfig,
				"localConfig":  localConfig,
			},
			kong.NamedMapper("yamlfile", kongyaml.YAMLFileMapper),
			kong.Description("The goali ball saving all in one app."),
			kong.UsageOnError(),
			kong.ConfigureHelp(kong.HelpOptions{
				Compact: true,
				Summary: false,
			}))
		// Call the Run() method of the selected parsed command.
		// Extra context arg? TODO
		err := ctx.Run(&profile{Yaml: struct{ Name string }{"running from main"}})
		ctx.FatalIfErrorf(err)
	} else { // interactive GUI mode
		a := app.New()
		w := a.NewWindow("Hello")

		hello := widget.NewLabel("Hello Fyne!")
		w.SetContent(container.NewVBox(
			hello,
			widget.NewButton("Hi!", func() {
				hello.SetText("Welcome :)")
			}),
		))

		w.ShowAndRun()
	}

	defer py.Py_Finalize()
	py.Py_Initialize()
	py.PyRun_SimpleString("print('hello world')")

	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
