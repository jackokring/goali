# Goali

It might need a `export PKG_CONFIG_PATH=/home/jackokring/goali` kind of thing
before the `go build` but it works as a demo. This might be cloned later
to adapt the definitions to remove errors and warnings.

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

## Modules Used

* `github.com/alecthomas/kong` - CLI parse
* `fyne.io/fyne/v2` - GUI toolkit
* `github.com/sublime-security/cpy3` - Python wrapper
* `github.com/charmbracelet/bubbletea` - TUI toolkit

## Python `venv`

A python virtual environment was added to the project. Its major directories
`lib`, and `bin` were added to `.gitignore`.

