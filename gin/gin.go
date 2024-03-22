package gin

//=====================================
//*********** Tea Models **************
//=====================================

import (
	"fmt"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackokring/goali/consts"
	fe "github.com/jackokring/goali/filerr"
)

// Run after completion interface for Model
type PostAction interface {
	RunAfter()
}

// A default tea Model
type Model struct {
	spinner    spinner.Model
	keys       keyMap
	help       help.Model
	text       string
	PostAction // why when unnamed and embedded leads to != nil tautology?
	// ah, the auto method inheritance of unnamed embedded interfaces
}

// A simple null function concrete run after implementation
type knol struct{}

func (k knol) RunAfter() {}

// Quit the TUI
type QuitMsg struct {
	tea.Msg // embedded container has receivers
}

// Action message (extend for more specifics)
type ActionMsg struct {
	tea.Msg
	// expanded data
	string // string to set on action
}

// PostAction message (after TUI close)
type PostActionMsg struct {
	tea.Msg
	PostAction
}

// User channel to return model on TUI quit
var userChan = make(chan Model)

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
		spinner:    s,
		keys:       keys,
		help:       help.New(),
		text:       "",
		PostAction: knol{}, // empty struct with no action
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
	str := fmt.Sprintf("\n\n\t%s %s ...\n\n", m.spinner.View(), m.text)
	helpView := m.help.View(m.keys)
	height := viewHeight - strings.Count(str, "\n") - strings.Count(helpView, "\n")
	return "\n" + str + strings.Repeat("\n", height) + helpView
}

// Model update function (uses select/case on systemChan as not mutable)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	case ActionMsg:
		// decode action messages
		m.text = msg.string
	case PostActionMsg:
		// set a clean up action based on final Model
		m.PostAction = msg.PostAction
	case QuitMsg:
		// exit request
		return m, tea.Quit
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil // default return values unless specified earlier
}

// Tea the gin TUI export of the send message function
var Tea func(tea.Msg)

// The IO channel file lock (prevents TUI IO)
var cond sync.Cond = *sync.NewCond(&lock)
var lock sync.Mutex

func wait() {
	cond.L.Lock() // has the IO been unlocked?
	cond.Wait()   // wait for IPO setup
	cond.L.Unlock()
}

func Signal() {
	cond.Signal()
}

// The TUI goroutine to thread the TUI (returns message send function pointer)
func Tui() {
	p := tea.NewProgram(initialModel())
	// p.send(msgType)
	// functional closure on p
	go func() {
		wait() // wait
		// the case for commoning out a lock and not an Tui in
		// each command comes down to the RunAfter() interface
		m, err := p.Run()
		if err != nil {
			close(userChan) // check _, ok := ... for error state on user channel via select/case
		}
		userChan <- m.(Model) // return of Model implies correct termination of user channel
	}()
	Tea = p.Send
}

// Get TUI model of TUI interaction if ok is true
func TuiGetModel() (m Model, ok bool) {
	select {
	case u, ok := <-userChan:
		if !ok {
			fe.Fatal(fmt.Errorf("internal gin channel closed unexpectedly"))
		}
		// completed
		return u, true
	default:
		// defaults to blank
		return Model{}, false
	}
}
