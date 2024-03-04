package mickey

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

type Command struct {
	FullScreen bool `help:"Use full screen for GUI window." short:"f"`
}

func (c *Command) Help() string {
	return `For a more graphical user experience.`
}

func (c *Command) Run(p *clit.Globals) error {
	// mickey command hook
	fe.SetGlobals(p)
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
