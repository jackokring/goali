# Goali

Not to be confused with the non-existent marsh-soup-eal the Koali.

## Install
```
$ sudo apt install git python-is-python3 golang python3-pip python3-dev
$ git clone git@github.com:jackokring/goali.git
$ cd goali
$ # the cpy3 submodule for embedding python
$ git submodule update --init --recursive
$ # pull and satisfy requirements then go build
$ ./require.sh
```

## Fork
```
$ # go to repository directory
$ cd goali
$ # use your fork URL
$ git remote set-url origin git@github.com:<username>/goali.git
$ git push origin master
$ # check requirements and git add commit push
$ ./freeze.sh
```
And then create pull requests if you feel like it.

## Links

* [Goali Blog][blog] A github pages site. Keeps the speculation off this readme.
* [TODO bot][todo] A github bot using commented `@todo`/`@body` to raise automatic issues.
* [A Goali][David] This repository has nothing to do with footie or crisps.

## `goali` Main Commands

* `unicorn` Unicode mangler. Also [The Word][the_word]...
* `snake` Embedded and expanded python...
* `mickey` Nice GUI for something...
* `knap` Web servia thing...

**Under Construction TODO...**

## Features

* `--pro-file=PROFILE` process configuration override with `.yaml` files.
* Logging and `os.Stderr` fallback.
* `os.Stdin`/`os.Stdout` use by `-` filename.
* GZip IO wrapping CLI options.
* Use `c`, `python` and/or `go` to develop possibilities.

## Coding Framework Features

* Debug 'n panic `Fatal(err) void`/`Error(err) bool` framework.
* `FilterWriter.Rollback()` for a non-commit `Close()`.

---

## Why Go?

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Go Modules Used (and Indirect)

* [kong][kong] - CLI parse
* [fyne][fyne] - GUI toolkit
* [chi][chi] - A kind of node.js express router in go
* [cpy3][cpy3] - Python 3.11 wrapper `git submodule update --init --recursive` ([CPY3.md][CPY3])
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

---

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

---

## Extracted Link Definitions (Invisible)

[blog]: jackokring.github.io/goali
[bubbles]: github.com/charmbracelet/bubbles
[bubbletea]: github.com/charmbracelet/bubbletea
[chi]: github.com/go-chi/chi/v5
[cpy3]: github.com/jackokring/cpy3
[CPY3]: CPY3.md
[fyne]: fyne.io/fyne/v2
[kong]: github.com/alecthomas/kong
[kong-yaml]: github.com/alecthomas/kong-yaml
[lipgloss]: github.com/charmbracelet/lipgloss
[todo]: https://todo.jasonet.co

[the_word]: docs.google.com/document/d/1rsPyq3c7uVzxpUb9JXtq0b603HSuju7NWeZ_aYfVkzs/edit?usp=sharing
[David]: en.wikipedia.org/wiki/David_Seaman