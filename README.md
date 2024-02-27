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

## Go Likes and Dislikes

### Likes

* It's fast enough at compile and doesn't clog my Chromebook.
* Receivers.
* No `switch` break madness fall through.
* Implicit interface implementations.

### Dislikes

* The postfix typing. As from a point of view it has no parse introduction of
type information **before** an identifier. This slightly increases parse complexity
and prevents some kind of edit selection of type DropDown inserting an
identifier TextBox.

## TODO

- [ ] `godoc` (version when # titles introduced is 1.19).
- [ ] more functionality.
- [ ] default `.py` file load.
- [ ] default add `PyObject` to namespace of python.
- [ ] use `.yaml` for CLI and general config options.

---

#### Extracted Link Definitions (Invisible)
[CPY3]: CPY3.md