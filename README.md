# Goali

**Under Construction**

---

## Why Go?

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Modules Used (and Indirect)

* `github.com/alecthomas/kong` - CLI parse
* `fyne.io/fyne/v2` - GUI toolkit
* `github.com/jackokring/cpy3` - Python 3.11 wrapper (submodule [CPY3][CPY3])
* `github.com/charmbracelet/bubbletea` - TUI toolkit

## Modules to Find

This place is kind of a brain storming section, stuff in the decision matrix.

## Python `venv`

A python virtual environment was added to the project. Its major directories
`lib`, and `bin` were added to `.gitignore`. This may be altered later 
depending on functional use.

---

#### Extracted Link Definitions (Invisible)
[CPY3]: CPY3.md