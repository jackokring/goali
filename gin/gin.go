package gin

//=====================================
//*********** Tea Models **************
//=====================================

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackokring/goali/consts"
	"github.com/jackokring/goali/filerr"
)

// A default tea Model
type Model struct {
	spinner spinner.Model
}

//TODO: Sizing/rate limiting

// User interaction channel
var userChan = make(chan Model)

// System state reporting channel
var systemChan = make(chan Model)

func initialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = consts.Gloss
	return Model{spinner: s}
}

// Model initialization function
func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Model view function
func (m Model) View() string {
	str := fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", m.spinner.View())
	//if m.quitting { // basically indicates a new line should follow the ending of the TUI
	//	return str + "\n"
	//}
	return str
}

// Model update function (uses select/case on systemChan as not mutable)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mu := m
	var ok bool
	func() {
		for {
			select {
			case mu, ok = <-systemChan:
				// new model in keeps UI parts of model
				if ok {
					mu.spinner = m.spinner
				}
			default:
				return
			}
		}
	}()
	if !ok {
		// exit UI
		return mu, tea.Quit
	}
	switch msg := msg.(type) {
	case tea.KeyMsg: //TODO
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			//m.quitting = true
			return mu, tea.Quit
		default:
			return mu, nil
		}
	default:
		var cmd tea.Cmd
		mu.spinner, cmd = mu.spinner.Update(msg)
		return mu, cmd
	}
}

// The TUI goroutine to thread the TUI
func Tui() {
	p := tea.NewProgram(initialModel())
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
