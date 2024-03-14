# Goali

**Under Construction**

- [ ] Unicode format checking and conversion.
- [ ] GUI interface.
- [ ] Python embedded, default extra available functions.
- [ ] TUI for all but GUI.
- [ ] ...

## Features

* User configuration `.yaml` files.
* `--pro-file=PROFILE` process configuration override.
* Logging and `os.Stderr` fallback.
* Debug 'n panic `Fatal(err) void`/`Error(err) bool` framework.
* `os.Stdin`/`os.Stdout` use by `-` filename.
* GZip IO wrapping CLI options.
* `FilterWriter.Rollback()` for a non-commit `Close()`.
* `snake.py` python module for interfacing and `mypy` type checking.

---

## Why Go?

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Modules Used (and Indirect)

* [kong][kong] - CLI parse
* [fyne][fyne] - GUI toolkit
* [chi][chi] - A kind of node.js express router in go
* [cpy3][cpy3] - Python 3.11 wrapper (`git submodule` [CPY3][CPY3])
* [bubbletea][bubbletea] - TUI toolkit
* [lipgloss][lipgloss] - ANSI sequences
* [bubbles][bubbles] - TUI extended controls
* [kong-yaml][kong-yaml] - config loader

## Modules to Find

This place is kind of a brain storming section, stuff in the decision matrix.

## Python `venv`

A python virtual environment was added to the project. Its major directories
`lib`, and `bin` were added to `.gitignore`. This may be altered later 
depending on functional use. A `snake_test.ipynb` test Jupyter notebook is included.

## Go Likes and Dislikes

### Likes

* It's fast enough at compile and doesn't clog my Chromebook.
* Receivers.
* No `switch` break madness fall through.
* Implicit interface implementations, without declaration and locking out
inbuilt type extension.
* Allows submodules and redirection via `replace ... => ...` to local a `git submodule`.
* Initial capitalization exports.
* `.(type)` switch variants. I like the two assignment tuple form.
* A mindset of default reentrant non-atomicity.
* Embedded `struct` types.
* Ah, `chan` of Occam. Or should that've been Ockham or Oakham? Nice `select` for
`ALT` too.
* The `error` strategy. Sure I've hacked it (see `filerr/filerr.go`).
* Although the idea of tacit tuple receivers excites me, they're not needed. 

### Dislikes

* The postfix typing. As from a point of view it has no parse introduction of
type information **before** an identifier. This slightly increases parse complexity
and prevents some kind of edit selection of type DropDown inserting an
identifier TextBox. It's not a deal breaker, as I've seen / used Pascal and
Modula-2 / Oberon before.
* The weird `import "github.com/jackokring/v2/cpy3"` and making a new `v2` branch
along with a `module github.com/jackokring/cpy3/v2` just to allow a
`go get github.com/jackokring/cpy3/v2@v2.0.0` after a following
`GOPROXY=proxy.golang.org go list -m github.com/jackokring/cpy3/v2@v2.0.0` but
only after a `git tag v2.0.0` itself after a push followed by a `git push origin v2.0.0`.
* Of course the above is made more irritating by node.js inside VSCode spamming
the `-jMAX` option and behaving in its docs like the other processes on the
system are the problem. Apparently, the terminal failure to initialize on first go
is a feature, and it can only count to 120. It doesn't understand `tmux`?
This is kind of "fixed" using VSCode as the terminal (apart from SIGHUP on close).
At least it has split views.
* The segmentation fault on `errors.New()`. But, it was new. Likely a pointer
to local returned from function and dereferenced.
* A mild `.cache` directory half a GB of "junk" from the language server. Perhaps
a `go tidy` for a chop down project global.
* The `import` deletion from a `^S` when there is an intent from `go get` and
no `go build` has been issued. I mean yes it does suggest and add, but does
it get the preferred named as data back (project or directory wide)?
* Mild annoyance with anonymous function closure syntax for `return` to achieve a
multi `break` and not even an ORIC-1 BASIC `POP:RETURN`. I'm not a fan of the
named labels and `goto` approach. Perhaps a `func {}` short form without `()`.
I mean I can't suggest a `return[level]` syntax as `[]int{x, y, z ...}` might
mix bad with the parser, being a prefixed typing.

### Wonderings

* Is `const` (as a type read only intent) inferred for auto `VAR` arguments (Pascal)?

### Things which are Really VSCode Things

* Eager launching of resource using tools. I'd like an awareness of `-j` which
does not use all my cores, as you know it's just a Debian container
also running a browser, with the music, and the reference internet. It
should hold off when close to memory 70% or core usage above 50% like
give the syntax check a stop for a minute or two. I only deleted a `}`. I'm sure
I'll manage without the extra or fewer items in the "intelli-drop-list".

## TODO

- [ ] `godoc` (version when # titles introduced is 1.19).
- [ ] more functionality.
- [ ] default `.py` file load.
- [ ] default add `PyObject` to namespace of python.
- [ ] sort out a good maths for verbose calculation (scaling?).
- [ ] UTF-8 processing decisions.

---

#### Extracted Link Definitions (Invisible)

[bubbles]: github.com/charmbracelet/bubbles
[bubbletea]: github.com/charmbracelet/bubbletea
[chi]: github.com/go-chi/chi/v5
[cpy3]: github.com/jackokring/cpy3
[CPY3]: CPY3.md
[fyne]: fyne.io/fyne/v2
[kong]: github.com/alecthomas/kong
[kong-yaml]: github.com/alecthomas/kong-yaml
[lipgloss]: github.com/charmbracelet/lipgloss