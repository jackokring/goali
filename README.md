# Goali

**Under Construction**

## Features

* User configuration `.yaml` files.
* `--pro-file=PROFILE` process configuration override.
* Logging and `os.Stderr` fallback.
* Debug 'n panic `Fatal(err) void`/`Error(err) bool` framework.
* `os.Stdin`/`os.Stdout` use by `-` filename. 

---

## Why Go?

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Modules Used (and Indirect)

* [kong][kong] - CLI parse
* [fyne][fyne] - GUI toolkit
* [cpy3][cpy3] - Python 3.11 wrapper (submodule [CPY3][CPY3])
* [bubbletea][bubbletea] - TUI toolkit
* [kong-yaml][kong-yaml] - config loader

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
* Allows submodules and redirection via `replace ... => ...` to local git submodules.
* Initial capitalization exports.
* `.(type)` switch variants.
* A mindset of default reentrant non-atomicity.

### Dislikes

* The postfix typing. As from a point of view it has no parse introduction of
type information **before** an identifier. This slightly increases parse complexity
and prevents some kind of edit selection of type DropDown inserting an
identifier TextBox.
* The weird `import "github.com/jackokring/v2/cpy3"` and making a new `v2` branch
along with a `module github.com/jackokring/cpy3/v2` just to allow a
`go get github.com/jackokring/cpy3/v2@v2.0.0` after a following
`GOPROXY=proxy.golang.org go list -m github.com/jackokring/cpy3/v2@v2.0.0` but
only after a `git tag v2.0.0` itself after a push followed by a `git push origin v2.0.0`.
* Of course the above is made more irritating by node.js inside VSCode spamming
the `-jMAX` option and behaving in its docs like the other processes on the
system are the problem. Apparently, the terminal failure to initialize on first go
is a feature, and it can only count to 120. 

## TODO

- [ ] `godoc` (version when # titles introduced is 1.19).
- [ ] more functionality.
- [ ] default `.py` file load.
- [ ] default add `PyObject` to namespace of python.
- [X] use `.goali.yaml` for CLI and general config options.

---

#### Extracted Link Definitions (Invisible)

[bubbletea]: github.com/charmbracelet/bubbletea
[cpy3]: github.com/jackokring/cpy3
[CPY3]: CPY3.md
[fyne]: fyne.io/fyne/v2
[kong]: github.com/alecthomas/kong
[kong-yaml]: github.com/alecthomas/kong-yaml