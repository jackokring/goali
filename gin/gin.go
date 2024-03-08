package gin

//=====================================
//*********** Tea Models **************
//=====================================

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackokring/goali/consts"
)

// A default tea Model
type Model struct {
	spinner spinner.Model
}

// User interaction channel
var UserChan = make(chan Model)

// System state reporting channel
var SystemChan = make(chan int)

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

// Model update function (use select/case on SystemChan)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			//m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

// The TUI goroutine to thread the TUI
func Tui() {
	p := tea.NewProgram(initialModel())
	// functional closure on p
	go func() {
		m, err := p.Run()
		if err != nil {
			close(UserChan) // check _, ok := ... for error state on user channel via select/case
		}
		UserChan <- m.(Model) // return of Model implies correct termination of user channel
	}()
}
