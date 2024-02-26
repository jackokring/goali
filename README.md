# Goali

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Modules Used

* `github.com/alecthomas/kong` - CLI parse
* `fyne.io/fyne/v2` - GUI toolkit
* `github.com/jackokring/cpy3` - Python 3.11 wrapper (submodule)
* `github.com/charmbracelet/bubbletea` - TUI toolkit

## Python `venv`

A python virtual environment was added to the project. Its major directories
`lib`, and `bin` were added to `.gitignore`. This may be altered later 
depending on functional use.

