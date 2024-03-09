package gin

//=====================================
//*********** Tea Models **************
//=====================================

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackokring/goali/consts"
	"github.com/jackokring/goali/filerr"
)

// A default tea Model
type Model struct {
	spinner spinner.Model
	keys    keyMap
	help    help.Model
}

//TODO: Sizing/rate limiting

// User interaction channel
var userChan = make(chan Model)

// System state reporting channel
var systemChan = make(chan Model)

type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		//{k.Up, k.Down, k.Left, k.Right}, // first column
		k.ShortHelp(), // default
	}
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func initialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = consts.Gloss
	return Model{
		spinner: s,
		keys:    keys,
		help:    help.New(),
	}
}

// Model initialization function
func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Default view height
const viewHeight = 8

// Model view function
func (m Model) View() string {
	str := fmt.Sprintf("\n\n   %s Loading forever...\n\n", m.spinner.View())
	helpView := m.help.View(m.keys)
	height := viewHeight - strings.Count(str, "\n") - strings.Count(helpView, "\n")
	return "\n" + str + strings.Repeat("\n", height) + helpView
}

// Model update function (uses select/case on systemChan as not mutable)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mu := m
	var ok bool = true // default ok with no channel read
	func() {
		for {
			select {
			case mu, ok = <-systemChan:
				// new model in keeps UI parts of model
				if ok {
					mu.spinner = m.spinner
					mu.keys = m.keys
					mu.help = m.help
				}
			default:
				// no message on channel
				return
			}
		}
	}()
	if !ok {
		// exit UI
		return mu, tea.Quit
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		mu.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, mu.keys.Help):
			mu.help.ShowAll = !mu.help.ShowAll
		case key.Matches(msg, mu.keys.Quit):
			return mu, tea.Quit
		}
	default:
		var cmd tea.Cmd
		mu.spinner, cmd = mu.spinner.Update(msg)
		return mu, cmd
	}
	return mu, nil // default return values unless specified earlier
}

// The TUI goroutine to thread the TUI
func Tui() {
	p := tea.NewProgram(initialModel())
	// p.send(msgType)
	// functional closure on p
	go func() {
		m, err := p.Run()
		if err != nil {
			close(userChan) // check _, ok := ... for error state on user channel via select/case
		}
		userChan <- m.(Model) // return of Model implies correct termination of user channel
	}()
}

// Get TUI model on completion of UI interaction if ok is true
func TuiGetModel() (m Model, ok bool) {
	select {
	case u, ok := <-userChan:
		if !ok {
			filerr.Fatal(fmt.Errorf("internal gin channel closed unexpectedly"))
		}
		// completed
		return u, true
	default:
		// defaults to blank
		return Model{}, false
	}
}

// Set TUI model (not rate limited in updates)
func TuiSetModel(m Model) {
	systemChan <- m
}

// Close TUI
func TuiClose() {
	close(systemChan)
}
